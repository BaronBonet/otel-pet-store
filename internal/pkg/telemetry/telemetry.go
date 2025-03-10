package telemetry

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/log"
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
	endpoint string
	Exporter Exporter
}
type OtelConfig struct {
	Service  OtelConfigService
	Exporter OtelConfigExporter
}

// NewOtelConfig creates a new configuration with the provided options.
func NewOtelConfig(config OtelConfig) (*OtelConfig, error) {
	config.Exporter.endpoint = "localhost:4317"

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	pid := os.Getpid()
	config.Service.instanceId = fmt.Sprintf("%s-%d", hostname, pid)
	return &config, nil
}

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
	// meterProvider, err := newMeterProvider(ctx, res, config)
	// if err != nil {
	// 	handleErr(err)
	// 	return
	// }
	// shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)
	// otel.SetMeterProvider(meterProvider)

	// Initialize the logger provider for logs with the configured exporter and resource.
	loggerProvider, err := newLoggerProvider(ctx, res, config)
	if err != nil {
		handleErr(err)
		return shutdown, err
	}
	shutdownFuncs = append(shutdownFuncs, loggerProvider.Shutdown)
	global.SetLoggerProvider(loggerProvider)

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
	opts := []otlptracegrpc.Option{
		otlptracegrpc.WithEndpoint(config.Exporter.endpoint),
	}

	if config.Exporter.Exporter == ExporterOTLPLocal {
		opts = append(opts, otlptracegrpc.WithInsecure())
	}

	traceExporter, err := otlptracegrpc.New(
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
// func newMeterProvider(
// 	ctx context.Context,
// 	res *resource.Resource,
// 	config OtelConfig,
// ) (*metric.MeterProvider, error) {
// 	opts := []otlpmetrichttp.Option{
// 		otlpmetrichttp.WithEndpoint(config.Exporter.endpoint),
// 		otlpmetrichttp.WithCompression(otlpmetrichttp.GzipCompression),
// 	}
//
// 	if config.Exporter.apiKey != "" {
// 		opts = append(opts, otlpmetrichttp.WithHeaders(map[string]string{
// 			"api-key": config.Exporter.apiKey,
// 		}))
// 	}
//
// 	if config.Exporter.Exporter == ExporterOTLPLocal {
// 		opts = append(opts, otlpmetrichttp.WithInsecure())
// 	}
//
// 	metricExporter, err := otlpmetrichttp.New(ctx, opts...)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	meterProvider := metric.NewMeterProvider(
// 		metric.WithReader(
// 			metric.NewPeriodicReader(
// 				metricExporter,
// 				metric.WithInterval(10*time.Second),
// 			),
// 		),
// 		metric.WithResource(res),
// 	)
//
// 	// Start runtime instrumentation
// 	if err := runtime.Start(runtime.WithMeterProvider(meterProvider)); err != nil {
// 		return nil, fmt.Errorf("failed to start runtime instrumentation: %w", err)
// 	}
//
// 	return meterProvider, nil
// }

func newLoggerProvider(
	ctx context.Context,
	res *resource.Resource,
	config OtelConfig,
) (*log.LoggerProvider, error) {
	fmt.Printf("Setting up log exporter to endpoint: %s\n", config.Exporter.endpoint)

	opts := []otlploggrpc.Option{
		otlploggrpc.WithEndpoint(config.Exporter.endpoint),
		otlploggrpc.WithCompressor("gzip"),
	}

	if config.Exporter.Exporter == ExporterOTLPLocal {
		opts = append(opts, otlploggrpc.WithInsecure())
		fmt.Println("Using insecure connection for logs")
	}

	opts = append(opts, otlploggrpc.WithRetry(otlploggrpc.RetryConfig{
		Enabled:         true,
		InitialInterval: 1 * time.Second,
		MaxInterval:     5 * time.Second,
	}))

	logExporter, err := otlploggrpc.New(ctx, opts...)
	if err != nil {
		fmt.Printf("ERROR creating log exporter: %v\n", err)
		return nil, err
	}

	// Create a processor pipeline for log exporting
	processor := log.NewBatchProcessor(
		logExporter,
	)

	// Create the logger provider with trace context, ensuring logs are correlated with traces
	loggerProvider := log.NewLoggerProvider(
		log.WithProcessor(processor),
		log.WithResource(res),
	)

	fmt.Println("Log provider successfully initialized and set as global provider")

	return loggerProvider, nil
}
