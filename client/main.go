package main

import (
	"context"

	"github.com/ethanmil/col/client/game"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/sirupsen/logrus"
)

var art *ebiten.Image
var delta float64
var err error

var ctx context.Context

func init() {
	art, _, err = ebitenutil.NewImageFromFile("assets/art.png", ebiten.FilterDefault)
	if err != nil {
		logrus.Fatal(err)
	}

	ctx = context.Background()
}

func main() {
	// create game object
	col := game.New(art)

	// run our ebiten game
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Colonists")
	ebiten.SetRunnableOnUnfocused(true)

	err = ebiten.RunGame(col)
	if err != nil {
		logrus.Fatal(err)
	}
}
