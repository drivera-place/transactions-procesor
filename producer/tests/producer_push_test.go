package tests

import (
	"fmt"
	"producer/pkg/imp2"
	"producer/pkg/interfaces"
	"testing"
	"time"
	"transactions/pkg/imp"
)

func TestPushTransaction(t *testing.T) {
	var db interfaces.Producer
	db = &imp2.DBProducer{}

	tm, _ := time.Parse("2024-Jan-01", "2024-May-15")
	err := db.Push(&imp.Row{Id: 5, Date: tm, Transaction: 1500.00})

	if err != nil {
		t.Fatalf("Could not save transaction.")
		fmt.Println(err)
	}

}
