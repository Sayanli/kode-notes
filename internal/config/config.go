package config

import (
	"log"
	"path"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App    `yaml:"app"`
		HTTP   `yaml:"http"`
		Log    `yaml:"log"`
		PG     `yaml:"postgres"`
		JWT    `yaml:"jwt"`
		Hasher `yaml:"hasher"`
	}

	App struct {
		Name    string `yaml:"name"    env:"APP_NAME"`
		Version string `yaml:"version" env:"APP_VERSION"`
	}

	HTTP struct {
		Host string `yaml:"host" env:"HTTP_HOST"`
		Port string `yaml:"port" env:"HTTP_PORT"`
	}

	Log struct {
		Level string `yaml:"level" env:"LOG_LEVEL"`
	}

	PG struct {
		MaxPoolSize int    `yaml:"max_pool_size" env:"PG_MAX_POOL_SIZE"`
		URL         string `yaml:"pg_url"`
	}

	JWT struct {
		SignKey  string        `yaml:"jwt_sign_key"`
		TokenTTL time.Duration `yaml:"jwt_token_ttl"`
	}

	Hasher struct {
		Salt string `yaml:"hasher_salt"`
	}
)

func NewConfig(configPath string) (*Config, error) {
	cfg := &Config{}
	if configPath == "" {
		log.Fatal("config path is not set")
	}

	err := cleanenv.ReadConfig(path.Join("./", configPath), cfg)
	if err != nil {
		log.Fatalf("error reading config file: %s", err)
	}

	return cfg, nil
}
