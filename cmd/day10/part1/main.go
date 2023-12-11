package main

import (
	"fmt"
	"log"

	"github.com/criscokid/aoc2023/internal/fileinput"
	"github.com/criscokid/aoc2023/internal/grids"
)

var allowDirections map[rune][]grids.Direction = map[rune][]grids.Direction{
	'|': {grids.UP, grids.DOWN},
	'-': {grids.LEFT, grids.RIGHT},
	'L': {grids.UP, grids.RIGHT},
	'J': {grids.UP, grids.LEFT},
	'7': {grids.DOWN, grids.LEFT},
	'F': {grids.DOWN, grids.RIGHT},
}

func main() {
	file_path := "input.txt"
	lines, err := fileinput.ReadLines(file_path)
	if err != nil {
		log.Fatal(err)
	}

	grid := grids.NewGridWithRows(lines)
	reader := grids.NewGridReader(grid)
	navgiateMaze('J', &reader)
}

func navgiateMaze(replaceStartWith rune, reader *grids.GridReader) bool {
	reader.SetCurrentCoords(grids.GridCoords{X: 0, Y: 0})
	reader.AdvanceUntil('S')
	start := reader.GetCurrentCoords()
	fmt.Println(start)
	seen := []grids.GridCoords{}

	adjs := getValidAdjcents(reader, replaceStartWith)
	ok := reader.TryMoveDirection(adjs[0].Direction)
	seen = append(seen, start, reader.GetCurrentCoords())
	for ok {
		adjs = getValidAdjcents(reader, 0)
		var dirToMove grids.AdjResult
		previous := seen[len(seen)-2]
		for _, a := range adjs {
			if a.X == previous.X && a.Y == previous.Y {
				continue
			}
			dirToMove = a
		}

		if dirToMove.X == start.X && dirToMove.Y == start.Y {
			break
		}

		ok = reader.TryMoveDirection(dirToMove.Direction)
		seen = append(seen, reader.GetCurrentCoords())
	}

	fmt.Println(seen)
	fmt.Println(len(seen)/2-1)
	return false
}

func getValidAdjcents(reader *grids.GridReader, overrideCurrentValue rune) []grids.AdjResult {
	adjs := reader.GetAdjacentValues(false)
	validDirs := allowDirections[reader.CurrentValue()]
	if overrideCurrentValue > 0 {
		validDirs = allowDirections[overrideCurrentValue]
	}
	valid := []grids.AdjResult{}

	for _, a := range adjs {
		if validDirs[0] == a.Direction || validDirs[1] == a.Direction {
			valid = append(valid, a)
		}
	}

	return valid
}
