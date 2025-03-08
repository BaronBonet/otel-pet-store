package connect

import "time"

type Config struct {
	ListenAddress        string        `env:"PORT"                   envDefault:"8080"` // Set by heroku
	MaxConcurrentStreams uint32        `env:"MAX_CONCURRENT_STREAMS" envDefault:"1000"`
	IdleTimeout          time.Duration `env:"IDLE_TIMEOUT"           envDefault:"120s"`
}
