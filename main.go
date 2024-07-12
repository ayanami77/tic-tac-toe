package main

import (
	board "tic-tac-toe/cmd/board"
	player "tic-tac-toe/cmd/player"
	"tic-tac-toe/cmd/ui"

	"github.com/eiannone/keyboard"
)

func main() {
	pc := player.PlayerController{}
	pc.Init()

	b := board.Board{}
	b.Init()
	b.RenderPreValues(pc.CurrentPlayer)

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
			b.RenderPreValues(pc.CurrentPlayer)
		}

		if event.Key == keyboard.KeyArrowLeft {
			b.Manipulate("left")
			b.RenderPreValues(pc.CurrentPlayer)
		}

		if event.Key == keyboard.KeyArrowUp {
			b.Manipulate("up")
			b.RenderPreValues(pc.CurrentPlayer)
		}

		if event.Key == keyboard.KeyArrowDown {
			b.Manipulate("down")
			b.RenderPreValues(pc.CurrentPlayer)
		}

		if event.Key == keyboard.KeySpace && !b.CheckIsFlipped() {
			b.Flip(b.CurrentPattern)

			b.Render(pc.CurrentPlayer)

			if b.CheckHasWinner() {
				ui.GreenPrintf("\n\n\nWinner: %s", pc.CurrentPlayer)
				break
			}

			pc.Switch()
			b.SetPattern(pc.Players[pc.CurrentPlayer].Pattern)
		}

		if event.Key == keyboard.KeyCtrlC || event.Key == keyboard.KeyEsc {
			break
		}
	}
}
