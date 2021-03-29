package game

type Tile struct {
	Terrain string
	X, Y    int
}

func NewTile(x, y int, terrain string) Tile {
	return Tile{
		X:       x,
		Y:       y,
		Terrain: terrain,
	}
}
