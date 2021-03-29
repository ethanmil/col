package game

type Map [][]Tile

func NewMap(w, h int) Map {
	m := [][]Tile{{}}
	for y := 0; y < h; y++ {
		m[y] = []Tile{}
		for x := 0; x < w; x++ {
			m[y][x] = NewTile(x, y, "forest")
		}
	}

	return m
}
