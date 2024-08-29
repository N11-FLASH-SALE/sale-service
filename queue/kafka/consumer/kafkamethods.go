package consumer

import (
	"context"
	"log/slog"
	"sale/service"
)

type KafkaMethods interface {
	UpdateProduct(ctx context.Context, topic string)
	DeleteProduct(ctx context.Context, topic string)
}

type KafkaMethodsImpl struct {
	brokers          []string
	msgBrokerService service.ProductKafkaService
	logger           *slog.Logger
}

func NewKafkaMethods(brokers []string, msgBrokerService service.ProductKafkaService, logger *slog.Logger) KafkaMethods {
	return &KafkaMethodsImpl{
		brokers:          brokers,
		msgBrokerService: msgBrokerService,
		logger:           logger,
	}
}

func (km *KafkaMethodsImpl) UpdateProduct(ctx context.Context, topic string) {
	reader := NewKafkaConsumer(km.brokers, topic, "", km.logger)
	defer reader.Close()

	km.logger.Info("Starting consumer for topic")

	err := reader.ConsumeMessage(ctx, km.msgBrokerService.UpdateProduct)
	if err != nil {
		km.logger.Error("Error consuming messages", "error", err)
		return
	}
}

func (km *KafkaMethodsImpl) DeleteProduct(ctx context.Context, topic string) {
	reader := NewKafkaConsumer(km.brokers, topic, "", km.logger)
	defer reader.Close()

	km.logger.Info("Starting consumer for topic")

	err := reader.ConsumeMessage(ctx, km.msgBrokerService.DeleteProduct)
	if err != nil {
		km.logger.Error("Error consuming messages", "error", err)
		return
	}
}
