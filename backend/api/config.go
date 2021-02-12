package api

import "strconv"

var (
	DefaultEnv     = "day"
	DefaultHost    = "localhost"
	DefaultPort    = "3030"
	DefaultAPIPath = "/"
)

type Config struct {
	Env     string
	Host    string
	Port    int
	APIPath string
}

func NewConfig(env, host, apipath string, port int) *Config {
	return &Config{
		Env:     env,
		Host:    host,
		Port:    port,
		APIPath: apipath,
	}
}

func DefaultConfig() *Config {
	port, _ := strconv.Atoi(DefaultPort)
	return &Config{
		Env:     DefaultEnv,
		Host:    DefaultHost,
		Port:    port,
		APIPath: DefaultAPIPath,
	}
}
