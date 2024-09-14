package aws

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/aws_msk_iam_v2"
)

type KafkaReaderProperty struct {
	Brokers []string
	Topic   string
	GroupID string
	Enable  bool
}

type KafkaWriterProperty struct {
	Brokers   []string
	Topic     string
	BatchSize int
}

func KafkaReader(p KafkaReaderProperty) *kafka.Reader {
	if !p.Enable {
		return nil
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}
	mechanism := aws_msk_iam_v2.NewMechanism(cfg)

	tmpConfig := kafka.ReaderConfig{
		Brokers: p.Brokers,
		GroupID: p.GroupID,
		Topic:   p.Topic,
		Dialer: &kafka.Dialer{
			Timeout:       10 * time.Second,
			DualStack:     true,
			SASLMechanism: mechanism,
			TLS:           &tls.Config{},
		},
	}
	return kafka.NewReader(tmpConfig)
}

func KafkaWriter(p KafkaWriterProperty) *kafka.Writer {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	mechanism := aws_msk_iam_v2.NewMechanism(cfg)
	w := &kafka.Writer{
		Addr:      kafka.TCP(p.Brokers...),
		Balancer:  &kafka.LeastBytes{},
		Topic:     p.Topic,
		BatchSize: p.BatchSize,
		Transport: &kafka.Transport{
			DialTimeout: 10 * time.Second,
			TLS:         &tls.Config{},
			SASL:        mechanism,
		},
	}

	return w
}
