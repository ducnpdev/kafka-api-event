package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type ReaderRepository interface {
	FetchMessage(ctx context.Context) (*kafka.Message, error)
	CommitMessages(ctx context.Context, m *kafka.Message) error
	// automatically commit the offset when called
	ReadMessage(ctx context.Context) (kafka.Message, error)
}

type readerRepository struct {
	reader *kafka.Reader
}

func NewReaderRepository(r *kafka.Reader) ReaderRepository {
	return &readerRepository{
		reader: r,
	}
}

func (r *readerRepository) FetchMessage(ctx context.Context) (*kafka.Message, error) {
	msg, err := r.reader.FetchMessage(ctx)
	return &msg, err
}

func (r *readerRepository) CommitMessages(ctx context.Context, m *kafka.Message) error {
	return r.reader.CommitMessages(ctx, *m)
}

func (r *readerRepository) ReadMessage(ctx context.Context) (kafka.Message, error) {
	return r.reader.ReadMessage(ctx)
}
