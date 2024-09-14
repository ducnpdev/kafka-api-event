package gin

import (
	"event-tracking/internal/common"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ContextGin struct {
	*gin.Context
}
type ResponseData struct {
	ResponseCode    string      `json:"responseCode,omitempty"`
	ResponseMessage string      `json:"responseMessage,omitempty"`
	ResponseId      string      `json:"responseId,omitempty"`
	ResponseTime    string      `json:"responseTime,omitempty"`
	Data            interface{} `json:"data,omitempty"`
}
type HandlerFunc func(ctx *ContextGin)

func WithContext(handler HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handler(&ContextGin{
			ctx,
		})
	}
}

func (c *ContextGin) BadRequestV2(err error, id string) {
	resp := ResponseData{}
	if err != nil {
		resp.ResponseCode = common.ParseError(err).Code()
		resp.ResponseMessage = common.ParseError(err).Message()
		resp.ResponseTime = time.Now().Format(time.RFC3339)
		resp.ResponseId = id
	}
	c.responseJson(http.StatusBadRequest, resp)
}

func (c *ContextGin) BadRequest(err error) {
	strErr := ""
	if err != nil {
		strErr = err.Error()
	}
	resp := ResponseData{
		ResponseMessage: strErr,
	}
	c.responseJson(http.StatusBadRequest, resp)
}

func (c *ContextGin) NotFound(err error) {
	c.responseJson(http.StatusNotFound, err)
}

func (c *ContextGin) BadLogic(err error) {
	resp := ResponseData{
		ResponseCode:    common.ParseError(err).Code(),
		ResponseMessage: common.ParseError(err).Message(),
		ResponseTime:    time.Now().Format("2006-01-02T15:04:05.000-07:00"),
	}
	c.responseJson(http.StatusOK, resp)
}

func (c *ContextGin) BadLogicV2(err error, id string) {
	resp := ResponseData{
		ResponseCode:    common.ParseError(err).Code(),
		ResponseMessage: common.ParseError(err).Message(),
		ResponseTime:    time.Now().Format("2006-01-02T15:04:05.000-07:00"),
		ResponseId:      id,
	}
	c.responseJson(http.StatusOK, resp)
}

func (c *ContextGin) OKResponse(data interface{}, id string) {
	resp := ResponseData{
		ResponseCode:    "00",
		ResponseMessage: "successfully",
		ResponseTime:    time.Now().Format("2006-01-02T15:04:05.000-07:00"),
		ResponseId:      id,
	}
	if data != nil {
		resp.Data = data
	}
	c.responseJson(http.StatusOK, resp)
}

func (c *ContextGin) responseJson(code int, data interface{}) {
	c.JSON(code, data)
	if code != http.StatusOK {
		c.Abort()
	}
}

func (c *ContextGin) TokenNotFound() {
}
