package coord

import (
	"fmt"
)

// Cartesian represent a ser of cartesian coordinates, x and y
type Cartesian struct{
	x, y int
}

// NewCartesian is the constructer for cartesian
func NewCartesian(x, y int) Cartesian {
	return Cartesian{x, y}
}

// Coord returns x if n==0, y if n==1
func (c Cartesian) Coord(n int) (int, error){

	switch n{
	case 0:
		return c.y, nil
	case 1:
		return c.x, nil
	}
	return 0, fmt.Errorf("unknow coord component %d", n)
}

func (c Cartesian) String() string {
	return fmt.Sprintf("%c1", 65 + c.y, c.x + 1)
}