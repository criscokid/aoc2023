package grids

import (
	"fmt"
	"io"
	"slices"
)

type Grid struct {
	Data [][]rune
}

func NewGrid() Grid {
	data := make([][]rune, 0)
	return Grid{Data: data}
}

func NewGridWithRows(input []string) Grid {
	grid := NewGrid()
	grid.AddRows(input)
	return grid
}

func (g *Grid) AddRow(input string) {
	g.Data = append(g.Data, []rune(input))
}

func (g *Grid) AddRows(input []string) {
	for _, v := range input {
		g.AddRow(v)
	}
}

func (g *Grid) InsertRowAt(row []rune, y int) {
	if len(g.Data) == y {
		g.Data = append(g.Data, row)
		return
	}

	g.Data = append(g.Data[:y+1], g.Data[y:]...)
	g.Data[y] = row
}

func (g *Grid) InsertColumnAt(col []rune, x int){
	for y := 0; y < len(g.Data); y++ {
		if len(g.Data[y]) == x {
			g.Data[y] = append(g.Data[y], col[y])
			continue
		}

		g.Data[y] = append(g.Data[y][:x+1], g.Data[y][x:]...)
		g.Data[y][x] = col[y]
	}
}

type GridReader struct {
	*Grid
	currentCoords GridCoords
}

type GridCoords struct {
	X int
	Y int
}

type AdjResult struct {
	GridCoords
	Char rune
	Direction
}

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
	UP_LEFT
	UP_RIGHT
	DOWN_LEFT
	DOWN_RIGHT
)

var mainDirections []Direction = []Direction{UP, DOWN, LEFT, RIGHT}
var diagonalDirections []Direction = []Direction{UP_LEFT, UP_RIGHT, DOWN_LEFT, DOWN_RIGHT}

func NewGridReader(grid *Grid) GridReader {
	return GridReader{Grid: grid, currentCoords: GridCoords{X: 0, Y: 0}}
}

func (gr *GridReader) CurrentValue() rune {
	return gr.Data[gr.currentCoords.Y][gr.currentCoords.X]
}

func (gr *GridReader) SetCurrentValue(v rune) {
	gr.Data[gr.currentCoords.Y][gr.currentCoords.X] = v
}

func (gr *GridReader) Advance() error {
	return gr.advanceCoords()
}

func (gr *GridReader) AdvanceUntil(v rune) bool {
	for gr.Data[gr.currentCoords.Y][gr.currentCoords.X] != v {
		err := gr.advanceCoords()
		if err == io.EOF {
			return false
		}
	}
	return true
}

func (gr GridReader) GetAdjacentValues(includeDiagonal bool) []AdjResult {
	directions := mainDirections
	if includeDiagonal {
		directions = append(directions, diagonalDirections...)
	}
	values := []AdjResult{}
	for _, d := range directions {
		ok, coords := gr.getCoordsForDIrection(d)
		if ok {
			char := gr.Data[coords.Y][coords.X]
			values = append(values, AdjResult{Char: char, GridCoords: coords, Direction: d})
		}
	}
	return values
}

func (gr *GridReader) advanceCoords() error {
	coords := gr.currentCoords
	coords.X += 1
	if coords.X >= len(gr.Data[gr.currentCoords.Y]) {
		coords.X = 0
		coords.Y += 1
	}
	if coords.Y >= len(gr.Data) || coords.X >= len(gr.Data[coords.Y]) {
		return io.EOF
	}
	gr.currentCoords = coords
	return nil
}

func (gr GridReader) GetCurrentCoords() GridCoords {
	return gr.currentCoords
}

func (gr *GridReader) SetCurrentCoords(gc GridCoords) {
	gr.currentCoords = gc
}

func (gr *GridReader) GetValueAtCoords(gc GridCoords) rune {
	return gr.Data[gc.Y][gc.X]
}

func (gr *GridReader) SetValueAtCoords(gc GridCoords, value rune) {
	gr.Data[gc.Y][gc.X] = value
}

func (gr *GridReader) TryMoveDirection(d Direction) bool {
	ok, coords := gr.getCoordsForDIrection(d)
	if ok {
		gr.currentCoords = coords
	}
	return ok
}

func (g Grid) PrintGrid() {
	for y := 0; y < len(g.Data); y++ {
		for x := 0; x < len(g.Data[y]); x++ {
			fmt.Print(string(g.Data[y][x]))
		}
		fmt.Print("\n")
	}
}

func (gr *GridReader) CheckRowForValue(row int, val rune) int {
	return slices.Index(gr.Data[row], val)
}

func (gr *GridReader) CheckColumnForValue(col int, val rune) int {
	for i := 0; i < len(gr.Data); i++ {
		char := gr.Data[i][col]
		if char == val {
			return i
		} 
	}

	return -1
}


func (gr *GridReader) getCoordsForDIrection(d Direction) (bool, GridCoords) {
	switch d {

	case LEFT:
		if gr.currentCoords.X-1 >= 0 {
			return true, GridCoords{X: gr.currentCoords.X - 1, Y: gr.currentCoords.Y}
		}
	case RIGHT:
		if gr.currentCoords.X+1 < len(gr.Data[gr.currentCoords.Y]) {
			return true, GridCoords{X: gr.currentCoords.X + 1, Y: gr.currentCoords.Y}
		}
	case UP:
		if gr.currentCoords.Y-1 >= 0 {
			return true, GridCoords{X: gr.currentCoords.X, Y: gr.currentCoords.Y - 1}
		}
	case DOWN:
		if gr.currentCoords.Y+1 < len(gr.Data) {
			return true, GridCoords{X: gr.currentCoords.X, Y: gr.currentCoords.Y + 1}
		}
	case UP_LEFT:
		if gr.currentCoords.Y-1 >= 0 && gr.currentCoords.X-1 >= 0 {
			return true, GridCoords{X: gr.currentCoords.X - 1, Y: gr.currentCoords.Y - 1}
		}
	case UP_RIGHT:
		if gr.currentCoords.X+1 < len(gr.Data[gr.currentCoords.Y]) && gr.currentCoords.X-1 >= 0 {
			return true, GridCoords{X: gr.currentCoords.X + 1, Y: gr.currentCoords.Y - 1}
		}
	case DOWN_LEFT:
		if gr.currentCoords.X-1 >= 0 && gr.currentCoords.Y+1 < len(gr.Data) {
			return true, GridCoords{X: gr.currentCoords.X - 1, Y: gr.currentCoords.Y + 1}
		}
	case DOWN_RIGHT:
		if gr.currentCoords.X+1 < len(gr.Data[gr.currentCoords.Y]) && gr.currentCoords.Y+1 < len(gr.Data) {
			return true, GridCoords{X: gr.currentCoords.X + 1, Y: gr.currentCoords.Y + 1}
		}
	}

	return false, GridCoords{}
}
