package model

import "gorm.io/datatypes"

type Event struct {
	RequestID  string            `gorm:"column:request_id" json:"requestId"`
	Partner    string            `gorm:"column:partner" json:"partner"`
	Channel    string            `gorm:"column:channel" json:"channel"`
	AppVersion string            `gorm:"column:app_version" json:"appVersion"`
	UserAgent  string            `gorm:"column:user_agent" json:"userAgent"`
	User       string            `gorm:"column:user" json:"user"`
	SdkVersion string            `gorm:"column:sdk_version" json:"sdkVersion"`
	Action     string            `gorm:"column:action" json:"action"`
	EventType  string            `gorm:"column:event_type" json:"eventType"`
	Time       string            `gorm:"column:time" json:"time"`
	SourceIP   string            `gorm:"column:source_ip" json:"sourceIp"`
	SessionID  string            `gorm:"column:session_id" json:"sessionId"`
	DeviceID   string            `gorm:"column:device_id" json:"deviceId"`
	Meta       datatypes.JSONMap `gorm:"column:meta" json:"meta"`
}
