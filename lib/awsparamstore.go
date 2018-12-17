package lib

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"strings"
)

func getSession(awsRegion string) (*session.Session, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String(awsRegion)},
		SharedConfigState: session.SharedConfigEnable,
	})

	if err != nil {
		return nil, err
	}

	return sess, nil
}

// GetAwsParamStoreData retrieves everything from parameter store based on path and region
func GetAwsParamStoreData(awsParamStorePath string, awsRegion string) (map[string]string, error) {
	sess, err := getSession(awsRegion)

	if err != nil {
		return nil, err
	}

	ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion(awsRegion))
	withDecryption := true

	configMap := make(map[string]string)
	param, err := ssmsvc.GetParametersByPath(&ssm.GetParametersByPathInput{Path: &awsParamStorePath, WithDecryption: &withDecryption})

	for param.NextToken != nil {
		if err != nil {
			return nil, err
		}

		for _, element := range param.Parameters {
			var key = (*element.Name)[strings.LastIndex(*element.Name, "/")+1 : len(*element.Name)]
			configMap[key] = *element.Value
		}

		param, err = ssmsvc.GetParametersByPath(&ssm.GetParametersByPathInput{Path: &awsParamStorePath, WithDecryption: &withDecryption, NextToken: param.NextToken})
	}

	return configMap, nil
}
