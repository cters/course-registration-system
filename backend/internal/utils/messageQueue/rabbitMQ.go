package messageQueue

import (
	"context"

	"github.com/QuanCters/backend/global"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)



func PublishMessage(exchange string, queueName string, body []byte) error {
	// 1. create channel
	ch, err := global.Rmq.Channel()
	if err != nil {
		global.Logger.Error("Failed to open channel", zap.Error(err))
		return err
	}
	defer ch.Close()

	// 2. send message
	err = ch.PublishWithContext(
		context.Background(),
		exchange,
		queueName,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body: body,
			DeliveryMode: amqp091.Persistent,
		})
	
	if err != nil {
		global.Logger.Error("Failed to publish message", zap.Error(err))
		return err
	}

	return nil
}

