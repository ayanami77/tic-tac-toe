package player

import (
	"fmt"
	"tic-tac-toe/cmd/ui"
)

const (
	PLAYER1 = "player1"
	PLAYER2 = "player2"
)

type Player struct {
	Pattern string
}

type PlayerController struct {
	Players       map[string]Player
	CurrentPlayer string
}

func (c *PlayerController) Init() {
	c.Players = map[string]Player{
		PLAYER1: {
			Pattern: "X",
		},
		PLAYER2: {
			Pattern: "O",
		},
	}

	c.CurrentPlayer = PLAYER1
}

func (c *PlayerController) Switch() {
	if c.CurrentPlayer == PLAYER1 {
		c.CurrentPlayer = PLAYER2
	} else {
		c.CurrentPlayer = PLAYER1
	}
}

func NotifyRules(currentPlayer string) {
	if currentPlayer == PLAYER1 {
		ui.BluePrintf("\nNow %s's turn!", currentPlayer)
	} else {
		ui.RedPrintf("\nNow %s's turn!", currentPlayer)
	}

	fmt.Print("\n\n# Press Space to flip.\n# Press Ctrl+C or ESC to quit.")
}

func NotifyWinner(currentPlayer string) {
	ui.GreenPrintf("\nWinner: %s\n", currentPlayer)
}

func NotifyDraw() {
	ui.GreenPrintf("\nDraw:)\n")
}
