package imp2

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

type DBProducer struct {
}

func (b *DBProducer) Push(r *imp.Row) error {

	err := insertTransaction(r)

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func insertTransaction(r *imp.Row) error {
	sess := getCredentials()
	svc := dynamodb.New(sess)

	ti, err := dynamodbattribute.MarshalMap(r)
	if err != nil {
		log.Fatalf("Got error marshalling new transaction item: %s", err)
	}

	tableName := "Transactions"

	input := &dynamodb.PutItemInput{Item: ti, TableName: aws.String(tableName)}

	_, err = svc.PutItem(input)

	if err != nil {
		log.Fatalf("Got error calling Put Item into DynamoDB: %s", err)
		return err
	}

	fmt.Println("Successfully added '" + strconv.Itoa(r.Id) + " to table " + tableName)

	return nil
}

func getCredentials() client.ConfigProvider {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return sess
}
