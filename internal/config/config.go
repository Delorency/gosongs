package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type ConfigHTTPServer struct {
	Host string `env:"HOST" env-default:"localhost"`
	Port string `env:"PORT" env-default:"8080"`
}

type ConfigDatabase struct {
	Role string `env:"DB_ROLE"`
	Pass string `env:"DB_PASS"`
	Host string `env:"DB_HOST"`
	Port string `env:"DB_PORT"`
	Name string `env:"DB_NAME"`
}

type Config struct {
	HTTPServer ConfigHTTPServer
	Db         ConfigDatabase
}

func MustLoad() *Config {
	godotenv.Load()

	var cfgHttpServer ConfigHTTPServer
	var cfgDatabase ConfigDatabase

	if err := cleanenv.ReadEnv(&cfgHttpServer); err != nil {
		panic("Must be implemented")
	}
	if err := cleanenv.ReadEnv(&cfgDatabase); err != nil {
		panic("Must be implemented")
	}

	return &Config{cfgHttpServer, cfgDatabase}
}
