package main

import (
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/criscokid/aoc2023/internal/fileinput"
	"github.com/criscokid/aoc2023/internal/grids"
	"github.com/criscokid/aoc2023/internal/mathutils"
)

type galaxy struct {
	number           string
	location         grids.GridCoords
	originalLocation grids.GridCoords
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

	expansion := 1000000

	grid := grids.NewGridWithRows(lines)
	reader := grids.NewGridReader(&grid)

	galaxies := []galaxy{}

	i := 1
	for {
		if reader.CurrentValue() == '#' {
			val := strconv.Itoa(i)
			galaxies = append(galaxies, galaxy{number: val, location: reader.GetCurrentCoords(), originalLocation: reader.GetCurrentCoords()})
			i++
		}

		err := reader.Advance()
		if err == io.EOF {
			break
		}
	}

	for i := 0; i < len(grid.Data); i++ {
		idx := reader.CheckRowForValue(i, '#')
		if idx == -1 {
			for j := 0; j < len(galaxies); j++ {
				if galaxies[j].originalLocation.Y > i {
					galaxies[j].location.Y += expansion - 1
				}
			}
		}
	}

	for i := 0; i < len(grid.Data[0]); i++ {
		idx := reader.CheckColumnForValue(i, '#')
		if idx == -1 {
			for j := 0; j < len(galaxies); j++ {
				if galaxies[j].originalLocation.X > i {
					galaxies[j].location.X += expansion - 1
				}
			}
		}
	}

	fmt.Println(galaxies)

	combos := createGalaxyCombos(galaxies)
	totalLength := 0
	for _, gc := range combos {
		length := findManhattanDistance(gc)
		totalLength += length
	}
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

func findManhattanDistance(gc galaxyConnection) int {
	xDiff := mathutils.Abs(gc.start.location.X - gc.end.location.X)
	yDiff := mathutils.Abs(gc.start.location.Y - gc.end.location.Y)
	return xDiff + yDiff

}
