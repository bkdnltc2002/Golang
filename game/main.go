package main

import (
	"log"
	"math"

	"image/color"


	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

)

type Game struct {
	TextX float64
}

func (g *Game) Update() error {
	g.TextX += 1 // Adjust the speed as needed
	if g.TextX > 640 {
		g.TextX = math.Mod(g.TextX, 640)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	ebitenutil.DebugPrintAt(screen, "Hello, World!", int(g.TextX), 240)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
