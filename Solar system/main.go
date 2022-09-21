package main

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth, screenHeight = 900, 700
)

var (
	err        	error
	background 	*ebiten.Image
)

var (


	imageTerra 	*ebiten.Image
	imageLava  	*ebiten.Image
	imageIce	*ebiten.Image

	terra planet
	baren planet
	lava planet
	ice planet
)

type planet struct {
	image *ebiten.Image
	radius float64
	rotation float64
	speed float64
	scale float64
	flag int
}

func sizeFloat64(img *ebiten.Image) (x float64, y float64) {
	px, py := img.Size()
	return float64(px), float64(py)
}

func init() {

	background, _, err = ebitenutil.NewImageFromFile("assets/space.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	imageTerra, _, err = ebitenutil.NewImageFromFile("assets/Terran.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	imageLava, _, err = ebitenutil.NewImageFromFile("assets/Lava.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	imageIce, _, err = ebitenutil.NewImageFromFile("assets/Ice.png", ebiten.FilterDefault)


	terra = planet{imageTerra, 100, 0, 10, 0.5, 1}
	lava = planet{imageLava, 250, 0, 6, 0.5, 1}
	ice = planet{imageIce, 320, 0, 4, 0.5, 1}

}

func movePlanet(globe *planet, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(globe.scale, globe.scale)
	op.GeoM.Translate(screenWidth/2 - globe.radius*math.Cos(globe.rotation), screenHeight/2 - globe.radius*math.Sin(globe.rotation))
	screen.DrawImage(globe.image, op)

	if globe.scale <= 0 { 
		globe.flag = 0
	} else if globe.scale >= 1 {
		globe.flag = 1
	}

	if globe.flag == 0 {
		globe.scale += globe.speed/5000
	} else {
		globe.scale -= globe.speed/5000
	}

	globe.rotation += globe.speed/1000
}

func update(screen *ebiten.Image) error {

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(background, op)

	movePlanet(&terra, screen)
	movePlanet(&lava, screen)
	movePlanet(&ice, screen)

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Solar system"); err != nil {
		log.Fatal(err)
	}
}