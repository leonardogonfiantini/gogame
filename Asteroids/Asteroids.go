package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)


const (
	screenWidth  = 600
	screenHeight = 600
)

var (

)

type Game struct {
	counter int
}

var (	
)

func (g *Game) Update() error {
	g.counter++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	Player := ebiten.NewImage(screenWidth,screenHeight)
	
	ebitenutil.DrawRect(Player, 300-20, 300-20, 40, 40, color.White)

	screen.DrawImage(Player, nil)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Asteroids")

	g := &Game{
		counter: 0,
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}