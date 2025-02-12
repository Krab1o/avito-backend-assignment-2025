package config

type HTTPConfig interface {
	Address() string
}

type PGConfig interface {
	DSN() string
}