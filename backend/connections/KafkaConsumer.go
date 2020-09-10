package connections

import "github.com/Shopify/sarama"

//IConsumer - Kafka consumer Services
type IConsumer interface {
	CreateConsumer()
	Setup(sarama.ConsumerGroupSession) error
	Cleanup(sarama.ConsumerGroupSession) error
	ConsumeClaim(sarama.ConsumerGroupSession, sarama.ConsumerGroupClaim) error
}

//AbstractConsumer - Abstract Kafka Consumer
type AbstractConsumer struct {
	IConsumer
}
