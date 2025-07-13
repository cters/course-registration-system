package setting

import "time"

type Config struct {
	DBUser     string `mapstructure:"DB_USER"`
	DBPass     string `mapstructure:"DB_PASSWORD"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	DBSSL      string `mapstructure:"DB_SSL"`
	ServerPort string `mapstructure:"SERVER_PORT"`

	JwtExpiration time.Duration `mapstructure:"JWT_EXPIRATION"`
	SecretKey     string    `mapstructure:"SECRET_KEY"`

	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     int    `mapstructure:"REDIS_PORT"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisDB       int    `mapstructure:"REDIS_DB"`
	RedisExpire   int    `mapstructure:"REDIS_EXPIRE_MINUTES"`

	LogLevel     string `mapstructure:"LOG_LEVEL"`
	LogFile      string `mapstructure:"LOG_FILE"`
	LogMaxSize   int    `mapstructure:"LOG_MAX_SIZE"`
	LogMaxBackup int    `mapstructure:"LOG_MAX_BACKUP"`
	LogMaxAge    int    `mapstructure:"LOG_MAX_AGE"`
	LogCompress  bool   `mapstructure:"LOG_COMPRESS"`
}