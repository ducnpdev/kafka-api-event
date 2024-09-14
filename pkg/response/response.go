package response

import (
	"encoding/json"
	"event-tracking/pkg/common"
	ginContext "event-tracking/pkg/gin"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Response struct {
	ResponseID      string      `json:"responseId"`
	ResponseTime    string      `json:"responseTime"`
	ResponseCode    string      `json:"responseCode"`
	ResponseMessage string      `json:"responseMessage"`
	Data            interface{} `json:"data,omitempty"`
}

func NewResponse(data interface{}, mgs, code string) Response {
	return Response{
		ResponseID:      uuid.New().String(),
		ResponseTime:    time.Now().Format(time.RFC3339),
		ResponseCode:    code,
		ResponseMessage: mgs,
		Data:            data,
	}
}

func NewResponseError(data interface{}, e error, mgs, code string) Response {
	if common.ParseError(e).IsDefined() {
		return Response{
			ResponseID:      uuid.New().String(),
			ResponseTime:    time.Now().Format(time.RFC3339),
			ResponseCode:    common.ParseError(e).Code(),
			ResponseMessage: common.ParseError(e).Message(),
			Data:            data,
		}
	}

	return Response{
		// TODO, get responseId is a requestId
		ResponseID:      uuid.New().String(),
		ResponseTime:    time.Now().Format(time.RFC3339),
		ResponseCode:    code,
		ResponseMessage: e.Error(),
		Data:            data,
	}
}

func JSONResponse(g *ginContext.ContextGin, statusCode int, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		Error(g, statusCode, fmt.Errorf("failed to marshal response: %w |  %+v", err, data))
		return
	}

	g.Data(statusCode, "application/json", b)
}

func Error(g *ginContext.ContextGin, statusCode int, err error) {
	g.JSON(statusCode, err)
}
