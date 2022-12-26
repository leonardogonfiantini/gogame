package main

import (
	"fmt"
	"log"
	"math"

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

	ballSize = 10
)

var (
	p1score = 0
	p2score = 0

	P1pos float64 = 0
	P2pos float64 = 0
	PlayerSpeed float64 = 6

	ball_pos_x float64 = screenWidth / 2
	ball_pos_y float64 = screenHeight / 2
	
	ball_dir_x float64 = -1
	ball_dir_y float64 = 0
	
	ball_speed float64 = 4
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

func vec2_norm() {
	// sets a vectors length to 1 (which means that x + y == 1)
	length := math.Sqrt((ball_dir_x * ball_dir_x) + (ball_dir_y * ball_dir_y));
	if (length != 0.0) {
		length = 1.0 / length;
		ball_dir_x *= length;
		ball_dir_y *= length;
	}
}

func updateBall() {
	// fly a bit
    ball_pos_x += ball_dir_x * ball_speed
    ball_pos_y += ball_dir_y * ball_speed

    //left 
    if (ball_pos_x < (20 + barWidth) && ball_pos_x > 20 &&
        ball_pos_y < (screenHeight/2+P1pos + barHeight/2) && ball_pos_y > (screenHeight/2+P1pos - barHeight/2)) {

        t := ((ball_pos_y - screenHeight/2+P1pos) / barHeight) - 0.5
        ball_dir_x = math.Abs(ball_dir_x)
		ball_dir_y = t
    }

	//right
    if (ball_pos_x < (screenWidth-40 + barWidth) && ball_pos_x > (screenWidth-40) && 
        ball_pos_y < (screenHeight/2+P2pos + barHeight/2) && ball_pos_y > (screenHeight/2+P2pos - barHeight/2)) {
        
        t := ((ball_pos_y - -screenHeight/2+P2pos) / barHeight) - 0.5
        ball_dir_x = -math.Abs(ball_dir_x)
        ball_dir_y = t;
    }

    //left wall
    if (ball_pos_x < 0) {
        p2score++
        ball_pos_x = screenWidth / 2
        ball_pos_y = screenHeight / 2
        ball_dir_x = math.Abs(ball_dir_x);
        ball_dir_y = 0
    }

    //right wall
    if (ball_pos_x > screenWidth) {
		p1score++
        ball_pos_x = screenWidth / 2
        ball_pos_y = screenHeight / 2
        ball_dir_x = -math.Abs(ball_dir_x)
        ball_dir_y = 0
    }

    //top wall
    if (ball_pos_y > screenHeight) {
        ball_dir_y = -math.Abs(ball_dir_y)
    }

    //bottom wall
    if (ball_pos_y < 0) {
        ball_dir_y = math.Abs(ball_dir_y)
	}

    vec2_norm();
}

func (g *Game) Update() error {

	pressed := inpututil.PressedKeys()
	for _, p := range pressed {
		for _, k := range keys {
			if p == k {
				switch k {
				case keys[0]:
					if (P1pos - PlayerSpeed) >= -(screenHeight/2-barHeight/2) {
						P1pos -= PlayerSpeed
					}
				case keys[1]:
					if (P1pos + PlayerSpeed) <= (screenHeight/2-barHeight/2) {
						P1pos += PlayerSpeed
					}
				case keys[2]:
					if (P2pos - PlayerSpeed) >= -(screenHeight/2-barHeight/2) {
						P2pos -= PlayerSpeed
					}
				case keys[3]:
					if (P2pos + PlayerSpeed) <= (screenHeight/2-barHeight/2) {
						P2pos += PlayerSpeed
					}
				}
			}
		}
	}


	updateBall()

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
	ebitenutil.DrawRect(Ball, ball_pos_x - ballSize/2, ball_pos_y - ballSize/2, ballSize, ballSize, color.White)

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