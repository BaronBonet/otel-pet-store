package main

import (
	"context"
	"fmt"

	"github.com/BaronBonet/otel-pet-store/internal/adapters/handler/connect"
	"github.com/BaronBonet/otel-pet-store/internal/adapters/repository/postgres"
	"github.com/BaronBonet/otel-pet-store/internal/core"
	"github.com/BaronBonet/otel-pet-store/internal/infrastructure"
	"github.com/BaronBonet/otel-pet-store/internal/pkg/logger"
	"github.com/BaronBonet/otel-pet-store/internal/pkg/telemetry"
)

const (
	name = "http"
)

func main() {
	ctx := context.Background()
	cfg, err := infrastructure.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Could not load config %v", err))
	}

	otelConfig, err := telemetry.NewOtelConfig(
		telemetry.OtelConfig{
			Service: telemetry.OtelConfigService{
				Name:      name,
				Version:   infrastructure.Version,
				NameSpace: "petstore",
			},
			Exporter: telemetry.OtelConfigExporter{
				Exporter: telemetry.ExporterOTLPLocal,
			},
		},
	)
	if err != nil {
		panic(fmt.Sprintf("Could not create Otel config %v", err))
	}

	shutdownOtel, err := telemetry.SetupOTelSDK(
		ctx,
		*otelConfig,
	)
	if err != nil {
		panic(fmt.Sprintf("Could not set up OpenTelemetry SDK %v", err))
	}

	logger := logger.NewOTelLogger(name, infrastructure.Version)

	defer func() {
		logger.Info(ctx, "Shutting down Otel")
		if err := shutdownOtel(ctx); err != nil {
			logger.Fatal(ctx, "Failed to shutdown OpenTelemetry SDK", "error", err)
		}
	}()

	pool, err := postgres.CreateDBPool(ctx, cfg.Postgres, name)
	if err != nil {
		logger.Fatal(ctx, "Couldn't create pool", "error", err)
	}
	repo := postgres.New(pool)

	service := core.NewService(repo, logger)

	handler, err := connect.New(ctx, cfg.Handler, service, logger)
	if err != nil {
		logger.Fatal(ctx, "Could not create handler", "error", err)
	}

	handler.Serve(ctx)
}
