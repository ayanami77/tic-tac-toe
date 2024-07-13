package main

import (
	"tic-tac-toe/internal/board"
	"tic-tac-toe/internal/player"

	"github.com/eiannone/keyboard"
)

func main() {
	pc := player.PlayerController{}
	pc.Init()

	b := board.Board{}
	b.Init()

	b.RenderPreValues()

	pc.NotifyRules()

	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		event := <-keysEvents
		if event.Err != nil {
			panic(event.Err)
		}

		if event.Key == keyboard.KeyArrowRight {
			b.Manipulate("right")
			b.RenderPreValues()
		}

		if event.Key == keyboard.KeyArrowLeft {
			b.Manipulate("left")
			b.RenderPreValues()
		}

		if event.Key == keyboard.KeyArrowUp {
			b.Manipulate("up")
			b.RenderPreValues()
		}

		if event.Key == keyboard.KeyArrowDown {
			b.Manipulate("down")
			b.RenderPreValues()
		}

		if event.Key == keyboard.KeySpace && !b.CheckIsFlipped() {
			b.Flip(b.CurrentPattern)

			b.Render(pc.CurrentPlayer)

			if b.CheckHasWinner() {
				pc.NotifyWinner()
				break
			}

			if b.CheckIfDraw() {
				pc.NotifyDraw()
				break
			}

			pc.Switch()
			b.SetPattern(pc.Players[pc.CurrentPlayer].Pattern)
		}

		if event.Key == keyboard.KeyCtrlC || event.Key == keyboard.KeyEsc {
			break
		}

		pc.NotifyRules()
	}
}
