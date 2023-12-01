package main

import (
	"cmp"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/criscokid/aoc2023/internal/fileinput"
)

var valuesMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

var numberRunes = []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

type numberLocation struct {
	number string
	idx    int
}

var possibleValues = make([]string, 0)

func main() {

	for k, v := range valuesMap {
		possibleValues = append(possibleValues, k, v)
	}

	file_path := "input.txt"
	lines, err := fileinput.ReadLines(file_path)
	if err != nil {
		log.Fatal(err)
	}

	values := make([]int, 0)

	for _, line := range lines {
		fmt.Println(line)
		numbers := findNumbers(line)
		numVal, err := getNumberFromLineRunes(numbers)
		if err != nil {
			log.Fatal(err.Error())
		}
		values = append(values, numVal)
	}
	fmt.Println(sum(values))
}

func findNumbers(input string) []rune {
	foundNumbers := make([]numberLocation, 0)
	for _, v := range possibleValues {
		fmt.Println("value is: " + v)
		currentIdx := 0
		for currentIdx > -1 {
			idx := indexStartingAt(input, v, currentIdx)
			fmt.Println(idx)
			if idx > -1 {
				foundNum := numberLocation{number: v, idx: idx}
				foundNumbers = append(foundNumbers, foundNum)
				idx = idx + len(v)
			}
			currentIdx = idx
		}
	}
	fmt.Println(foundNumbers)
	firstDigit := slices.MinFunc(foundNumbers, func(a, b numberLocation) int {
		return cmp.Compare(a.idx, b.idx)
	})
	lastDigit := slices.MaxFunc(foundNumbers, func(a, b numberLocation) int {
		return cmp.Compare(a.idx, b.idx)
	})

	var firstRune, lastRune rune

	if len(firstDigit.number) > 1 {
		firstRune = rune(valuesMap[firstDigit.number][0])
	} else {
		firstRune = rune(firstDigit.number[0])
	}

	if len(lastDigit.number) > 1 {
		lastRune = rune(valuesMap[lastDigit.number][0])
	} else {
		lastRune = rune(lastDigit.number[0])
	}

	return []rune{firstRune, lastRune}
}

func indexStartingAt(input string, subStr string, startIdx int) int {
	idx := strings.Index(input[startIdx:], subStr)
	if idx == -1 {
		return idx
	}

	return idx + startIdx
}

func getNumberFromLineRunes(numbers []rune) (int, error) {
	var sb strings.Builder
	sb.WriteRune(numbers[0])
	sb.WriteRune(numbers[len(numbers)-1])
	return strconv.Atoi(sb.String())
}

func sum(values []int) int {
	total := 0
	for _, n := range values {
		total = total + n
	}
	return total
}
