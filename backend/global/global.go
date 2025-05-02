package global

import (
	"database/sql"

	"github.com/QuanCters/backend/pkg/logger"
	"github.com/QuanCters/backend/pkg/setting"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Rdb    *redis.Client
	Pdbc *sql.DB
	Rmq *amqp091.Connection
)