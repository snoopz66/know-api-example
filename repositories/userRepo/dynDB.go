package userRepo

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/snoopz66/know-api-example/models"
	"github.com/snoopz66/know-api-example/repositories/config"
)

type DynDB struct {
	DynDBConnection *dynamodb.DynamoDB
	SrvConf         *config.ServiceConf
}

func New(conf *config.ServiceConf, dynDBConnection *dynamodb.DynamoDB) *DynDB {
	return &DynDB{SrvConf: conf, DynDBConnection: dynDBConnection}
}

func (d *DynDB) CreateUser(user *models.User) error {
	item, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{Item: item, TableName: aws.String(d.SrvConf.TableName)}
	_, err = d.DynDBConnection.PutItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (d *DynDB) GetUser(uuid string) (*models.User, error) {
	queryInput := &dynamodb.QueryInput{
		TableName: aws.String(d.SrvConf.TableName),
		KeyConditions: map[string]*dynamodb.Condition{
			"uuid": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(uuid),
					},
				},
			},
		},
	}
	res, err := d.DynDBConnection.Query(queryInput)
	if err != nil {
		return nil, err
	}
	var user *models.User
	err = dynamodbattribute.UnmarshalMap(res.Items[0], &user)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshall the response from dynamodb")
	}
	return user, nil
}
