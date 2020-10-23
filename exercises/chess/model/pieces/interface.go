// Package pieces will handle game pieces
package pieces

import (
	"fmt"

	"github.com/MartinGongSing/4eail40_2020/exercices/chess/model/coord"
	"github.com/MartinGongSing/4eail40_2020/exercices/chess/model/player"
)

// Piece is an interface for pieces on CB
type Piece interface {
	fmt.Stringer
	Color() player.Color
	Moves() []coord.ChessCoordinates
}

func test(){
	var p Piece
	// Check if move is contained in the map
	if _, move := range p.Moves{
		// Check if it is contained, if yes brek out of loop
	}
}