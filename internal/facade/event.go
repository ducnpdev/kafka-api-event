package facade

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"event-tracking/config"
	"event-tracking/internal/dto"
	kafkaRepo "event-tracking/internal/repository/kafka"
	"event-tracking/pkg/logger"

	"github.com/segmentio/kafka-go"
)

type EventFacade interface {
	FetchMessage(ctx context.Context, reqDTO dto.FetchMessageDTO) (*kafka.Message, error)
	CommitMessages(ctx context.Context, reqDTO dto.CommitMessageDTO, msg *kafka.Message) error

	PushMessage(ctx context.Context, msg kafka.Message) error
}

func NewEventFacade(
	cfg *config.Config,
	kafKaReader kafkaRepo.ReaderRepository,
	kafkaWriter kafkaRepo.WriteMsgRepository,
) EventFacade {
	return &eventFacade{
		cfg:         cfg,
		kafKaReader: kafKaReader,
		kafkaWriter: kafkaWriter,
	}
}

type eventFacade struct {
	cfg         *config.Config
	kafKaReader kafkaRepo.ReaderRepository
	kafkaWriter kafkaRepo.WriteMsgRepository
}

func (c *eventFacade) FetchMessage(ctx context.Context, reqDTO dto.FetchMessageDTO) (*kafka.Message, error) {
	var (
		fieldWorker  = zap.String("workerID", fmt.Sprintf("%d", reqDTO.WorkerID))
		consumerName = logger.GetFieldsKafkaMessageType(reqDTO.ConsumerName)
	)
	logger.GLogger.InfoWithField("facade fetch message started", fieldWorker, consumerName)
	defer func() {
		logger.GLogger.InfoWithField("facade fetch message stopped", fieldWorker, consumerName)
	}()
	if c.kafKaReader == nil {
		return nil, fmt.Errorf("facade fetch message msk nil, please init it's")
	}
	return c.kafKaReader.FetchMessage(ctx)

}

func (c *eventFacade) CommitMessages(ctx context.Context, reqDTO dto.CommitMessageDTO, msg *kafka.Message) error {
	var (
		fieldWorker = zap.String("workerID", fmt.Sprintf("%d", reqDTO.WorkerID))
	)
	logger.GLogger.InfoWithField("facade commit message started", fieldWorker)
	defer func() {
		logger.GLogger.InfoWithField("facade commit message stopped", fieldWorker)
	}()
	if c.kafKaReader == nil {
		return fmt.Errorf("facade commit message msk nil, please init it's")
	}
	return c.kafKaReader.CommitMessages(ctx, msg)
}

func (c *eventFacade) PushMessage(ctx context.Context, msg kafka.Message) error {
	err := c.kafkaWriter.WriteMessage(ctx, msg)
	return err
}
