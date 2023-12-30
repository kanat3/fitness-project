package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env            string `yaml:"env" env:"ENV" env-default:"local" env-required:"true"`
	HTTPServer     `yaml:"http_server"`
	DatabaseConfig `yaml:"database"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host" env:"Host" env-default:"localhost"`
	Port     string `yaml:"port" env:"Port" env-default:"5432"`
	User     string `yaml:"user" env:"User" env-required:"true"`
	Password string `yaml:"password" env:"Password" env-required:"true"`
	DBname   string `yaml:"dbname" env:"DBname" env-default:"postgresql"`
}

type HTTPServer struct {
	Address        string        `yaml:"address" env:"Address" env-default:"localhost:80"`
	Timeout        time.Duration `yaml:"timeout" env:"Timeout" timeout-defaul:"4s"`
	Middle_timeout time.Duration `yaml:"middle_timeout" env:"Middle_timeout" env-default:"60s"`
}

func GetConfig() *Config {
	/* not implemented in makefile */
	//configPath := os.Getenv("CONFIG_PATH")

	configPath := "config/local.yaml"

	if configPath == "" {
		log.Fatal("Env $CONFIG_PATH isn't set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("File in %s path doesn't exit", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Can't read file from %s", configPath)
	}

	return &cfg
}
