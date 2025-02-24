package config

import "time"

type Config struct {
	Env      string `env-default:"local"`
	Dsn      string `env-required:"true"`
	GRPC     GRPCConfig
	TokenTTL time.Duration `env-default:"1h"`
}
type GRPCConfig struct {
	Port    int
	Timeout time.Duration
}
