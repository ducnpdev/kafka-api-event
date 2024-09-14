package dto

type CommonReqDTO struct {
	ChannelId   string `json:"channelId"`
	RequestId   string `json:"requestId" validate:"required,max=50,min=3"`
	RequestTime string `json:"requestTime" validate:"required,max=50,min=3"`
	Signature   string `json:"signature" validate:"required,max=100,min=5"`
}

type CommonRespDTO struct {
	RequestId       string `json:"requestId"`
	ResponseCode    string `json:"responseCode"`
	ResponseTime    string `json:"responseTime"`
	ResponseMessage string `json:"responseMessage"`
}
