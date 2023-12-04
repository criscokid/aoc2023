package grids

import (
	"io"
)

type Grid struct {
	data [][]rune
}

func NewGrid() Grid {
	data := make([][]rune, 0)
	return Grid{data: data}
}

func (g *Grid) AddRow(input string) {
	g.data = append(g.data, []rune(input))
}

type GridReader struct {
	Grid
	currentCoords GridCoords
}

type GridCoords struct {
	x int
	y int
}

type AdjResult struct {
	GridCoords
	Char rune
}

func NewGridReader(grid Grid) GridReader {
	return GridReader{Grid: grid, currentCoords: GridCoords{x: 0, y: 0}}
}

func (gr *GridReader) CurrentValue() rune {
	return gr.data[gr.currentCoords.y][gr.currentCoords.x]
}

func (gr *GridReader) Advance() error {
	return gr.advanceCoords()
}

func (gr GridReader) GetAdjacentValues() []AdjResult {
	coords := gr.findAdjuscentCoords()
	values := []AdjResult{}
	for _, v := range coords {
		char := gr.data[v.y][v.x]
		values = append(values, AdjResult{Char: char, GridCoords: v})
	}
	return values
}

func (gr *GridReader) advanceCoords() error {
	coords := gr.currentCoords
	coords.x += 1
	if coords.x >= len(gr.data[gr.currentCoords.y]) {
		coords.x = 0
		coords.y += 1
	}
	if coords.y >= len(gr.data) || coords.x >= len(gr.data[coords.y]) {
		return io.EOF
	}
	gr.currentCoords = coords
	return nil
}

func (gr GridReader) findAdjuscentCoords() []GridCoords {
	coords := []GridCoords{}

	if gr.currentCoords.x - 1 >= 0 {
		coords = append(coords, GridCoords{x: gr.currentCoords.x - 1, y: gr.currentCoords.y})
	}

	if gr.currentCoords.x + 1 < len(gr.data[gr.currentCoords.y]) {
		coords = append(coords, GridCoords{x: gr.currentCoords.x + 1, y: gr.currentCoords.y})
	}

	if gr.currentCoords.y - 1 >= 0 {
		coords = append(coords, GridCoords{x: gr.currentCoords.x, y: gr.currentCoords.y - 1})
	}

	if gr.currentCoords.y + 1 < len(gr.data) {
		coords = append(coords, GridCoords{x: gr.currentCoords.x, y: gr.currentCoords.y + 1})
	}

	if gr.currentCoords.y - 1 >= 0 {
		//up and left
		if gr.currentCoords.x - 1 >= 0 {
			coords = append(coords, GridCoords{x: gr.currentCoords.x - 1, y: gr.currentCoords.y - 1})
		}
		
		//up and right
		if gr.currentCoords.x + 1 < len(gr.data[gr.currentCoords.y]) {
			coords = append(coords, GridCoords{x: gr.currentCoords.x + 1, y: gr.currentCoords.y - 1})
		}
	}

	if gr.currentCoords.y + 1 < len(gr.data) {
		//down and left
		if gr.currentCoords.x - 1 >= 0 {
			coords = append(coords, GridCoords{x: gr.currentCoords.x - 1, y: gr.currentCoords.y + 1})
		}
		
		//down and right
		if gr.currentCoords.x + 1 < len(gr.data[gr.currentCoords.y]) {
			coords = append(coords, GridCoords{x: gr.currentCoords.x + 1, y: gr.currentCoords.y + 1})
		}
	}
	return coords
}

func (gr GridReader) GetCurrentCoords() GridCoords {
	return gr.currentCoords
}

func (gr *GridReader) SetCurrentCoords(gc GridCoords) {
	gr.currentCoords = gc
}
