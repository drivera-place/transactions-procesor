package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"producer/pkg/imp2"
	"producer/pkg/interfaces"
	"strconv"
	"time"
	"transactions/pkg/imp"
)

func main() {

	filePath := "txns.csv"
	var db interfaces.Producer
	db = &imp2.DBProducer{}

	tm, _ := time.Parse("2024-Jan-01", "2024-May-15")
	err := db.Push(&imp.Row{Id: 5, Date: tm, Transaction: 1500.00})

	if err != nil {
		fmt.Println(err)
	}

	rows, err := readTnxs(filePath)

	fmt.Println(rows)
}

func readTnxs(filePath string) ([]imp.Row, error) {
	var rows []imp.Row

	if _, err := os.Stat(filePath); err != nil {
		fmt.Println("File does not exist\n")
	}

	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		fmt.Println("Error while reading the file.\n")
	}

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

		rows = append(rows, imp.Row{Id: id, Date: date, Transaction: tnx})
	}

	return rows, nil
}
