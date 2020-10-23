// Package board will content types and logic for handling chess board(s)
package board

import (
	"fmt"

	"github.com/MartinGongSing/4eail40_2020/exercices/chess/model/coord"
	"github.com/MartinGongSing/4eail40_2020/exercices/chess/model/piece"
)

// Board is an interface to a chess board.
// It defines methods for handling pieces on it
type Board interface {
	fmt.Stringer
	// PieceAT retrives pieces at given coordinates
	// Returns nill if no piece was found
	PieceAt(at coord.ChessCoordinates) piece.Piece
	// MovePiece moves a piece from given coordinates
	// to given coordinates
	// Returns an error if destination was occupied
	MovePiece(from, to coord.ChessCoordinates) error
	// PlacePieceAt ^places a given piece at given location
	// Returns an error if destination was occupied
	PlacePieceAt(p piece.Piece, at coord.ChessCoordinates) error
}
