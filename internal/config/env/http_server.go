package env

import (
	"errors"
	"fmt"
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

func NewHTTPConfig() (config.HTTPConfig, error) {
	host := os.Getenv(httpHostEnvName)
	if len(host) == 0 {
		return nil, errors.New(fmt.Sprintf(config.ErrorMessage, httpHostEnvName))
	}
	port := os.Getenv(httpPortEnvName)
	if len(port) == 0 {
		return nil, errors.New(fmt.Sprintf(config.ErrorMessage, httpPortEnvName))
	}
	
	return &httpConfig{
		Host: host,
		Port: port,
	}, nil
}

func (c *httpConfig) Address() string {
	return net.JoinHostPort(c.Host, c.Port)
}