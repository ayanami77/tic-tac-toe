package ui

import "fmt"

func NotifySituation(currentPlayer string) {
	if currentPlayer == "player1" {
		BluePrintf("\nNow %s's turn!", currentPlayer)
	} else {
		RedPrintf("\nNow %s's turn!", currentPlayer)
	}
	fmt.Print("\n\n# Press Space to flip.\n# Press Ctrl+C or ESC to quit.")
}
