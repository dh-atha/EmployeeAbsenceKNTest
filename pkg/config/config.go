package config

import (
	"time"

	"github.com/spf13/viper"
)

var Configuration Config

type Config struct {
	Server struct {
		Mode            string        `mapstructure:"mode"`
		Port            int           `mapstructure:"port"`
		ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"`
		JWTSecret       string        `mapstructure:"jwt_secret"`
		TokenDuration   time.Duration `mapstructure:"token_duration"`
	} `mapstructure:"server"`

	Postgres struct {
		Host     string   `mapstructure:"host"`
		Port     int      `mapstructure:"port"`
		Database string   `mapstructure:"database"`
		Username string   `mapstructure:"username"`
		Password string   `mapstructure:"password"`
		Options  []string `mapstructure:"options"`
	} `mapstructure:"postgres"`

	Logger struct {
		Dir        string `mapstructure:"dir"`
		FileName   string `mapstructure:"file_name"`
		MaxBackups int    `mapstructure:"max_backups"`
		MaxSize    int    `mapstructure:"max_size"`
		MaxAge     int    `mapstructure:"max_age"`
		Compress   bool   `mapstructure:"compress"`
		LocalTime  bool   `mapstructure:"local_time"`
	} `mapstructure:"logger"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (err error) {
	viper.SetConfigFile(path)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	var config Config
	err = viper.Unmarshal(&config)
	Configuration = config
	return
}
