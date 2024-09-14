package api

import (
	"github.com/gin-gonic/gin"
)

func (h *healthRouter) HealthRouter(router *gin.RouterGroup) {
	router.GET("/liveness", h.liveness())
	router.GET("/readiness", h.readiness())

}

func (h *eventRouter) EventRouter(router *gin.RouterGroup) {
	router.POST("receiver", h.receiverEvent())
}
