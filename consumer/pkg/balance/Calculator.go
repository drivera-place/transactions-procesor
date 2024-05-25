package balance

import (
	"fmt"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"transactions/pkg/imp"
)

type Calculator struct {
	Summary Summary
}

func (b *Calculator) Query() []imp.Row {
	tnxs := []imp.Row{}

	return tnxs
}

func prepareQuery() error {
	sess := getCredentials()
	svc := dynamodb.New(sess)

	var queryInput = &dynamodb.QueryInput{
		TableName: aws.String("Transactions"),
		//IndexName: aws.String("Date"),
		KeyConditions: map[string]*dynamodb.Condition{
			"Date": {
				ComparisonOperator: aws.String("CONTAINS"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String("-01-"),
					},
				},
			},
			"date": {
				ComparisonOperator: aws.String("BEGINS_WITH"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String("2024"),
					},
				},
			},
		},
	}

	resp, err := svc.Query(queryInput)

	if err != nil {
		log.Fatalf("Got error calling Query: %s", err)
		return err
	}

	log.Println("Query successfully executed")

	for _, row := range resp {
		fmt.Println(row)
	}

	return nil
}

func getCredentials() client.ConfigProvider {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return sess
}
