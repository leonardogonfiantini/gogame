package main

import (
	"fmt"
	"log"

	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)


const (
	screenWidth  = 400
	screenHeight = 300

	barWidth = 15
	barHeight = 60
)

var (

	P1pos float64 = 0
	P2pos float64 = 0

	PlayerSpeed float64 = 5
)

type Game struct {
	counter int
}

var (
	keys = []ebiten.Key{
		//player 1
		ebiten.KeyA, //up
		ebiten.KeyS, //down

		//player 2
		ebiten.KeyK, //up
		ebiten.KeyL, //down
	}
)

func (g *Game) Update() error {

	pressed := inpututil.PressedKeys()
	for _, p := range pressed {
		for _, k := range keys {
			if p == k {
				switch k {
				case keys[0]:
					P1pos -= PlayerSpeed
				case keys[1]:
					P1pos += PlayerSpeed
				case keys[2]:
					P2pos -= PlayerSpeed
				case keys[3]:
					P2pos += PlayerSpeed
				}
			}
		}
	}

	g.counter++

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	Player1 := ebiten.NewImage(screenWidth, screenHeight)
	Player2 := ebiten.NewImage(screenWidth, screenHeight)
	Ball := ebiten.NewImage(screenWidth, screenHeight)

	ebitenutil.DrawRect(Player1, 20, screenHeight/2-barHeight/2 + P1pos, barWidth, barHeight, color.White)
	ebitenutil.DrawRect(Player2, screenWidth-40, screenHeight/2-barHeight/2 + P2pos, barWidth, barHeight, color.White)
	ebitenutil.DrawCircle(Ball, screenWidth/2-4, screenHeight/2-4, 4, color.White)

	screen.DrawImage(Player1, nil)	
	screen.DrawImage(Player2, nil)	
	screen.DrawImage(Ball, nil)	


	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Ping Pong")

	g := &Game{
		counter: 0,
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}