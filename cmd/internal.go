package main

import (
	"encoding/csv"
	"fmt"
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

func restoreCSV(reader *csv.Reader) (string, error) {
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
	if table[0][0] != "" {
		return "", fmt.Errorf("there is no first column in the table")
	}

	// columns map stores titles of columns
	columns = map[string]int{}
	for i, col := range table[0] {
		if _, ok := columns[col]; ok {
			return "", fmt.Errorf("column %s is duplicated", col)
		}
		columns[col] = i
	}
	// rows map stores titles of rows
	rows = map[string]int{}
	for i, row := range table {
		if _, ok := rows[row[0]]; ok {
			return "", fmt.Errorf("row %s is duplicated", row)
		}
		rows[row[0]] = i
	}

	awaited = map[[2]int]int{}
	for r := 1; r < len(table); r++ {
		for c := 1; c < len(table[r]); c++ {
			if table[r][c] == "" {
				return "", fmt.Errorf("cell value is empty string")
			}
			if table[r][c][0] == '=' {
				err := evaluateCell(r, c, false)
				if err != nil {
					return "", err
				}
			}
		}
	}
	last := len(awaited)
	for len(awaited) != 0 {
		for k := range awaited {
			err := evaluateCell(k[0], k[1], true)
			if err != nil {
				return "", err
			}
		}
		if len(awaited) == last {
			return "", fmt.Errorf("impossible to solve problem, there are cyclic links in table")
		}
		last = len(awaited)
	}
	var res []string
	for _, r := range table {
		res = append(res, strings.Join(r, ","))
	}
	return strings.Join(res, "\n"), nil
}

// evaluateCell handles a cell with a formula
// the flag indicates if we should remove the row-column pair from the map
func evaluateCell(r, c int, flag bool) error {
	operations := map[rune]int{'+': 1, '-': 1, '/': 1, '*': 1}
	for i, char := range table[r][c] {
		if _, ok := operations[char]; ok {
			leftArg, err := splitAndFindCell(table[r][c][1:i])
			if err != nil {
				return err
			}
			rightArg, err := splitAndFindCell(table[r][c][i+1:])
			if err != nil {
				return err
			}
			if leftArg[0] == '=' || rightArg[0] == '=' {
				if !flag {
					awaited[[2]int{r, c}]++
				}
				return nil
			}
			if flag {
				delete(awaited, [2]int{r, c})
			}
			res, err := calculate(leftArg, rightArg, char)
			if err != nil {
				return err
			}
			table[r][c] = res
			return nil
		}
	}
	return fmt.Errorf("there is no operand in the cell: %s [%d][%d]", table[r][c], r, c)
}

// splitAndFindCell splits string into row and column and
// returns value of the cell
func splitAndFindCell(s string) (string, error) {
	for i, v := range s {
		_, err := strconv.Atoi(string(v))
		if err != nil {
			continue
		}
		if _, ok := rows[s[i:]]; !ok {
			return "", fmt.Errorf("row '%s' not in table", s[i:])
		}
		if _, ok := columns[s[:i]]; !ok {
			return "", fmt.Errorf("column '%s' not in table", s[:i])
		}
		return table[rows[s[i:]]][columns[s[:i]]], nil
	}
	return "", fmt.Errorf("argument has invalid format, %s", s)
}

// calculate performs operations based on the received operand
func calculate(leftArg, rightArg string, operand rune) (string, error) {
	arg1, err := strconv.Atoi(leftArg)
	if err != nil {
		return "", err
	}
	arg2, err := strconv.Atoi(rightArg)
	if err != nil {
		return "", err
	}
	switch operand {
	case '+':
		return strconv.Itoa(arg1 + arg2), nil
	case '*':
		return strconv.Itoa(arg1 * arg2), nil
	case '/':
		if arg2 == 0 {
			return "", fmt.Errorf("zero division")
		}
		return strconv.Itoa(arg1 / arg2), nil
	default:
		return strconv.Itoa(arg1 - arg2), nil
	}
}
