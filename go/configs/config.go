package configs

import "github.com/spf13/viper"

type DatabaseConfig struct {
	Port string `mapstructure:"DB_PORT"`
	Host string `mapstructure:"DB_HOST"`
	User string `mapstructure:"DB_USERNAME"`
	Pass string `mapstructure:"DB_PASSWORD"`
	Name string `mapstructure:"DB_NAME"`
}

var DBConfig *DatabaseConfig

func InitializeEnv() {
	viper.AddConfigPath("..")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&DBConfig); err != nil {
		panic(err)
	}
}
