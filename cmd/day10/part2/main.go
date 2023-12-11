package main

import (
	"fmt"
	"io"
	"log"
	"slices"

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

var wallRunes []rune = []rune{'|', '-', 'L', 'J', '7', 'F', 'S'}

func main() {
	file_path := "input.txt"
	lines, err := fileinput.ReadLines(file_path)
	if err != nil {
		log.Fatal(err)
	}
	
	//Create exapnded gird
	grid := grids.NewGrid()
	verticalRow := []rune{}
	for i := 0; i < len(lines[0])*2+1; i++ {
		verticalRow = append(verticalRow, '|')
	}
	grid.AddRow(string(verticalRow))
	for _, line := range lines {
		newLine := []rune{}
		for _, r := range line {
			newLine = append(newLine, '-', r)
		}
		newLine = append(newLine, '-')
		grid.AddRow(string(newLine))
		grid.AddRow(string(verticalRow))
	}

	reader := grids.NewGridReader(grid)
	seen := navgiateMaze('J', &reader)

	//Fill in empty spaces with 0
	reader.SetCurrentCoords(grids.GridCoords{X: 0, Y: 0})
	for {
		currentCoords := reader.GetCurrentCoords()
		idx := slices.IndexFunc(seen, func(gc grids.GridCoords) bool {
			return currentCoords.X == gc.X && currentCoords.Y == gc.Y
		})

		if idx == -1 {
			reader.SetCurrentValue('0')
		}
		err := reader.Advance()
		if err == io.EOF {
			break
		}
	}

	//fill with ones
	floodfill(&grid, 0, 0)


	//count remaining 0's with skipping the padded rows
	total := 0
	for y := 0; y < len(grid.Data); y++ {
		if y % 2 == 0 {
			continue
		}

		for x := 0; x < len(grid.Data[y]); x++ {
			if x % 2 == 0 {
				continue
			}

			if grid.Data[y][x] == '0' {
				total++
			}
		}
	}

	fmt.Println(total)

}

func navgiateMaze(replaceStartWith rune, reader *grids.GridReader) []grids.GridCoords {
	reader.SetCurrentCoords(grids.GridCoords{X: 0, Y: 0})
	reader.AdvanceUntil('S')
	start := reader.GetCurrentCoords()
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
	return seen
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

func floodfill(grid *grids.Grid, x, y int) {
	val := grid.Data[y][x]
	idx := slices.Index(wallRunes, grid.Data[y][x])
	if idx > -1 || val == '1' {
		return
	}

	grid.Data[y][x] = '1'
	if x+1 < len(grid.Data[y]) {
		floodfill(grid, x+1, y)
	}

	if x-1 >= 0 {
		floodfill(grid, x-1, y)
	}

	if y + 1 < len(grid.Data) {
		floodfill(grid, x, y+1)
	}

	if y - 1 >= 0 {
		floodfill(grid, x, y-1)
	}

	return
}
