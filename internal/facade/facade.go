package facade

import (
	"context"
	"event-tracking/config"
)

type Facade struct {
}

func InitFacade(
	ctx context.Context,
	cfg *config.Config,
) (*Facade, error) {
	return &Facade{}, nil
}
