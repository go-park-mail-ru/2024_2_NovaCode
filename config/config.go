package config

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	Logger LoggerConfig `yaml:"logger"`
}

type LoggerConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

type ServerConfig struct {
	Port           string        `yaml:"port"`
	ReadTimeout    time.Duration `yaml:"readTimeout"`
	WriteTimeout   time.Duration `yaml:"writeTimeout"`
	IdleTimeout    time.Duration `yaml:"idleTimeout"`
	ContextTimeout time.Duration `yaml:"contextTimeout"`
}

func New() (*Config, error) {
	viper, err := newViper()
	if err != nil {
		return nil, fmt.Errorf("cannot create config: %v", err)
	}

	cfg, err := parseConfig(viper)
	if err != nil {
		return nil, fmt.Errorf("cannot parse config: %v", err)
	}

	return cfg, nil
}

func newViper() (*viper.Viper, error) {
	v := viper.New()

	v.AddConfigPath(os.Getenv("CONFIG_PATH"))
	v.SetConfigName(os.Getenv("CONFIG_NAME"))
	v.SetConfigType(os.Getenv("CONFIG_TYPE"))

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}

		return nil, err
	}

	return v, nil
}

func parseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config

	err := v.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to decode config into struct: %v", err)
	}

	return &cfg, nil
}
