package aws

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/service/ssm"
)

type ParameterStore struct {
	SecretKey       string `json:"secretKey"`
	Url             string `json:"url"`
	KeyEAS          string `json:"keyEAS"`
	HMacKey         string `json:"HMacKey"`
	DynamoTableName string `json:"dynamoTableName"`
}

func GetParameter(name *string) (results *ssm.GetParameterOutput, err error) {
	session := GetSession()
	if name == nil {
		return results, fmt.Errorf("get parameter required field name is required")
	}
	return ssm.New(session).GetParameter(&ssm.GetParameterInput{
		Name: name,
	})
}

// get aws parameter-store
func GetSecretKey(ctx context.Context, key string) (ParameterStore, error) {
	var (
		param = ParameterStore{}
		err   error
	)
	paramOut, err := GetParameter(&key)
	if err != nil {
		return param, err
	}
	value := *paramOut.Parameter.Value
	if value == "" {
		return param, fmt.Errorf("value of parameter %s value empty", key)
	}
	err = json.Unmarshal([]byte(value), &param)
	if err != nil {
		return param, fmt.Errorf("json Unmarshal of partner %s, err %s", key, err)
	}

	return param, err
}
