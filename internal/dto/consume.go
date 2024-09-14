package dto

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type MsgDTO struct {
	Msg         *kafka.Message
	TraceID     string
	WorkerID    uint8
	MessageType string
	ProcessTime string
}

type FetchMessageDTO struct {
	ConsumerName string
	WorkerID     uint8
}

type CommitMessageDTO struct {
	WorkerID uint8
}

type ConsumeDTO struct {
	NumberWorker     uint8
	KafkaMessageType string
	ProcessMsgFunc   func(context.Context, MsgDTO) error
}
