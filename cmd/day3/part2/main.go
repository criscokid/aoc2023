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

type Gear struct {
	grids.GridCoords
	val string
}

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
	connectedVals := []Gear{}
	main_loop:
	for {
		num := []rune{}
		if slices.Index(numbers, reader.CurrentValue()) > -1 {
			num = append(num, reader.CurrentValue())
			gear := Gear{}
			add := false
			for {
				adjs := reader.GetAdjacentValues()
				for _, a := range adjs {
					if a.Char == '*' {
						gear.GridCoords = a.GridCoords
						add = true
					}
				}
				err := reader.Advance()
				if err == io.EOF {
					if add {
						gear.val = string(num)
						connectedVals = append(connectedVals, gear)
					}
					break main_loop
					
				}
				if slices.Index(numbers, reader.CurrentValue()) == -1 {
					if add {
						gear.val = string(num)
						connectedVals = append(connectedVals, gear)
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

	gearGroups := map[grids.GridCoords][]string{}
	for _, v := range connectedVals {
		if _, ok := gearGroups[v.GridCoords]; ok {
			gearGroups[v.GridCoords] = append(gearGroups[v.GridCoords], v.val)
		} else {
			gearGroups[v.GridCoords] = []string{v.val}
		}
	}

	fmt.Println(gearGroups)
	
	sum := 0
	for _, v := range gearGroups {
		if len(v) < 2 {
			continue
		}
	
		product := 1
		for _, teeth := range v {
			t, err := strconv.Atoi(teeth)
			if err != nil {
				log.Fatal(err.Error())
			}
			product *= t
		}
		sum += product
	}

	fmt.Println(sum)
}
