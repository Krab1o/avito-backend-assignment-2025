package env

import (
	"errors"
	"fmt"
	"os"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/config"
)

const jwtEnvName = "JWT_SECRET"

type jwtConfig struct {
	jwtSecret	string
}

func NewJWTConfig() (config.JWTConfig, error) {
	errorMessage := "%s is empty or not read"
	jwt := os.Getenv(jwtEnvName)
	if len(jwtEnvName) == 0 {
		return nil, errors.New(fmt.Sprintf(errorMessage, pgHostEnvName))
	}
	return &jwtConfig{
		jwtSecret: jwt,
	}, nil
}

func (c *jwtConfig) Secret() string {
	return c.jwtSecret
}