package postgres

type PoolConfig struct {
	MinConnections int32 `env:"DATABASE_POOL_MIN_CONN" envDefault:"1"`
	MaxConnections int32 `env:"DATABASE_POOL_MAX_CONN" envDefault:"20"`
}

type Config struct {
	DatabaseURL string `env:"DATABASE_URL"`
	Timezone    string `env:"DATABASE_TIMEZONE" envDefault:"UTC"`
	PoolConfig  PoolConfig
}
