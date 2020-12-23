package game

import (
	"github.com/hajimehoshi/ebiten"
)

// Orch -
type Orch struct {
	Art *ebiten.Image
}

// New -
func New(art *ebiten.Image) *Orch {
	return &Orch{
		Art: art,
	}
}

// Update -
func (Orch) Update(screen *ebiten.Image) error {
	return nil
}

// Draw -
func (Orch) Draw(screen *ebiten.Image) error {
	return nil
}

// Layout -
func (Orch) Layout(width, height int) (int, int) {
	return width, height
}
