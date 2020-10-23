// Package coord will contain types and logic for handling chess board coordinates
package coord

import (
	"fmt"
)
// ChessCoordinates is an interface giving the position of a chess piece
type ChessCoordinates interface {
	fmt.Stringer // "A7"
	//Coord gets n'th coordinate comp. 
	Coord(n int) error
	
}
