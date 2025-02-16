package env

import (
	"errors"
	"fmt"
	"net"
	"os"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/config"
)

const (
	httpHostAccessEnvName = "ACCESS_HOST"
	httpPortEnvName = "HTTP_PORT"
)

type httpConfig struct {
	host	string
	port	string
}

func NewHTTPConfig() (config.HTTPConfig, error) {
	host := os.Getenv(httpHostAccessEnvName)
	if len(host) == 0 {
		return nil, errors.New(fmt.Sprintf(config.ErrorMessage, httpHostAccessEnvName))
	}
	
	port := os.Getenv(httpPortEnvName)
	if len(port) == 0 {
		return nil, errors.New(fmt.Sprintf(config.ErrorMessage, httpPortEnvName))
	}
	
	return &httpConfig{
		host: host,
		port: port,
	}, nil
}

func (c *httpConfig) Port() string {
	return fmt.Sprintf(":%s", c.port)
}

func (c *httpConfig) AccessAddress() string {
	return net.JoinHostPort(c.host, c.port)
}