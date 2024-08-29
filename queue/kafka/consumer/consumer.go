package consumer

import (
	"context"
	"log/slog"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer interface {
	ConsumeMessage(ctx context.Context, handler func(message []byte)) error
	Close() error
}
type KafkaConsumerImpl struct {
	reader *kafka.Reader
	logger *slog.Logger
}

func NewKafkaConsumer(brokers []string, topic string, groupId string, logger *slog.Logger) KafkaConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  groupId,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})

	return &KafkaConsumerImpl{
		reader: reader,
		logger: logger,
	}
}

func (c *KafkaConsumerImpl) ConsumeMessage(ctx context.Context, handler func(message []byte)) error {
	defer c.Close()
	for {
		select {
		case <-ctx.Done():
			c.logger.Info("Consumer shutting down")
			return nil
		default:
			m, err := c.reader.ReadMessage(ctx)
			if err != nil {
				c.logger.Error("Error reading message", "error", err)
				continue
			}
			go handler(m.Value)
		}

	}
}

func (c *KafkaConsumerImpl) Close() error {
	return c.reader.Close()
}
