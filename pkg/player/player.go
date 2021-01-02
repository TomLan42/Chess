package player

import (
	"fmt"

	"github.com/TomLan42/Chess/pkg/piece"
)

// Colour is either black or white
type Colour int

const (
	//BLACK representes black side
	BLACK Colour = iota
	//WHITE representes black side
	WHITE
)

// Player is one of the two chess players in a game
type Player struct {
	Side     Colour
	Assets   map[string]*piece.Piece
	Captured map[string]*piece.Piece
}

// Select a piece according to its piece id
func (p *Player) Select(id string) (*piece.Piece, error) {
	if piece, found := p.Assets[id]; !found {
		return piece, nil
	}
	return nil, fmt.Errorf("Piece %s is not in player %d's asset", id, p.Side)
}
