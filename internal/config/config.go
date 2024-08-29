package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env      string        `yaml:"env" env-default:"local"`
	PathDb   string        `yaml:"db_path" env-required:"true"`
	TokenTTL time.Duration `yaml:"token_ttl" env-default:"1h"`
	Grpc     *GRPC         `yaml:"grpc"`
}

type GRPC struct {
	Port    string        `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {
	confPath := fetchPath()
	cfg := &Config{}

	if err := cleanenv.ReadConfig(confPath, cfg); err != nil {
		panic("Failed to read config: " + err.Error())
	}

	return cfg
}

func fetchPath() string {
	var res string

	flag.StringVar(&res, "config", "/home/katia/GolandProjects/PizzaOrderApp(gRPC)/cmd/pizza-order/config/config.yaml", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
