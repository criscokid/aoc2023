package main

import (
	"fmt"
	"io"
	"log"
	"slices"
	"strconv"

	"github.com/criscokid/aoc2023/internal/fileinput"
	"github.com/criscokid/aoc2023/internal/grids"
)

var numbers = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
var nonSpecialCharacters = append(numbers, '.')

func main() {
	file_path := "input.txt"
	lines, err := fileinput.ReadLines(file_path)
	if err != nil {
		log.Fatal(err)
	}

	grid := grids.NewGrid()

	for _, line := range lines {
		grid.AddRow(line)
	}

	reader := grids.NewGridReader(grid)
	connectedVals := []string{}
	main_loop:
	for {
		num := []rune{}
		if slices.Index(numbers, reader.CurrentValue()) > -1 {
			num = append(num, reader.CurrentValue())
			add := false
			for {
				adjs := reader.GetAdjacentValues()
				for _, a := range adjs {
					if slices.Index(nonSpecialCharacters, a.Char) == -1 {
						add = true
					}
				}
				err := reader.Advance()
				if err == io.EOF {
					if add {
						connectedVals = append(connectedVals, string(num))
					}
					break main_loop
					
				}
				if slices.Index(numbers, reader.CurrentValue()) == -1 {
					if add {
						connectedVals = append(connectedVals, string(num))
					}
					break
				}
				num = append(num, reader.CurrentValue())
			}
		}
		err := reader.Advance()
		if err == io.EOF {
			break
		}
	}

	fmt.Println(connectedVals)
	sum := 0
	for _, v := range connectedVals {
		val, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err.Error())
		}
		sum += val
	}

	fmt.Println(sum)
}
