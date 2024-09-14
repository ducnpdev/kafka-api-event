package usecase

import (
	"context"
	"encoding/json"
	"event-tracking/internal/dto"
	"event-tracking/pkg/logger"
	"fmt"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

func (u *eventUseCase) buildMessage(message dto.EventReqDTO) (msg kafka.Message, err error) {
	msgByte, err := json.Marshal(&message)
	if err != nil {
		return msg, err
	}

	msg = kafka.Message{
		Value: msgByte,
	}

	return msg, nil
}
func (u *eventUseCase) EventReceiver(ctx context.Context, message dto.EventReqDTO) error {
	msg, err := u.buildMessage(message)
	if err != nil {
		logger.GVA_LOG.Error("write notification message: build kafka message error ", zap.Error(err))
		return fmt.Errorf("build message err %s", err)
	}
	err = u.event.PushMessage(ctx, msg)
	if err != nil {
		logger.GVA_LOG.Error("write notification message to kafka err", zap.Error(err), zap.String("trace", message.TraceId), zap.String("reqId", message.RequestId))
	}
	return err
}
