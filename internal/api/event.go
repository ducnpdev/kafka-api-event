package api

import (
	"event-tracking/internal/dto"
	"event-tracking/internal/usecase"
	ginContext "event-tracking/pkg/gin"
	"event-tracking/pkg/logger"
	"event-tracking/pkg/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type eventRouter struct {
	eventUseCase usecase.EventUseCase
}

func NewEventHandler(
	eventUseCase usecase.EventUseCase,
) eventRouter {
	return eventRouter{
		eventUseCase: eventUseCase,
	}
}

func (h *eventRouter) receiverEvent() gin.HandlerFunc {
	return ginContext.WithContext(func(ctx *ginContext.ContextGin) {
		var (
			reqDTO  = dto.EventReqDTO{}
			err     error
			traceId = utils.GenLogID()
		)

		if err = ctx.ShouldBindJSON(&reqDTO); err != nil {
			logger.GVA_LOG.Error("write notification message error while bind json", zap.Error(err), zap.String("trace", traceId))
			ctx.BadRequestV2(err, "")
			return
		}
		reqDTO.TraceId = traceId

		err = h.eventUseCase.EventReceiver(ctx, reqDTO)
		if err != nil {
			logger.GVA_LOG.Error("write notification message to kafka err", zap.Error(err), zap.String("trace", traceId), zap.String("reqId", reqDTO.RequestId))
			ctx.BadRequestV2(err, reqDTO.RequestId)
			return
		}
		ctx.OKResponse(nil, "")
	})
}
