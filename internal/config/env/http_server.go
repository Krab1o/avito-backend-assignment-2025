package env

import (
	"net"
	"os"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/config"
)

const (
	httpHostEnvName = "HTTP_HOST"
	httpPortEnvName = "HTTP_PORT"
)

type httpConfig struct {
	Host	string
	Port	string
}

func NewHTTPConfig() config.HTTPConfig {
	host := os.Getenv(httpHostEnvName)
	port := os.Getenv(httpPortEnvName)
	
	return &httpConfig{
		Host: host,
		Port: port,
	}
}

func (c *httpConfig) Address() string {
	return net.JoinHostPort(c.Host, c.Port)
}