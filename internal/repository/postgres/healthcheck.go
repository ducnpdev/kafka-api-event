package postgres

import (
	"context"

	"gorm.io/gorm"
)

type HealthCheckRepository interface {
	Ping(ctx context.Context) error
}

type healthCheckRepository struct {
	db *gorm.DB
}

func NewHealthCheckRepository(db *gorm.DB) HealthCheckRepository {
	return &healthCheckRepository{db: db}
}

func (r *healthCheckRepository) Ping(ctx context.Context) error {
	var result int
	err := r.db.WithContext(ctx).Raw("SELECT 1").Scan(&result).Error
	return err
}
