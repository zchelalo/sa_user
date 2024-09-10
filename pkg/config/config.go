package config

import (
	"github.com/spf13/viper"
)

// Config contains the configuration of the application.
// The values are read from environment variables.
type Config struct {
	DBHost                string `mapstructure:"DB_HOST"`
	DBPort                int32  `mapstructure:"DB_PORT"`
	DBUser                string `mapstructure:"DB_USER"`
	DBPassword            string `mapstructure:"DB_PASS"`
	DBName                string `mapstructure:"DB_NAME"`
	Port                  int32  `mapstructure:"PORT"`
	PaginatorLimitDefault int32  `mapstructure:"PAGINATOR_LIMIT_DEFAULT"`
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
