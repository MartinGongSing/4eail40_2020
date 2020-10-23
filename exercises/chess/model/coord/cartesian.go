package coord

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
	panic("TODO")
}

func (c Cartesian) String() string {
	panic("TODO")
}