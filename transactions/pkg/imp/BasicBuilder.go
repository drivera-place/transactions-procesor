package imp

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type BasicBuilder struct {
	Lines int
	Name  string
	Path  string
	Rows  []Row
}

func (b *BasicBuilder) Create(lines int) (string, error) {

	w, err := os.Create(b.Path)

	if err != nil {
		panic(err)
	}
	defer w.Close()

	n := csv.NewWriter(w)
	err = n.Write([]string{"Id", "Date", "Transaction"})

	if err != nil {
		return "", err
	}

	b.createTransactions(lines)

	for _, row := range b.Rows {

		err := n.Write(formatRow(row))

		if err != nil {
			return "", err
		}
	}

	n.Flush()
	return b.Path, n.Error()
}

func (b *BasicBuilder) createTransactions(lines int) {

	b.Rows = make([]Row, lines)

	for i := 0; i < lines; i++ {
		b.Rows[i] = Row{Id: i, Date: randomDate(), Transaction: randomAmmount()}
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

func formatRow(r Row) []string {
	s := []string{}

	s = append(s, strconv.Itoa(r.Id))
	s = append(s, strconv.Itoa(r.Date.Day())+"/"+strconv.Itoa(int(r.Date.Month())))
	s = append(s, fmt.Sprintf("%+.2f", r.Transaction))

	return s
}
