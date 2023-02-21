package main

import (
	"encoding/csv"
	"log"
	"sort"
	"strconv"
	"strings"
)

var (
	table   [][]string
	rows    map[string]int
	columns map[string]int
	awaited map[[2]int]int
)

// TODO: add error handling
// TODO: make better algo
// TODO: makefile
// TODO: Readme.md
func restoreCSV(reader *csv.Reader) string {
	table = [][]string{}
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		table = append(table, record)
	}
	sort.Slice(table, func(i int, j int) bool {
		return table[i][0] < table[j][0]
	})

	// columns map stores titles of columns
	columns = map[string]int{}
	for i, col := range table[0] {
		columns[col] = i
	}
	// columns map stores titles of rows
	rows = map[string]int{}
	for i, row := range table {
		rows[row[0]] = i
	}

	awaited = map[[2]int]int{}
	for r := 1; r < len(table); r++ {
		for c := 1; c < len(table[r]); c++ {
			if table[r][c][0] == '=' {
				evaluateCell(r, c, false)
			}
		}
	}
	last := len(awaited)
	for len(awaited) != 0 {
		for k := range awaited {
			evaluateCell(k[0], k[1], true)
		}
		if len(awaited) == last {
			log.Println("Impossible to solve problem, there are cyclic links")
			break
		}
		last = len(awaited)
	}
	var res []string
	for _, r := range table {
		res = append(res, strings.Join(r, ","))
	}
	return strings.Join(res, "\n")
}

func evaluateCell(r, c int, flag bool) {
	operations := map[rune]int{'+': 1, '-': 1, '/': 1, '*': 1}
	for i, char := range table[r][c] {
		if _, ok := operations[char]; ok {
			leftArg := splitAndFindCell(table[r][c][1:i])
			rightArg := splitAndFindCell(table[r][c][i+1:])
			if leftArg[0] == '=' || rightArg[0] == '=' {
				if !flag {
					awaited[[2]int{r, c}]++
				}
				continue
			}
			if flag {
				delete(awaited, [2]int{r, c})
			}
			table[r][c] = operate(leftArg, rightArg, char)
		}
	}
}

func splitAndFindCell(s string) string {
	for i, v := range s {
		_, err := strconv.Atoi(string(v))
		if err != nil {
			continue
		}
		return table[rows[s[i:]]][columns[s[:i]]]
	}
	return ""
}

func operate(leftArg, rightArg string, operand rune) string {
	arg1, _ := strconv.Atoi(leftArg)
	arg2, _ := strconv.Atoi(rightArg)
	switch operand {
	case '+':
		return strconv.Itoa(arg1 + arg2)
	case '*':
		return strconv.Itoa(arg1 * arg2)
	case '/':
		return strconv.Itoa(arg1 / arg2)
	default:
		return strconv.Itoa(arg1 - arg2)
	}
}
