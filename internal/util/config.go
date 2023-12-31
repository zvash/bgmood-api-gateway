package util

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	AppName                        string `mapstructure:"APP_NAME"`
	Environment                    string `mapstructure:"ENVIRONMENT"`
	GRPCServerAddress              string `mapstructure:"GRPC_SERVER_ADDRESS"`
	PrometheusHTTPServerAddress    string `mapstructure:"PROMETHEUS_HTTP_SERVER_ADDRESS"`
	AuthServiceGRPCServerAddress   string `mapstructure:"AUTH_SERVICE_GRPC_SERVER_ADDRESS"`
	FileServiceGRPCServerAddress   string `mapstructure:"FILE_SERVICE_GRPC_SERVER_ADDRESS"`
	CircleServiceGRPCServerAddress string `mapstructure:"CIRCLE_SERVICE_GRPC_SERVER_ADDRESS"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
