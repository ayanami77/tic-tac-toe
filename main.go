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
			b.ControlPosition("right")
			b.MakePreValues()
			b.RenderPreValues()
		}

		if event.Key == keyboard.KeyArrowLeft {
			b.ControlPosition("left")
			b.MakePreValues()
			b.RenderPreValues()
		}

		if event.Key == keyboard.KeyArrowUp {
			b.ControlPosition("up")
			b.MakePreValues()
			b.RenderPreValues()
		}

		if event.Key == keyboard.KeyArrowDown {
			b.ControlPosition("down")
			b.MakePreValues()
			b.RenderPreValues()
		}

		if event.Key == keyboard.KeySpace && !b.CheckIsFlipped() {
			b.Flip()
			b.Render()

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

		if event.Key == keyboard.KeyArrowRight ||
			event.Key == keyboard.KeyArrowLeft ||
			event.Key == keyboard.KeyArrowUp ||
			event.Key == keyboard.KeyArrowDown ||
			event.Key == keyboard.KeySpace {
			pc.NotifyRules()
		}
	}
}
