package initialize

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/QuanCters/backend/global"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

var consumerStopChan chan struct {}

func InitConsumer() {
	consumerStopChan = make(chan struct{})
	go Consumer(1, consumerStopChan)
	go ConsumerDLX(consumerStopChan)
}

func ShutDownConsumers() {
	if consumerStopChan != nil {
		close(consumerStopChan)
		time.Sleep(2 * time.Second)
	}
}

// Consumer “sống” cho tới khi ctx huỷ hoặc tín hiệu hệ thống.
func Consumer(id int, stopChan chan struct{}) error {
	// 0. func test
	handler := func(deliveryTag uint64, msgBody string) error {
		global.Logger.Info("Received message",
			zap.Int("consumer_id", id),
			zap.Uint64("delivery_tag", deliveryTag),
			zap.String("body", msgBody),
		)

		// Mô phỏng thời gian xử lý
		time.Sleep(100 * time.Millisecond)

		// Mô phỏng lỗi xử lý để kiểm tra Nack và DLX
		if strings.Contains(msgBody, "error_dlx") {
			errMsg := fmt.Sprintf("simulated processing error for message: %s", msgBody)
			global.Logger.Warn(errMsg, zap.Int("consumer_id", id), zap.Uint64("delivery_tag", deliveryTag))
			return fmt.Errorf("%s", errMsg)
		}

		global.Logger.Info("Message processed successfully by handler",
			zap.Int("consumer_id", id),
			zap.Uint64("delivery_tag", deliveryTag),
		)
		return nil
	}

	// 1. mở channel
	ch, err := global.Rmq.Channel()
	if err != nil {
		global.Logger.Error("Open channel failed", zap.Error(err))
		return err
	}

	// đảm bảo channel đóng khi hàm kết thúc
	defer ch.Close()

	// 2. QoS – cho phép n message song song
	if err := ch.Qos(5, 0, false); err != nil {
		return err
	}

	// 3. đăng ký consumer
	consumerTag := fmt.Sprintf("consumer-%d", id)
	msgs, err := ch.Consume(
		queueName,
		consumerTag,
		false, // auto-ack
		false, // exclusive
		false, // no-local (deprecated)
		false, // no-wait
		nil,   // args
	)

	if err != nil {
		global.Logger.Error("Register consumer failed", zap.Error(err))
		return err
	}

	cancelChan := ch.NotifyCancel(make(chan string))
	chanClose  := ch.NotifyClose(make(chan *amqp091.Error))
    connClose  := global.Rmq.NotifyClose(make(chan *amqp091.Error))

	global.Logger.Info("Consumer registered OK",
		zap.Int("consumer_id", id),
		zap.String("queue", queueName),
		zap.String("tag", consumerTag),
	)

	for {
		select {

		case <-stopChan:
			s := fmt.Sprintf("Stopping consumer %d …", id)
			global.Logger.Info(s)
			return nil

		case cancel := <- cancelChan:
			global.Logger.Warn("Consumer cancelled by broker",
				zap.Int("consumer_id", id),
				zap.String("tag", cancel),
			)
			return fmt.Errorf("consumer cancelled: %s", cancel)

		case cerr := <-chanClose:
            global.Logger.Error("Channel closed", zap.Error(cerr))
            return cerr

		case cerr := <-connClose:
            global.Logger.Error("Connection closed", zap.Error(cerr))
            return cerr

		case d, ok := <-msgs:
			if !ok {
				global.Logger.Warn("Message channel closed. Consumer stopping...", zap.Int("consumer_id", id), zap.String("tag", d.ConsumerTag))
				return amqp091.ErrClosed
			}

			// processing message
			if err := handler(d.DeliveryTag, string(d.Body)); err != nil {
				// Xử lý tin nhắn thất bại -> Nack
				global.Logger.Error(
					"Handle message failed",
					zap.Int("consumer_id", id),
					zap.Uint64("delivery_tag",
					d.DeliveryTag),
					zap.Error(err),
					zap.Bool("requeue", false), // Sẽ không requeue, gửi tới DLX nếu có
				)
				// Gửi Nack, không requeue (tin nhắn sẽ đi vào DLX nếu được cấu hình, hoặc bị loại bỏ)
				if nackErr := d.Nack(false, false); nackErr != nil {
					global.Logger.Error("Failed to Nack message",
						zap.Int("consumer_id", id),
						zap.Uint64("delivery_tag", d.DeliveryTag),
						zap.Error(nackErr),
					)
				}
			} else {
				// Xử lý tin nhắn thành công -> Ack
				global.Logger.Info("Message handled successfully by handler, attempting to Ack",
					zap.Int("consumer_id", id),
					zap.Uint64("delivery_tag", d.DeliveryTag),
				)
				// Gửi Ack, xác nhận tin nhắn đã được xử lý
				if ackErr := d.Ack(false); ackErr != nil {
					global.Logger.Error("Failed to Ack message",
						zap.Int("consumer_id", id),
						zap.Uint64("delivery_tag", d.DeliveryTag),
						zap.Error(ackErr),
					)
				}
			}
		}
	}
}

func ConsumerDLX(stopChan chan struct{}) error {
	// 1. mở channel
	ch, err := global.Rmq.Channel()
	if err != nil {
		global.Logger.Error("Open channel failed", zap.Error(err))
		return err
	}
	// đảm bảo channel đóng khi hàm kết thúc
	defer ch.Close()

	msgs, err := ch.Consume(
		queueNameDlx,
		"dlx-consumer",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		global.Logger.Error("Failed to register a DLX consumer", zap.Error(err))
		return err
	}

	log.Printf("[*] DLX Consumer is waiting letter from %s", queueNameDlx)

	for {
		select {
		case <-stopChan:
			log.Println("DLX Consumer is stopping...")
			return nil
		case d, ok := <-msgs:
			if !ok {
				log.Println("DLX consumer is closed.")
				return nil
			}
			log.Printf("DLX Consumer received message: %s\n", d.Body)
		}
	}
}