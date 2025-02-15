package env

import (
	"errors"
	"fmt"
	"os"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/config"
)

const (
	pgHostEnvName 		= "PG_HOST"
	pgPortEnvName 		= "PG_PORT"
	pgUserEnvName 		= "PG_USER"
	pgPasswordEnvName 	= "PG_PASSWORD"
	pgDatabaseEnvName 	= "PG_DB"
)

type pgConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
}

func NewPGConfig() (config.PGConfig, error) {
	host := os.Getenv(pgHostEnvName)
	if len(host) == 0 {
		return nil, errors.New(fmt.Sprintf(config.ErrorMessage, pgHostEnvName))
	}
	port := os.Getenv(pgPortEnvName)
	if len(host) == 0 {
		return nil, errors.New(fmt.Sprintf(config.ErrorMessage, pgPortEnvName))
	}
	user := os.Getenv(pgUserEnvName)
	if len(host) == 0 {
		return nil, errors.New(fmt.Sprintf(config.ErrorMessage, pgUserEnvName))
	}
	password := os.Getenv(pgPasswordEnvName)
	if len(host) == 0 {
		return nil, errors.New(fmt.Sprintf(config.ErrorMessage, pgPasswordEnvName))
	}
	dbname := os.Getenv(pgDatabaseEnvName)
	if len(host) == 0 {
		return nil, errors.New(fmt.Sprintf(config.ErrorMessage, pgDatabaseEnvName))
	}

	return &pgConfig{
		Host : host,
		Port : port,
		User : user,
		Password: password,
		DBName: dbname,
	}, nil
}

func (c *pgConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.DBName,
	)
}