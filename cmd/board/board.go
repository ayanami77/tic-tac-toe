package board

import (
	"fmt"
	"slices"
	"tic-tac-toe/cmd/ui"
)

type Board struct {
	CurrentX       int
	CurrentY       int
	CurrentPattern string
	Values         [][]string
	PreValues      [][]string
}

func (b *Board) Init() {
	b.CurrentX = 0
	b.CurrentY = 0
	b.CurrentPattern = "X"
	b.Values = [][]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}
	b.PreValues = b.makePreValues(b.Values)
	b.PreValues[0][0] = "@"
}

func (b *Board) Manipulate(direction string) {
	b.controlPosition(direction)

	b.setPreValues()
}

func (b *Board) controlPosition(direction string) {
	const (
		leftEdge  = 0
		rightEdge = 2
		upEdge    = 0
		downEdge  = 2
	)

	if direction == "left" && leftEdge <= b.CurrentX-1 {
		b.CurrentX--
	}
	if direction == "right" && b.CurrentX+1 <= rightEdge {
		b.CurrentX++
	}
	if direction == "up" && upEdge <= b.CurrentY-1 {
		b.CurrentY--
	}
	if direction == "down" && b.CurrentY+1 <= downEdge {
		b.CurrentY++
	}
}

func (b *Board) SetPattern(pattern string) {
	b.CurrentPattern = pattern
}

func (b *Board) setPreValues() {
	preValues := b.makePreValues(b.Values)

	if preValues[b.CurrentY][b.CurrentX] == "-" {
		preValues[b.CurrentY][b.CurrentX] = "@"
	}

	b.PreValues = preValues
}

func (b *Board) makePreValues(values [][]string) [][]string {
	preValues := make([][]string, len(values))

	for i := range values {
		preValues[i] = make([]string, len(values[i]))
		copy(preValues[i], values[i])
	}

	return preValues
}

func (b *Board) RenderPreValues() {
	ui.Clear()

	b.renderHelper(b.PreValues)
}

func (b *Board) renderHelper(values [][]string) {
	for i := 0; i < len(values); i++ {
		for _, v := range values[i] {
			if v == "X" {
				ui.BluePrintf("%s ", v)
			} else if v == "O" {
				ui.RedPrintf("%s ", v)
			} else if v == "@" {
				if b.CurrentPattern == "X" {
					ui.BluePrintf("- ")
				} else {
					ui.RedPrintf("- ")
				}
			} else {
				fmt.Printf("%s ", v)
			}
		}

		fmt.Print("\n")
	}
}

func (b *Board) CheckIsFlipped() bool {
	return b.Values[b.CurrentY][b.CurrentX] != "-"
}

func (b *Board) Flip(pattern string) {
	b.Values[b.CurrentY][b.CurrentX] = pattern
}

func (b *Board) Render(currentPlayer string) {
	ui.Clear()

	b.renderHelper(b.Values)
}

func (b *Board) CheckHasWinner() bool {
	patterns := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}

	values := slices.Concat(b.Values...)

	for _, pattern := range patterns {
		if values[pattern[0]] != "-" &&
			values[pattern[1]] != "-" &&
			values[pattern[2]] != "-" &&
			values[pattern[0]] == values[pattern[1]] &&
			values[pattern[1]] == values[pattern[2]] {
			return true
		}
	}

	return false
}
