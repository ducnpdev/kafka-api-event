package usecase

import (
	"context"
	"event-tracking/config"
	"event-tracking/internal/facade"
)

type HealthCheckUsecase interface {
	Ping(ctx context.Context) error
}

type healthCheckUsecase struct {
	cfg               *config.Config
	healthCheckFacade facade.HealthCheckFacade
}

func NewHealthCheckUsecase(
	cfg *config.Config,
	healthCheckFacade facade.HealthCheckFacade,
) HealthCheckUsecase {
	return &healthCheckUsecase{
		cfg:               cfg,
		healthCheckFacade: healthCheckFacade,
	}
}

func (u *healthCheckUsecase) Ping(ctx context.Context) error {
	err := u.healthCheckFacade.Ping(ctx)
	if err != nil {
		return err
	}
	return nil
}
