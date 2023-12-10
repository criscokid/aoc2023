package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/criscokid/aoc2023/internal/fileinput"
)

func main() {
	file_path := "input.txt"
	lines, err := fileinput.ReadLines(file_path)
	if err != nil {
		log.Fatal(err)
	}

	rows := parseValues(lines)
	exSum := 0
	for i := 0; i < len(rows); i++ {
		row := rows[i]
		slices.Reverse(row)
		exValue := findExtroplatedValue(row)
		exSum += exValue
	}
	fmt.Println()
	fmt.Println(exSum)
}

func parseValues(input []string) [][]int {
	rows := make([][]int, len(input))
	for i := 0; i < len(input); i++ {
		values := strings.Fields(input[i])
		row := make([]int, len(values))
		for j := 0; j < len(values); j++ {
			val, err := strconv.Atoi(values[j])
			if err != nil {
				log.Fatal(err)
			}
			row[j] = val
		}
		rows[i] = row
	}

	return rows
}

func findExtroplatedValue(row []int) int {
	currentRow := row
	fmt.Println()
	fmt.Println(currentRow)
	exValue := row[len(row)-1]
	fmt.Println(exValue)
	for {
		newRow := []int{}
		allZero := true
		for i := 0; i < len(currentRow); i++ {
			if i+1 < len(currentRow) {
				diff := currentRow[i+1] - currentRow[i]
				if diff != 0 {
					allZero = false
				}
				newRow = append(newRow, diff)
			}
		}
		currentRow = newRow
		fmt.Println(currentRow)
		if allZero || len(currentRow) == 1 {
			break
		} else {
			exValue += currentRow[len(currentRow)-1]
		}

		fmt.Println(exValue)
	}

	return exValue
}
