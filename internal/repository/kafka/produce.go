package kafka

import (
	"context"
	"event-tracking/config"

	"github.com/segmentio/kafka-go"
)

type WriteMsgRepository interface {
	WriteMessage(ctx context.Context, msg kafka.Message) error
}

type writeMsgRepository struct {
	cfg    *config.Config
	writer *kafka.Writer
}

func NewWriterRepository(cfg *config.Config, writer *kafka.Writer) WriteMsgRepository {
	return &writeMsgRepository{
		cfg:    cfg,
		writer: writer,
	}
}

func (r *writeMsgRepository) WriteMessage(ctx context.Context, msg kafka.Message) error {
	return r.writer.WriteMessages(ctx, msg)
}
