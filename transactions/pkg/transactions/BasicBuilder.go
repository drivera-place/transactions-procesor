package transactions

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
	"domain/pkg/domain"
)

type FileBuilder struct {
	FilePath string
	Rows     []*domain.Transaction
}

func (fb *FileBuilder) CreateFile(lines int) (string, error) {

	w, err := os.Create(fb.FilePath)

	if err != nil {
		panic(err)
	}
	defer w.Close()

	n := csv.NewWriter(w)
	err = n.Write([]string{"Id", "Date", "Transaction"})

	if err != nil {
		return "", err
	}

	fb.createTransactions(lines)

	for _, row := range fb.Rows {

		err := n.Write(formatRow(row))

		if err != nil {
			return "", err
		}
	}

	n.Flush()
	return fb.FilePath, n.Error()
}

func (fb *FileBuilder) createTransactions(lines int) {

	var txn domain.Transaction
	fb.Rows = make([]*domain.Transaction, lines)

	for i := 0; i < lines; i++ {
		fb.Rows[i] = txn.New(i, randomDate(), randomAmmount())
	}
}

func randomDate() time.Time {

	currentYear := time.Now().Year()
	nextYear := currentYear + 1

	min := time.Date(currentYear, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(nextYear, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min

	return time.Unix(sec, 0)
}

func randomAmmount() float64 {
	max := 10000.00
	min := -10000.00

	return rand.Float64()*(max-min) + min
}

func formatRow(r *domain.Transaction) []string {
	s := []string{}

	s = append(s, strconv.Itoa(r.Id()))
	s = append(s, strconv.Itoa(r.Date().Day())+"/"+strconv.Itoa(int(r.Date().Month())))
	s = append(s, fmt.Sprintf("%+.2f", r.Txn()))

	return s
}
