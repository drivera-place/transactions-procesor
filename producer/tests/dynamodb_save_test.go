package tests

import (
	"domain/pkg/domain"
	"producer/pkg/persistence/implementation"
	"producer/pkg/persistence/interfaces"
	"testing"
	"time"
)

func TestSaveTransaction(t *testing.T) {

	// Arrange
	var db interfaces.TxnsDB = &implementation.DynamoDB{}
	date, _ := time.Parse("2024-January-01", "2024-May-15")
	var txn domain.Transaction

	// Act
	id, err := db.Save(txn.New(24, date, 1500.00))

	// Assert
	t.Logf("Save Txn row with id: %v", id)
	if err != nil {
		t.Fatalf("Could not save transaction. %v", err)
	}
}

/* func TestSaveBulkTnxs(t *testing.T) {

	filePath := "tnxs.csv"
	f,err := os.Open(filePath);

	if  err != nil {
		panic(err)
	}

	defer f.Close()

	r := csv.NewReader(f)
	data, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	for i, line := range data {
		if i > 0 {
			for j, field := range line {

			}
		}
	}
} */
