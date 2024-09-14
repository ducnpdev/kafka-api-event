package facade

import (
	"context"
	"event-tracking/config"
	"event-tracking/internal/repository/postgres"
)

type HealthCheckFacade interface {
	Ping(ctx context.Context) error
}

type healthCheckFacade struct {
	cfg             *config.Config
	healthCheckRepo postgres.HealthCheckRepository
}

func NewHealthCheckFacade(
	cfg *config.Config,
	healthCheckRepo postgres.HealthCheckRepository,
) HealthCheckFacade {
	return &healthCheckFacade{
		cfg:             cfg,
		healthCheckRepo: healthCheckRepo,
	}
}

func (u *healthCheckFacade) Ping(ctx context.Context) error {
	err := u.healthCheckRepo.Ping(ctx)
	if err != nil {
		return err
	}
	return nil
}
