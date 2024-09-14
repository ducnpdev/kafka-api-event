package aws

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetSession() *session.Session {
	s := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config:            GetConfig(),
	}))
	return s
}
func DynamoClient(s *session.Session) *dynamodb.DynamoDB {
	if s == nil {
		s = GetSession()
	}
	return dynamodb.New(s)
}

func GetConfig() aws.Config {
	region := GetRegion()
	config := aws.NewConfig()
	config.Region = aws.String(region)
	return *config
}
func GetRegion() string {
	region := os.Getenv("REGION")
	return region
}
