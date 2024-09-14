package usecase

import (
	"context"

	"event-tracking/config"
	"event-tracking/internal/dto"
	"event-tracking/internal/facade"
)

type EventUseCase interface {
	EventReceiver(ctx context.Context, msg dto.EventReqDTO) error
}

type eventUseCase struct {
	cfg    *config.Config
	event  facade.EventFacade
	status int32
}

func NewEventUseCase(
	cfg *config.Config,
	e facade.EventFacade,
) EventUseCase {
	return &eventUseCase{
		cfg:    cfg,
		event:  e,
		status: 1,
	}
}
