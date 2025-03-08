package telemetry

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

type Exporter int

const (
	ExporterUnknown Exporter = iota
	ExporterOTLPLocal
	ExporterNewRelic
)

type OtelConfigService struct {
	NameSpace  string
	Name       string
	Version    string
	instanceId string
}
type OtelConfigExporter struct {
	apiKey   string
	endpoint string
	Exporter Exporter
}
type OtelConfig struct {
	Service  OtelConfigService
	Exporter OtelConfigExporter
}

// Option applies an option to the configuration.
type Option func(config *OtelConfig) error

func WithAPIKey(key string) Option {
	return func(config *OtelConfig) error {
		if key == "" {
			return errors.New("API key cannot be empty")
		}
		config.Exporter.apiKey = key
		return nil
	}
}

// NewOtelConfig creates a new configuration with the provided options.
func NewOtelConfig(config OtelConfig, opts ...Option) (*OtelConfig, error) {
	for _, opt := range opts {
		if err := opt(&config); err != nil {
			return nil, err
		}
	}

	switch config.Exporter.Exporter {
	case ExporterOTLPLocal:
		config.Exporter.endpoint = "localhost:4318"
	case ExporterNewRelic:
		config.Exporter.endpoint = "otlp.nr-data.net"
		if config.Exporter.apiKey == "" {
			return nil, errors.New("API key is required for New Relic exporter")
		}
	case ExporterUnknown:
		return nil, errors.New("unknown exporter")
	}

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	pid := os.Getpid()
	config.Service.instanceId = fmt.Sprintf("%s-%d", hostname, pid)
	return &config, nil
}

// Taken from https://github.com/newrelic/newrelic-opentelemetry-examples/tree/main/getting-started-guides/go

// SetupOTelSDK bootstraps the OpenTelemetry pipeline.
// It initializes the resources, propagators, and providers for traces, metrics, and logs.
// If it does not return an error, make sure to call shutdown for proper cleanup.
func SetupOTelSDK(
	ctx context.Context,
	config OtelConfig,
) (shutdown func(context.Context) error, err error) {
	var shutdownFuncs []func(context.Context) error
	// shutdown is a closure that calls cleanup functions registered via shutdownFuncs.
	// The errors from the calls are joined.
	// Each registered cleanup will be invoked once.
	shutdown = func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}
	// handleErr calls shutdown for cleanup and makes sure that all errors are returned.
	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}
	// Create a resource with service metadata and system information.
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(config.Service.Name),
			semconv.ServiceVersion(config.Service.Version),
			semconv.ServiceNamespace(config.Service.NameSpace),
			semconv.ServiceInstanceID(config.Service.instanceId),
		),
		resource.WithOS(),
		resource.WithHost(),
		resource.WithProcess(),
	)
	if err != nil {
		handleErr(err)
		return
	}
	// Set up a propagator for context propagation across service boundaries.
	prop := newPropagator()
	otel.SetTextMapPropagator(prop)
	// Initialize the trace provider with the configured exporter and resource.
	tracerProvider, err := newTraceProvider(ctx, res, config)
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)
	otel.SetTracerProvider(tracerProvider)
	// Initialize the meter provider for metrics with the configured exporter and resource.
	meterProvider, err := newMeterProvider(ctx, res, config)
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)
	if config.Exporter.Exporter == ExporterNewRelic {
		otel.SetMeterProvider(meterProvider)
		// Start runtime instrumentation to collect runtime metrics.
		if err := runtime.Start(runtime.WithMeterProvider(meterProvider)); err != nil {
			handleErr(err)
			return shutdown, err
		}
		// Initialize the logger provider for logs with the configured exporter and resource.
		loggerProvider, err := newLoggerProvider(ctx, res, config)
		if err != nil {
			handleErr(err)
			return shutdown, err
		}
		shutdownFuncs = append(shutdownFuncs, loggerProvider.Shutdown)
		global.SetLoggerProvider(loggerProvider)
	}

	fmt.Println("OpenTelemetry SDK setup complete")
	fmt.Printf("Service: %s/%s (version: %s)\n",
		config.Service.NameSpace,
		config.Service.Name,
		config.Service.Version)
	fmt.Printf("Exporter: %s\n", config.Exporter.Exporter)

	return
}

// newPropagator sets up a composite propagator for context propagation.
// It includes trace context and baggage propagation.
func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

// newTraceProvider initializes a trace provider with a batch span processor.
// The trace provider is responsible for creating tracers and exporting trace data.
// Trace provide will use insecure http if we are running this locally
func newTraceProvider(
	ctx context.Context,
	res *resource.Resource,
	config OtelConfig,
) (*trace.TracerProvider, error) {
	headers := map[string]string{"api-key": config.Exporter.apiKey}

	opts := []otlptracehttp.Option{
		otlptracehttp.WithEndpoint(config.Exporter.endpoint),
		otlptracehttp.WithHeaders(headers),
	}

	if config.Exporter.Exporter == ExporterOTLPLocal {
		opts = append(opts, otlptracehttp.WithInsecure())
	}

	traceExporter, err := otlptracehttp.New(
		ctx,
		opts...,
	)
	if err != nil {
		return nil, err
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter,
			trace.WithMaxExportBatchSize(512),
			trace.WithBatchTimeout(5*time.Second),
			trace.WithMaxQueueSize(2048),
		),
		trace.WithResource(res),
		trace.WithSampler(trace.ParentBased(
			trace.AlwaysSample(),
			// Ensure children of dropped spans are also dropped
			trace.WithRemoteParentSampled(trace.AlwaysSample()),
		)),
	)
	return traceProvider, nil
}

// newMeterProvider initializes a meter provider for metrics collection.
// It uses a periodic reader to export metrics at regular intervals.
func newMeterProvider(
	ctx context.Context,
	res *resource.Resource,
	config OtelConfig,
) (*metric.MeterProvider, error) {
	metricExporter, err := otlpmetrichttp.New(
		ctx,
		otlpmetrichttp.WithEndpoint(config.Exporter.endpoint),
		otlpmetrichttp.WithCompression(otlpmetrichttp.GzipCompression),
		otlpmetrichttp.WithHeaders(map[string]string{
			"api-key": config.Exporter.apiKey,
		}),
	)
	if err != nil {
		return nil, err
	}
	meterProvider := metric.NewMeterProvider(
		metric.WithReader(
			metric.NewPeriodicReader(
				metricExporter,
				metric.WithInterval(3*time.Second),
			),
		),
		metric.WithResource(res),
	)
	return meterProvider, nil
}

func newLoggerProvider(
	ctx context.Context,
	res *resource.Resource,
	config OtelConfig,
) (*log.LoggerProvider, error) {
	// Print debug information about where logs are being sent
	fmt.Printf("Setting up log exporter to endpoint: %s\n", config.Exporter.endpoint)

	opts := []otlploghttp.Option{
		otlploghttp.WithEndpoint(config.Exporter.endpoint),
		otlploghttp.WithCompression(otlploghttp.GzipCompression),
	}

	// Only add API key if it's not empty
	if config.Exporter.apiKey != "" {
		opts = append(opts, otlploghttp.WithHeaders(map[string]string{
			"api-key": config.Exporter.apiKey,
		}))
	}

	if config.Exporter.Exporter == ExporterOTLPLocal {
		opts = append(opts, otlploghttp.WithInsecure())
		fmt.Println("Using insecure connection for logs")
	}

	// Add debug option to see what's happening with the exporter
	opts = append(opts, otlploghttp.WithRetry(otlploghttp.RetryConfig{
		Enabled:         true,
		InitialInterval: 1 * time.Second,
		MaxInterval:     5 * time.Second,
		MaxElapsedTime:  30 * time.Second,
	}))

	logExporter, err := otlploghttp.New(ctx, opts...)
	if err != nil {
		fmt.Printf("ERROR creating log exporter: %v\n", err)
		return nil, err
	}

	// Create a custom processor that logs when batches are exported
	processor := log.NewBatchProcessor(
		logExporter,
		// Add options to make batches flush more frequently during debugging
	)

	loggerProvider := log.NewLoggerProvider(
		log.WithProcessor(processor),
		log.WithResource(res),
	)

	fmt.Println("Log provider successfully initialized and set as global provider")

	return loggerProvider, nil
}
