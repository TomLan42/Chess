package piece

// Piece is an object on the chess board
type Piece interface {
	GetPieceID() string
	GetPosition() [2]int
	// Move takes a coordinate on the chess board
	// and validates if it is a valid move
	Move([2]int) bool
	// Sacrifice remove the piece from the chess board
	Sacrifice()
}

// BasePiece is a generic piece implementor with minimal functionality
type BasePiece struct {
	PieceID  string
	Position [2]int
}

// GetPieceID returns the piece id of the piece
func (b *BasePiece) GetPieceID() string {
	return b.PieceID
}

// GetPosition returns the current position of the piece
func (b *BasePiece) GetPosition() [2]int {
	return b.Position
}
