package common

type ResponseCode string

const (
	EmptyStr = ""
)

var ResponseMessage = map[string]string{}

func (rc ResponseCode) Code() string {
	return string(rc)
}

func (rc ResponseCode) IsDefined() bool {
	_, ok := ResponseMessage[rc.Code()]
	return ok
}

func (rc ResponseCode) Message() string {
	if value, ok := ResponseMessage[rc.Code()]; ok {
		return value
	}
	return EmptyStr
}

func ParseError(err error) ResponseCode {
	return ResponseCode(err.Error())
}
