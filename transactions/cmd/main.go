package main

import (
	"fmt"
	"transactions/pkg/interfaces"
	"transactions/pkg/transactions"
)

func main() {
	var builder interfaces.Builder
	lines := 1000
	fileName := "txns.csv"
	builder = &transactions.FileBuilder{FilePath: fileName}

	fmt.Println("Creating file...")
	_, err := builder.CreateFile(lines)

	if err != nil {
		fmt.Println(err)
	}
}
