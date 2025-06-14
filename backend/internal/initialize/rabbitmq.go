package initialize

import (
	"context"
	"fmt"
	"time"

	"github.com/QuanCters/backend/global"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

const (
	maxRabbitRetries      = 3
	initialBackoff  = 1 * time.Second
	backoffFactor   = 2 

	// main exchange and queue config
	exchangeName	= "registration.exchange"
  	queueName	= "registration.create"
  	routingKey	= "registration.create"
	mainExchangeType = "direct"

	// DLX config
  	exchangeNameDlx	= "registration.exchange.dlx"
	queueNameDlx = "registration.create.dlx"
  	routingKeyDlx	= "registration.create.dlx"
	dlxExchangeType = "direct"

	// TTL
	messageTTL = int32(10000)
)

func InitRabbitMQ(ctx context.Context) {
	conn, err := connectWithRetry(ctx, maxRabbitRetries)
	if err != nil {
		global.Logger.Fatal("RabbitMQ connection failed", zap.Error(err))
		return
	}
	
	global.Rmq = conn
	global.Logger.Info("RabbitMQ is ready 🚀")

	// create channel to declare resource
	ch, err := conn.Channel()
	if err != nil {
		global.Logger.Fatal("Failed to open a RabbitMQ channel", zap.Error(err))
		// attempt to close after connection failed 
		if closeErr := conn.Close(); closeErr != nil {
			global.Logger.Error("Failed to close RabbitMQ connection after channel opening error", zap.Error(closeErr))
		}
		return
	}
	defer ch.Close() // close channel after used

	// declare exchanges, queues, and bindings
	if err := declareInfrastructure(ch); err != nil {
		global.Logger.Fatal("RabbitMQ infrastructure declaration failed", zap.Error(err))
		// close if failed
		if closeErr := conn.Close(); closeErr != nil {
			global.Logger.Error("Failed to close RabbitMQ connection after infrastructure declaration error", zap.Error(closeErr))
		}
		return
	}
}

func connectWithRetry(ctx context.Context, retries int) (*amqp091.Connection, error) {
	cfg := global.Config.RabbitMQ
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		cfg.User, cfg.Password, cfg.Host, cfg.Port,
	)

	var (
		conn *amqp091.Connection
		err  error
	)

	backoff := initialBackoff
	for attempt := 0; attempt <= retries; attempt++ {
		// Cho phép huỷ bỏ bằng context
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}

		conn, err = amqp091.Dial(url)
		if err == nil {
			return conn, nil // Thành công
		}

		if attempt == retries {
			break
		}

		global.Logger.Warn("RabbitMQ dial failed, retrying...",
			zap.Int("attempt", attempt+1),
			zap.Error(err),
			zap.Duration("sleep", backoff),
		)

		select {
		case <-time.After(backoff):
			backoff *= backoffFactor
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}

	return nil, fmt.Errorf("dial to %s failed after %d attempts: %w", url, retries+1, err)
}

func declareInfrastructure(ch *amqp091.Channel) error {
	// a. DLX exchange 
	if err := ch.ExchangeDeclare(
		exchangeNameDlx,
		dlxExchangeType,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return fmt.Errorf("declare DLX exchange: %w", err)
	}
	// global.Logger.Info("Declared DLX exchange", zap.String("name", exchangeNameDlx), zap.String("type", dlxExchangeType))

	// b. DLX queue
	if _, err := ch.QueueDeclare(
		queueNameDlx,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return fmt.Errorf("declare DLX queue: %w", err)
	}
	// global.Logger.Info("Declared DLX queue", zap.String("name", queueNameDlx))

	// c. Bind queue <-> DLX exchange
	if err := ch.QueueBind(queueNameDlx, routingKeyDlx, exchangeNameDlx, false, nil); err != nil {
		return fmt.Errorf("failed to bind dead letter queue to DLX: %w", err)
	}

	// global.Logger.Info("Bound DLX queue to DLX exchange",
	// 	zap.String("queue", queueNameDlx),
	// 	zap.String("exchange", exchangeNameDlx),
	// 	zap.String("routing_key", routingKeyDlx),
	// )

	err := ch.ExchangeDeclare(
		exchangeName,     // name
		mainExchangeType, // type
		true,             // durable
		false,            // auto-deleted
		false,            // internal
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare main exchange '%s': %w", exchangeName, err)
	}
	// global.Logger.Info("Declared main exchange", zap.String("name", exchangeName), zap.String("type", mainExchangeType))

	args := amqp091.Table{
		"x-dead-letter-exchange": exchangeNameDlx,
		"x-dead-letter-routing-key": routingKeyDlx,
	}
	if messageTTL > 0 {
		args["x-message-ttl"] = messageTTL
	}

	_, err = ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		args,
	)

	if err != nil {
		global.Logger.Error("Failed to declare main queue", zap.Error(err))
		return err
	}

	// global.Logger.Info("Declared main queue", zap.String("name", queueName), zap.Any("arguments", args))

	err = ch.QueueBind(
		queueName,    // queue name
		routingKey,   // routing key
		exchangeName, // main exchange name
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to bind main queue '%s' to main exchange '%s' with key '%s': %w", queueName, exchangeName, routingKey, err)
	}
	// global.Logger.Info("Bound main queue to main exchange",
	// 	zap.String("queue", queueName),
	// 	zap.String("exchange", exchangeName),
	// 	zap.String("routing_key", routingKey),
	// )

	// global.Logger.Info("RabbitMQ infrastructure declaration completed successfully.")

	return nil
}