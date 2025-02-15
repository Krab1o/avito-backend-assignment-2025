package env

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/config"
)

const (
	jwtSecretEnvName = "JWT_SECRET"
	jwtTimeoutEnvName = "JWT_TIMEOUT"
	jwtTimeoutParseError = "Unable to parse JWT timeout"
)
type jwtConfig struct {
	jwtSecret	[]byte
	timeout		int
}

func NewJWTConfig() (config.JWTConfig, error) {
	jwt := os.Getenv(jwtSecretEnvName)
	if len(jwtSecretEnvName) == 0 {
		return nil, errors.New(fmt.Sprintf(config.ErrorMessage, pgHostEnvName))
	}
	time := os.Getenv(jwtTimeoutEnvName)
	if len(jwtSecretEnvName) == 0 {
		return nil, errors.New(fmt.Sprintf(config.ErrorMessage, pgHostEnvName))
	}
	timeVal, err := strconv.Atoi(time)
	if err != nil {
		return nil, errors.New(jwtTimeoutParseError)
	}
	return &jwtConfig{
		jwtSecret: []byte(jwt),
		timeout: timeVal,
	}, nil
}

func (c *jwtConfig) Secret() []byte {
	return c.jwtSecret
}

func (c *jwtConfig) Timeout() int {
	return c.timeout
}