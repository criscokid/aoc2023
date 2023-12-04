package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/criscokid/aoc2023/internal/fileinput"
	"github.com/criscokid/aoc2023/internal/mathutils"
)

var numberRunes = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func main() {
	file_path := "input.txt"
	lines, err := fileinput.ReadLines(file_path)
	if err != nil {
		log.Fatal(err)
	}

	values := make([]int, 0)

	for _, line := range lines {
		numbers := findNumbers(line)
		numVal, err := getNumberFromLineRunes(numbers)
		if err != nil {
			log.Fatal(err.Error())
		}
		values = append(values, numVal)
	}

	sum := mathutils.SumSlice(values, func(v int) int { return v })
	fmt.Println(sum)

}

func findNumbers(input string) []rune {
	numbers := make([]rune, 0)
	for _, c := range input {
		if slices.Contains(numberRunes, c) {
			numbers = append(numbers, c)
		}
	}
	return numbers
}

func getNumberFromLineRunes(numbers []rune) (int, error) {
	var sb strings.Builder
	sb.WriteRune(numbers[0])
	sb.WriteRune(numbers[len(numbers)-1])
	return strconv.Atoi(sb.String())
}
