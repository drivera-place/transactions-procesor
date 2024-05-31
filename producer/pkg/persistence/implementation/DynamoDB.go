package implementation

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"domain/pkg/domain"
)

type DynamoDB struct {
}

func (db *DynamoDB) Save(txn *domain.Transaction) (string, error) {

	tableName := "Transactions"

	sess := getCredentials()
	svc := dynamodb.New(sess)

	item, err := dynamodbattribute.MarshalMap(txn)
	if err != nil {
		log.Fatalf("Got error marshalling new transaction item: %s", err)
		return "", err
	}

	fmt.Println("#####")
	fmt.Println(item)

	input := &dynamodb.PutItemInput{Item: item, TableName: aws.String(tableName)}
	output, err := svc.PutItem(input)

	if err != nil {
		log.Fatalf("Got error calling Put Item into DynamoDB: %s", err)
		return "", err
	}

	id := output.Attributes["Id"].N

	log.Printf("Successfully added %v, to table %v", id, tableName)

	return derefString(id), nil
}

func derefString(s *string) string {
	if s != nil {
		return *s
	}

	return ""
}

func getCredentials() client.ConfigProvider {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return sess
}
