package config

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	HTTPServer HTTPServer `toml:"http-server"`
	Store      Store      `toml:"store"`
	Logger     Logger     `toml:"logger"`
}

type HTTPServer struct {
	BindAddr string `toml:"bind_addr"`
}

type Store struct {
	DatabaseURL string `toml:"database_url"`
}

type Logger struct {
	Level string `toml:"level"`
}

func MustLoad() *Config {
	var (
		configPath string
		config     Config
	)

	flag.StringVar(&configPath, "config-path", "config.toml", "path to config file")
	flag.Parse()
	_, err := toml.DecodeFile(configPath, &config)
	if err != nil {
		log.Fatal(err)
	}
	return &config
}
