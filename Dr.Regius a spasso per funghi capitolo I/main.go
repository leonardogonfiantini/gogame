package main

import (
	"log"
	"math/rand"
	"time"
	"container/list"
	"fmt"


	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Our game constants
const (
	screenWidth, screenHeight = 800, 600
	characterCentering = 32
	mushroomSize = 5
	catchSize = 32
)

// Create our empty vars
var (
	err        error
	background *ebiten.Image
	character  *ebiten.Image
	edulis *ebiten.Image
	playerOne  player
	mushrooms  *list.List
	points int
)

// Create the player class
type player struct {
	image      *ebiten.Image
	xPos, yPos float64
	speed      float64
}

type mushroom struct {
	image *ebiten.Image
	xPos, yPos float64
}

// Run this code once at startup
func init() {
	rand.Seed(time.Now().UnixNano())

	background, _, err = ebitenutil.NewImageFromFile("assets/caucaso.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	character, _, err = ebitenutil.NewImageFromFile("assets/me.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	edulis, _, err = ebitenutil.NewImageFromFile("assets/mushroom.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	playerOne = player{character, (screenWidth / 2.0) - characterCentering, (screenHeight / 2.0) - characterCentering, 10}
	mushrooms = list.New()

}

func createMushroom() mushroom {
	return mushroom{edulis, float64(rand.Intn(screenWidth-50)), float64(rand.Intn(screenHeight-50))}
}

func addMushroom() {
	if mushrooms.Len() < mushroomSize {
		mushrooms.PushBack(createMushroom())
	}
}

func updateMushrooms(screen *ebiten.Image) {
	addMushroom()

	for e := mushrooms.Front(); e != nil; e = e.Next() {
		m := mushroom(e.Value.(mushroom))
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(m.xPos), float64(m.yPos))
		screen.DrawImage(m.image, op)
	}

	for e := mushrooms.Front(); e != nil; e = e.Next() {
		m := mushroom(e.Value.(mushroom))

		if (playerOne.xPos <= m.xPos + catchSize && playerOne.xPos >= m.xPos - catchSize) &&
			(playerOne.yPos <= m.yPos + catchSize && playerOne.yPos >= m.yPos - catchSize) {
				mushrooms.Remove(e)
				points += 20
			}
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", points))
}

// Move the player depending on which key is pressed
func movePlayer() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if val := playerOne.yPos - playerOne.speed; val >= 0 {
			playerOne.yPos -= playerOne.speed
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if val := playerOne.yPos + playerOne.speed; val <= screenHeight - characterCentering {
			playerOne.yPos += playerOne.speed
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if val := playerOne.xPos + playerOne.speed; val >= 0 {
			playerOne.xPos -= playerOne.speed
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if val := playerOne.xPos + playerOne.speed; val <= screenWidth - characterCentering {
			playerOne.xPos += playerOne.speed
		}
	}
}

func update(screen *ebiten.Image) error {
	movePlayer()

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(background, op)

	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Translate(playerOne.xPos, playerOne.yPos)
	screen.DrawImage(playerOne.image, playerOp)

	updateMushrooms(screen)

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}