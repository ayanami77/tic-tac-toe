package player

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
