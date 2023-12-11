package main

import (
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/criscokid/aoc2023/internal/fileinput"
	"github.com/criscokid/aoc2023/internal/grids"
)

type galaxy struct {
	number   string
	location grids.GridCoords
}

type galaxyConnection struct {
	start    galaxy
	end      galaxy
	distance int
}

func main() {
	file_path := "input.txt"
	lines, err := fileinput.ReadLines(file_path)
	if err != nil {
		log.Fatal(err)
	}

	originalRowLength := len(lines[0])
	emptyRow := make([]rune, originalRowLength)
	for i := 0; i < len(emptyRow); i++ {
		emptyRow[i] = '.'
	}

	grid := grids.NewGridWithRows(lines)
	reader := grids.NewGridReader(&grid)

	for i := 0; i < len(grid.Data); i++ {
		idx := reader.CheckRowForValue(i, '#')
		if idx == -1 {
			grid.InsertRowAt(emptyRow, i+1)
			i++
		}
	}

	emptyCol := make([]rune, len(grid.Data))
	for i := 0; i < len(grid.Data); i++ {
		emptyCol[i] = '.'
	}

	for i := 0; i < len(grid.Data[0]); i++ {
		idx := reader.CheckColumnForValue(i, '#')
		if idx == -1 {
			grid.InsertColumnAt(emptyCol, i)
			i++
		}
	}

	galaxies := []galaxy{}

	i := 1
	for {
		if reader.CurrentValue() == '#' {
			val := strconv.Itoa(i)
			galaxies = append(galaxies, galaxy{number: val, location: reader.GetCurrentCoords()})
			i++
		}

		err := reader.Advance()
		if err == io.EOF {
			break
		}
	}

	combos := createGalaxyCombos(galaxies)
	totalLength := 0
	for _, gc := range combos {
		length := findPathLength(gc)
		totalLength += length
	}
	grid.PrintGrid()
	fmt.Println(totalLength)	
}

func createGalaxyCombos(galaxies []galaxy) []galaxyConnection {
	combos := []galaxyConnection{}
	for i := 0; i < len(galaxies); i++ {
		for j := i; j < len(galaxies); j++ {
			if i == j {
				continue
			}
			combos = append(combos, galaxyConnection{start: galaxies[i], end: galaxies[j]})
		}
	}

	return combos
}

func findPathLength(gc galaxyConnection) int {
	// fmt.Printf("%s -> %s", string(gc.start.number), string(gc.end.number))
	position := gc.start.location
	end := gc.end.location
	length := 0
	for {
		if position.Y < end.Y {
			length++
			position.Y++
		}

		if position.X < end.X {
			length++
			position.X++
		} else if position.X > end.X {
			length++
			position.X--
		}

		if position.X == end.X && position.Y == end.Y {
			break
		}

	}
	
	return length
}
