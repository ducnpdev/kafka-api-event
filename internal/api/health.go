package api

import (
	"event-tracking/internal/usecase"

	"event-tracking/config"

	ginContext "event-tracking/pkg/gin"

	"github.com/gin-gonic/gin"
)

type healthRouter struct {
	cfg               *config.Config
	heathCheckUseCase usecase.HealthCheckUsecase
}

func NewHealthHandler(cfg *config.Config,
	heathCheckUseCase usecase.HealthCheckUsecase) healthRouter {
	return healthRouter{
		cfg:               cfg,
		heathCheckUseCase: heathCheckUseCase,
	}
}

func (h *healthRouter) liveness() gin.HandlerFunc {
	return ginContext.WithContext(func(ctx *ginContext.ContextGin) {
	})
}

func (h *healthRouter) readiness() gin.HandlerFunc {
	return ginContext.WithContext(func(ctx *ginContext.ContextGin) {
	})
}
