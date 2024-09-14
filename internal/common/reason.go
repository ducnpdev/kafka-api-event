package common

type ResponseCode string

const ()

var ResponseMessage = map[string]string{}

func (rc ResponseCode) Code() string {
	return string(rc)
}

func (rc ResponseCode) Message() string {
	if value, ok := ResponseMessage[rc.Code()]; ok {
		return value
	}
	return ""
}

func ParseError(err error) ResponseCode {
	return ResponseCode(err.Error())
}
