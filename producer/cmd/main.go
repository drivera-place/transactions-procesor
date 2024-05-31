package main

import (
	"domain/pkg/domain"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"producer/pkg/persistence/implementation"
	"producer/pkg/persistence/interfaces"
	"strconv"
	"time"
)

func main() {

	filePath := "txns.csv"
	var db interfaces.TxnsDB = &implementation.DynamoDB{}
	var txn domain.Transaction

	
	date, _ := time.Parse("2024-Jan-01", "2024-May-15")
	_, err := db.Save(txn.New(5, date, 1500.00))

	if err != nil {
		fmt.Errorf("Could not save row, v%", err)
	}

	rows, err := readTnxs(filePath)

	fmt.Println(rows)
}

func readTnxs(filePath string) ([]*domain.Transaction, error) {
	var rows []*domain.Transaction
	var txn domain.Transaction

	if _, err := os.Stat(filePath); err != nil {
		fmt.Println("File does not exist")
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error while reading the file.")
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','

	_, err = reader.Read()
	if err != nil && err != io.EOF {
		return nil, err
	}

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		id, err := strconv.Atoi(row[0])
		if err != nil {
			return nil, err
		}

		l := "05/19/0000"
		date, err := time.Parse(l, row[1])

		if err != nil {
			return nil, err
		}

		tnx, err := strconv.ParseFloat(row[2], 64)

		if err != nil {
			return nil, err
		}

		rows = append(rows, txn.New(id, date, tnx))
	}

	return rows, nil
}
