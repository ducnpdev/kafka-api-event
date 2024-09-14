package kafka

import (
	"event-tracking/config"

	"github.com/segmentio/kafka-go"
)

func NewKafkaReader(k config.KafkaReader) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: k.BrokerAddress,
		Topic:   k.Topic,
		GroupID: k.GroupID,
	})
}

func NewKafkaWriter(k config.KafkaWriter) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers: k.BrokerAddress,
		Topic:   k.Topic,
	})
}
