package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var AppConfig Config

type Config struct {
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbName     string `mapstructure:"DB_NAME"`
	DbSslMode  string `mapstructure:"DB_SSL_MODE"`
	DbTimeZone string `mapstructure:"DB_TIMEZONE"`
	AppPort    string `mapstructure:"APP_PORT"`
}

func (c Config) GetDBURL() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s TimeZone=%s",
		c.DbHost, c.DbPort, c.DbUser, c.DbName, c.DbPassword, c.DbSslMode, c.DbTimeZone)
}

func LoadConfig(path string) (config Config) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.ReadInConfig()
	viper.Unmarshal(&config)
	AppConfig = config
	return AppConfig
}
