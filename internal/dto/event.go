package dto

type EventReqDTO struct {
	RequestId  string                 `json:"requestId"`
	Partner    string                 `json:"partner"`
	Channel    string                 `json:"channel"`
	AppVersion string                 `json:"appVersion"`
	UserAgent  string                 `json:"userAgent"`
	User       string                 `json:"user"`
	SdkVersion string                 `json:"sdkVersion"`
	Action     string                 `json:"action"`
	EventType  string                 `json:"eventType"`
	Time       string                 `json:"time"`
	SourceIp   string                 `json:"sourceIp"`
	SessionId  string                 `json:"sessionId"`
	DeviceId   string                 `json:"deviceId"`
	Meta       map[string]interface{} `json:"meta"`
	TraceId    string
}
