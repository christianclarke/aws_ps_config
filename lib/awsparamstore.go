package lib

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"strings"
)

func getAwsSession(awsRegion string) (*session.Session, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String(awsRegion)},
		SharedConfigState: session.SharedConfigEnable,
	})

	if err != nil {
		return nil, err
	}

	return sess, nil
}

func getAwsSsmSession(awsRegion string) (*ssm.SSM, error) {
	session, err := getAwsSession(awsRegion)

	if err != nil {
		return nil, err
	}

	return ssm.New(session, aws.NewConfig().WithRegion(awsRegion)), nil
}

// GetAwsParamStoreData retrieves everything from parameter store based on path and region
func GetAwsParamStoreData(awsParamStorePath string, awsRegion string) (map[string]string, error) {
	awsSSMSession, err := getAwsSsmSession(awsRegion)

	if err != nil {
		return nil, err
	}

	withDecryption := true
	configMap := make(map[string]string)
	param, err := awsSSMSession.GetParametersByPath(&ssm.GetParametersByPathInput{Path: &awsParamStorePath, WithDecryption: &withDecryption})

	for param.NextToken != nil {
		if err != nil {
			return nil, err
		}

		for _, element := range param.Parameters {
			var key = (*element.Name)[strings.LastIndex(*element.Name, "/")+1 : len(*element.Name)]
			configMap[key] = *element.Value
		}

		param, err = awsSSMSession.GetParametersByPath(&ssm.GetParametersByPathInput{Path: &awsParamStorePath, WithDecryption: &withDecryption, NextToken: param.NextToken})
	}

	return configMap, nil
}
