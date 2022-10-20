package config

import (
	"encoding/base64"
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

type ServiceConf struct {
	TableName string `json:"table_name"`
}

func Get() (*ServiceConf, error) {
	region := "eu-west-1"

	//Create a Secrets Manager client
	svc := secretsmanager.New(session.New(),
		aws.NewConfig().WithRegion(region))
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String("example-api"),
		VersionStage: aws.String("AWSCURRENT"),
	}
	result, err := svc.GetSecretValue(input)
	if err != nil {
		return nil, err
	}
	// Decrypts secret using the associated KMS CMK.
	// Depending on whether the secret is a string or binary, one of these fields will be populated.
	var secretString, decodedBinarySecret string
	var conf = ServiceConf{}
	if result.SecretString != nil {
		secretString = *result.SecretString
		err := json.Unmarshal([]byte(secretString), &conf)
		if err != nil {
			return nil, err
		}
	} else {
		decodedBinarySecretBytes := make([]byte, base64.StdEncoding.DecodedLen(len(result.SecretBinary)))
		len, err := base64.StdEncoding.Decode(decodedBinarySecretBytes, result.SecretBinary)
		if err != nil {
			return nil, err
		}
		decodedBinarySecret = string(decodedBinarySecretBytes[:len])
		err = json.Unmarshal([]byte(decodedBinarySecret), &conf)
		if err != nil {
			return nil, err
		}
	}
	return &conf, nil
}
