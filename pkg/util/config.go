package util

import (
	"sync"

	"github.com/spf13/viper"
)

var (
	config Config
	once   sync.Once
)

type Config struct {
	DBHost                string `mapstructure:"DB_HOST"`
	DBPort                int32  `mapstructure:"DB_PORT"`
	DBUser                string `mapstructure:"DB_USER"`
	DBPassword            string `mapstructure:"DB_PASS"`
	DBName                string `mapstructure:"DB_NAME"`
	Port                  int32  `mapstructure:"PORT"`
	PaginatorLimitDefault int32  `mapstructure:"PAGINATOR_LIMIT_DEFAULT"`
}

func LoadConfig(path string) (Config, error) {
	var err error
	once.Do(func() {
		viper.AddConfigPath(path)
		viper.SetConfigName("app")
		viper.SetConfigType("env")
		viper.AutomaticEnv()

		err = viper.ReadInConfig()
		if err != nil {
			return
		}

		err = viper.Unmarshal(&config)
	})
	return config, err
}

func GetConfig() Config {
	return config
}
