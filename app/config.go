package app

// Environment ..
type Environment string

// Development ..
const (
	Development Environment = "dev"
	Production  Environment = "prod"
)

// Config ..
type Config struct {
	HTTP        HTTPConfig
	MySQL       MySQLConfig
	Redis       RedisConfig
	Logger      LoggerConfig
	Environment Environment
}

// HTTPConfig ..
type HTTPConfig struct {
	Host string
	Port string
}

// MySQLConfig .
type MySQLConfig struct {
	User     string
	Password string
	Host     string
	Database string
}

// RedisConfig .
type RedisConfig struct {
	Host      string
	Port      string
	MaxIdle   int
	MaxActive int
}

// LoggerConfig .
type LoggerConfig struct {
	Path string
}
