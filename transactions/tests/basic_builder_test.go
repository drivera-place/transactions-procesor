package tests

import (
	"encoding/csv"
	"fmt"
	"os"
	"testing"
	"transactions/pkg/imp"
	"transactions/pkg/interfaces"
)

func TestCreateFile(t *testing.T) {

	// Arrange
	var builder interfaces.Builder
	lines := 1000
	expected := 1001
	builder = &imp.BasicBuilder{Lines: lines, Path: "txns.csv"}

	// Act
	fmt.Println("Creating file...")
	filePath, err := builder.Create(lines)
	if err != nil {
		t.Error(err)
	}

	// Assert
	output := readLines(filePath, t)
	assert(expected, output, t)
}

func readLines(filePath string, t *testing.T) int {

	if _, err := os.Stat(filePath); err != nil {
		t.Fatalf("File does not exist\n")
	}

	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		t.Fatalf("Error while reading the file.\n")
	}

	reader := csv.NewReader(file)
	reader.Comma = ','
	lines, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
	}

	return len(lines)
}

func assert(expected, output int, t *testing.T) {
	if expected != output {
		t.Fatalf(`Total lines in file does not match: expected: %v, output %v:`, expected, output)
	}
}
