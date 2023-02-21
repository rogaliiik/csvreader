package main

import (
	"encoding/csv"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestoreCSV_common(t *testing.T) {
	fileName := "./tables/test1.csv"
	file, _ := os.Open(fileName)

	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ','

	result, _ := restoreCSV(reader)
	expected := ",A,B,Cell\n" +
		"1,1,0,1\n" +
		"2,2,6,0\n" +
		"30,0,1,5"
	assert.Equal(t, expected, result)

}

func TestRestoreCSV_cyclicLinks(t *testing.T) {
	fileName := "./tables/test2.csv"
	file, _ := os.Open(fileName)

	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ','

	_, err := restoreCSV(reader)
	expected := "impossible to solve problem, there are cyclic links in table"
	assert.EqualError(t, err, expected)
}

func TestRestoreCSV_wrongData(t *testing.T) {
	fileName := "./tables/test3.csv"
	file, _ := os.Open(fileName)

	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ','

	_, err := restoreCSV(reader)
	expected := "row '777' not in table"
	assert.EqualError(t, err, expected)
}

func TestRestoreCSV_noColumns(t *testing.T) {
	fileName := "./tables/test4.csv"
	file, _ := os.Open(fileName)

	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ','

	_, err := restoreCSV(reader)
	expected := "there is no first column in the table"
	assert.EqualError(t, err, expected)
}

func TestRestoreCSV_wrongCellFormat(t *testing.T) {
	fileName := "./tables/test5.csv"
	file, _ := os.Open(fileName)

	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ','

	_, err := restoreCSV(reader)
	expected := "there is no operand in the cell: =A1B2 [1][1]"
	assert.EqualError(t, err, expected)
}

func TestRestoreCSV_duplicateRow(t *testing.T) {
	fileName := "./tables/test6.csv"
	file, _ := os.Open(fileName)

	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ','

	_, err := restoreCSV(reader)
	expected := "row [ A B Cell] is duplicated"
	assert.EqualError(t, err, expected)
}

func TestRestoreCSV_cellEmpty(t *testing.T) {
	fileName := "./tables/test7.csv"
	file, _ := os.Open(fileName)

	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ','

	_, err := restoreCSV(reader)
	expected := "cell value is empty string"
	assert.EqualError(t, err, expected)
}

func TestRestoreCSV_commonNegative(t *testing.T) {
	fileName := "./tables/test8.csv"
	file, _ := os.Open(fileName)

	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ','

	result, _ := restoreCSV(reader)
	expected := ",A,B,Cell\n" +
		"-9,0,-4,5\n" +
		"1,-4,0,-1\n" +
		"2,2,-9,0\n" +
		"30,1,5,-5"
	assert.Equal(t, expected, result)

}

func TestSplitAndFindCell_invalidFormat(t *testing.T) {
	_, err := splitAndFindCell("A")
	assert.EqualError(t, err, "argument has invalid format, A")
}

func TestCalculateAdd(t *testing.T) {
	result, _ := calculate("2", "1", '+')
	assert.Equal(t, "3", result)
}

func TestCalculateSubtract(t *testing.T) {
	result, _ := calculate("2", "1", '-')
	assert.Equal(t, "1", result)
}

func TestCalculateProduct(t *testing.T) {
	result, _ := calculate("2", "3", '*')
	assert.Equal(t, "6", result)
}

func TestCalculateDivision(t *testing.T) {
	result, _ := calculate("2", "2", '/')
	assert.Equal(t, "1", result)
}

func TestCalculateZeroDivision(t *testing.T) {
	_, err := calculate("2", "0", '/')
	assert.EqualError(t, err, "zero division")
}
