package main

import (
	"tic-tac-toe/cmd/board"
	"tic-tac-toe/cmd/player"

	"github.com/eiannone/keyboard"
)

func main() {
	pc := player.PlayerController{}
	pc.Init()

	b := board.Board{}
	b.Init()

	b.RenderPreValues()

	player.NotifyRules(pc.CurrentPlayer)

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
				player.NotifyWinner(pc.CurrentPlayer)
				break
			}

			if b.CheckIfDraw() {
				player.NotifyDraw()
				break
			}

			pc.Switch()
			b.SetPattern(pc.Players[pc.CurrentPlayer].Pattern)
		}

		player.NotifyRules(pc.CurrentPlayer)

		if event.Key == keyboard.KeyCtrlC || event.Key == keyboard.KeyEsc {
			break
		}
	}
}
