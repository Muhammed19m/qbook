package app

import "flag"

type Config struct {
	HttpServer HttpServerConfig
}

func InitConfig() (Config, error) {
	var cfg Config
	flag.StringVar(&cfg.HttpServer.Addr, "addr", ":8080", "addr to listen on")
	flag.Parse()

	return cfg, nil
}
