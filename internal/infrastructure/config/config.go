package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/stackus/dotenv"
	"os"
	"time"
)

type (
	GatewayConfig struct {
		URL        string `mapstructure:"url"`
		Timeout    int    `mapstructure:"timeout"`
		MaxRetries int    `mapstructure:"max_retries"`
	}

	AppConfig struct {
		Environment     string
		LogLevel        string                   `envconfig:"LOG_LEVEL" default:"DEBUG"`
		Gateways        map[string]GatewayConfig `mapstructure:"gateways"`
		ShutdownTimeout time.Duration            `envconfig:"SHUTDOWN_TIMEOUT" default:"30s"`
	}
)

func InitConfig() (cfg AppConfig, err error) {
	if err = dotenv.Load(dotenv.EnvironmentFiles(os.Getenv("ENVIRONMENT"))); err != nil {
		return
	}

	err = envconfig.Process("", &cfg)

	return
}
