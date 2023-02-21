package main

import (
	"encoding/csv"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRestoreCSV_common(t *testing.T) {
	fileName := "./tables/test1.csv"
	file, _ := os.Open(fileName)

	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ','

	result := restoreCSV(reader)
	expected := ",A,B,Cell\n" +
		"1,1,0,1\n" +
		"2,2,6,0\n" +
		"30,0,1,5"
	assert.Equal(t, expected, result)

}

func TestRestoreCSV_edge1(t *testing.T) {
	fileName := "./tables/test2.csv"
	file, _ := os.Open(fileName)

	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ','

	result := restoreCSV(reader)
	expected := ",A,B,Cell\n" +
		"1,1,0,1\n" +
		"2,2,=B30+Cell30,0\n" +
		"30,0,=B2+A1,5"
	assert.Equal(t, expected, result)
}

func TestRestoreCSV_edge2(t *testing.T) {
	fileName := "./tables/test3.csv"
	file, _ := os.Open(fileName)

	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ','

	result := restoreCSV(reader)
	expected := ",A,B,Cell\n" +
		"1,1,0,1\n" +
		"2,2,5,0\n" +
		"30,0,5,5"
	assert.Equal(t, expected, result)
}

func TestOperateAdd(t *testing.T) {
	result := operate("2", "1", '+')
	assert.Equal(t, "3", result)
}

func TestOperateSubtract(t *testing.T) {
	result := operate("2", "1", '-')
	assert.Equal(t, "1", result)
}

func TestOperateProduct(t *testing.T) {
	result := operate("2", "3", '*')
	assert.Equal(t, "6", result)
}

func TestOperateDivision(t *testing.T) {
	result := operate("2", "2", '/')
	assert.Equal(t, "1", result)
}
