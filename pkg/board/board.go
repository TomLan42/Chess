package board

import "github.com/TomLan42/Chess/pkg/piece"

// Board is a standard 8x8 chess board
type Board struct {
	Grids [8][8]*Grid
}

// Grid is a single cell on the chess board
type Grid struct {
	GirdID string
	// Owner is the piece that occupies the current gird
	Owner *piece.Piece
	// Taints are the pieces whose current attack ranges covers
	// this grid
	Taints []*piece.Piece
}
