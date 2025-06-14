package setting

type Config struct {
	Server     ServerSetting     `mapstructure:"server"`
	Postgresql PostgresQLSetting `mapstructure:"postgresql"`
	Logger     LoggerSetting     `mapstructure:"logger"`
	Redis      RedisSetting      `mapstructure:"redis"`
	RabbitMQ   RabbitMQSetting   `mapstructure:"rabbitmq"`
	JWT        JWTSetting        `mapstructure:"jwt"`
}

type ServerSetting struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type RedisSetting struct {
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	Password   string `mapstructure:"password"`
	Database   int    `mapstructure:"database"`
	ExpireTime int    `mapstructure:"expire_time"`
}

type PostgresQLSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

type LoggerSetting struct {
	Log_level     string `mapstructure:"log_level"`
	File_log_name string `mapstructure:"file_log_name"`
	Max_backups   int    `mapstructure:"max_backups"`
	Max_age       int    `mapstructure:"max_age"`
	Max_size      int    `mapstructure:"max_size"`
	Compress      bool   `mapstructure:"compress"`
}

type JWTSetting struct {
	TOKEN_HOUR_LIFESPAN uint   `mapstructure:"TOKEN_HOUR_LIFESPAN"`
	API_SECRET_KEY      string `mapstructure:"API_SECRET_KEY"`
	JWT_EXPIRATION      string `mapstructure:"JWT_EXPIRATION"`
}

type RabbitMQSetting struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
}