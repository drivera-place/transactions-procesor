package tests

import (
	"encoding/csv"
	"fmt"
	"os"
	"testing"
	"transactions/pkg/interfaces"
	"transactions/pkg/transactions"
)

func TestCreateFile(t *testing.T) {

	// Arrange
	var builder interfaces.Builder
	lines := 1000
	expected := 1001
	builder = &transactions.FileBuilder{FilePath: "./data/txns.csv"}

	// Act
	fmt.Println("Creating file...")
	filePath, err := builder.CreateFile(lines)
	if err != nil {
		t.Errorf("Unexpected error creating file %v", err)
	}

	// Assert
	output := readLines(filePath, t)
	assert(expected, output, t)
}

func readLines(filePath string, t *testing.T) int {

	file, err := os.Open(filePath)

	if err != nil {
		t.Errorf("Could not open file %v: %v",filePath, err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	lines, err := reader.ReadAll()

	if err != nil {
		t.Error(err)
	}

	return len(lines)
}

func assert(expected, output int, t *testing.T) {
	if expected != output {
		t.Fatalf(`Total lines in file does not match: expected: %v, output %v:`, expected, output)
	}
}
