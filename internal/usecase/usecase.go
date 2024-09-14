package usecase

import (
	"context"
	"event-tracking/config"

	"event-tracking/internal/repository"
)

type UseCase struct {
	HealthCheckCase HealthCheckUsecase
	EventUseCase    EventUseCase
}

func InitUseCase(
	ctx context.Context,
	cfg *config.Config,
	repository repository.Repository,
	kafkaRepository repository.KafkaRepository,
) *UseCase {
	return &UseCase{}
}
