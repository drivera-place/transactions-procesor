package main

import (
	"fmt"
	"transactions/pkg/imp"
	"transactions/pkg/interfaces"
)

func main() {
	var builder interfaces.Builder
	lines := 1000
	fileName := "txns.csv"
	builder = &imp.BasicBuilder{Lines: lines, Path: fileName}

	fmt.Println("Creating file...")
	_, err := builder.Create(lines)

	if err != nil {
		fmt.Println(err)
	}
}
