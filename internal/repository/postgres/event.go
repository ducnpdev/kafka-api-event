package postgres

import (
	"context"
	"event-tracking/internal/repository/postgres/model"

	"gorm.io/gorm"
)

type eventRepository struct {
	db *gorm.DB
}
type EventRepository interface {
	CreateEvent(ctx context.Context, event *model.Event) error
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db: db}
}

// CreateEvent inserts a new event into the database
func (r *eventRepository) CreateEvent(ctx context.Context, event *model.Event) error {
	if err := r.db.WithContext(ctx).Table("schema.events").Create(event).Error; err != nil {
		return err
	}
	return nil
}
