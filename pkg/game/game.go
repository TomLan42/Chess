package game

import (
	"time"
)

// Result is the final outcome of the game, it is
// either black wins, white wins or DRAW
type Result int

const (
	//BLACKWIN means the black side wins
	BLACKWIN Result = iota
	// WHITEWIN means the white side wins
	WHITEWIN
	// DRAW means stalemate happens or both
	// sides agree on a draw
	DRAW
)

// Game is a chess game
type Game struct {
	TimeLimit *time.Duration
	Trace     []string
	Result    Result
}

// Init the game with all pieces in place, timer ready
func (g *Game) Init() {

}
