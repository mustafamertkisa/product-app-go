package infrastructure

import "github.com/spf13/viper"

type Config struct {
	DBUrl string `mapstructure:"POSTGRES_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigType("env")
	viper.SetConfigName("app")
	viper.AddConfigPath(path)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
