package config

const ErrorMessage = "%s is empty or not read"

type HTTPConfig interface {
	Port() string
	AccessAddress() string
}

type PGConfig interface {
	DSN() string
}

type JWTConfig interface {
	Secret() []byte
	Timeout() int
}