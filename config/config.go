package config

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Service  ServiceConfig  `yaml:"service"`
	Postgres PostgresConfig `yaml:"postgres"`
	Minio    MinioConfig    `yaml:"minio"`
}

type ServiceConfig struct {
	Port           string        `yaml:"port"`
	ReadTimeout    time.Duration `yaml:"readTimeout"`
	WriteTimeout   time.Duration `yaml:"writeTimeout"`
	IdleTimeout    time.Duration `yaml:"idleTimeout"`
	ContextTimeout time.Duration `yaml:"contextTimeout"`

	CORS   CORSConfig   `yaml:"cors"`
	Logger LoggerConfig `yaml:"logger"`
	Auth   AuthConfig   `yaml:"auth"`
}

type CORSConfig struct {
	AllowOrigin      string `yaml:"allowOrigin"`
	AllowMethods     string `yaml:"allowMethods"`
	AllowHeaders     string `yaml:"allowHeaders"`
	ExposeHeaders    string `yaml:"exposeHeaders"`
	AllowCredentials bool   `yaml:"allowCredentials"`
}

type LoggerConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

type AuthConfig struct {
	CSRF CSRFConfig `yaml:"csrf"`
	Jwt  JwtConfig  `yaml:"jwt"`
}

type CSRFConfig struct {
	HeaderName string `yaml:"headerName"`
	Salt       string `yaml:"salt"`
}

type JwtConfig struct {
	Secret string          `yaml:"secret"`
	Expire time.Duration   `yaml:"expire"`
	Cookie JwtCookieConfig `yaml:"cookie"`
}

type JwtCookieConfig struct {
	Name     string `yaml:"name"`
	MaxAge   int    `yaml:"maxAge"`
	Secure   bool   `yaml:"secure"`
	HttpOnly bool   `yaml:"httpOnly"`
}

type PostgresConfig struct {
	Host                string `yaml:"host"`
	Port                string `yaml:"port"`
	User                string `yaml:"user"`
	Password            string `yaml:"password"`
	DBName              string `yaml:"dbName"`
	SSLMode             bool   `yaml:"sslMode"`
	Driver              string `yaml:"driver"`
	MaxOpenConns        int    `yaml:"maxOpenConns"`
	ConnMaxIdleLifetime int    `yaml:"connMaxLifetime"`
	MaxIdleConns        int    `yaml:"maxIdleConns"`
	ConnMaxIdleTime     int    `yaml:"connMaxIdleTime"`
}

type MinioConfig struct {
	InnerURL string `yaml:"innerURL"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	SSLMode  bool   `yaml:"sslMode"`
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
	v.SetConfigName(os.Getenv("ENV"))
	v.SetConfigType("yaml")

	err := bindEnv(v)
	if err != nil {
		return nil, fmt.Errorf("cannot bind env variables: %v", err)
	}

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf("config file not found")
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

func bindEnv(v *viper.Viper) error {
	envBindings := map[string]string{
		"postgres.port":           "POSTGRES_PORT",
		"postgres.dbname":         "POSTGRES_DB",
		"postgres.user":           "POSTGRES_USER",
		"postgres.password":       "POSTGRES_PASSWORD",
		"minio.user":              "MINIO_USER",
		"minio.password":          "MINIO_PASSWORD",
		"service.auth.csrf.salt":  "CSRF_SALT",
		"service.auth.jwt.secret": "JWT_SECRET",
	}

	for key, env := range envBindings {
		if err := v.BindEnv(key, env); err != nil {
			return err
		}
	}

	return nil
}
