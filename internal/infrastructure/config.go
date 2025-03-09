package infrastructure

import (
	"github.com/BaronBonet/otel-pet-store/internal/adapters/handler/connect"
	"github.com/BaronBonet/otel-pet-store/internal/adapters/repository/postgres"
	"github.com/caarlos0/env/v11"
	"go.opentelemetry.io/collector/service/telemetry"
)

type ApplicationConfig struct {
	Telemetry telemetry.Config
	Handler   connect.Config
	Postgres  postgres.Config
}

func LoadConfig() (*ApplicationConfig, error) {
	config := ApplicationConfig{}
	opts := env.Options{
		RequiredIfNoDef: true,
	}
	if err := env.ParseWithOptions(&config, opts); err != nil {
		return nil, err
	}
	return &config, nil
}
