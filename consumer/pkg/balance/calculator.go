package balance

import (
	"domain/pkg/domain"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Calculator struct {
	Summary domain.Summary
}

func (b *Calculator) Query() []*domain.Transaction {
	var tnxs []*domain.Transaction

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

	for _, row := range resp.Items {
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
