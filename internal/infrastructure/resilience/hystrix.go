package resilience

import (
	"github.com/afex/hystrix-go/hystrix"
)

// TODO: move to a config file
func ConfigureHystrix() {
	hystrix.DefaultTimeout = 5000
	hystrix.DefaultMaxConcurrent = 20
	hystrix.DefaultErrorPercentThreshold = 50
	hystrix.DefaultSleepWindow = 10000
}

type HystrixConfig struct {
	Name                   string
	Timeout                int // Timeout in milliseconds
	ErrorPercentThreshold  int // Percentage of errors to open the circuit
	RequestVolumeThreshold int // Minimum number of requests before the circuit can trip
	SleepWindow            int // How long (in milliseconds) to wait after tripping before retrying
}

func InitializeHystrixCommand(config HystrixConfig) {
	hystrix.ConfigureCommand(config.Name, hystrix.CommandConfig{
		Timeout:                config.Timeout,
		ErrorPercentThreshold:  config.ErrorPercentThreshold,
		RequestVolumeThreshold: config.RequestVolumeThreshold,
		SleepWindow:            config.SleepWindow,
	})
}

// RunWithHystrix runs a function with a circuit breaker, allowing for retries and fallback.
func RunWithHystrix(commandName string, runFunc func() error, fallbackFunc func(error) error) error {
	// Execute the command with Hystrix circuit breaker.
	return hystrix.Do(commandName, runFunc, fallbackFunc)
}
