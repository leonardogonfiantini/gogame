package main

import (
	"log"
	"math"
	"image"
	"image/png"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth, screenHeight = 2000, 1400


	frameOX     = 0
	frameOY     = 0
	frameWidth  = 100
	frameHeight = 100
	frameCount  = 50
)

var (
	err        	error
	background 	*ebiten.Image
)

var (
	imageTerra 	*ebiten.Image
	imageSolar 	*ebiten.Image
	imageGiove 	*ebiten.Image
	imageMars 	*ebiten.Image

	terra 		planet
	solar 		star
	giove 		planet
	mars 		planet
)

type planet struct {
	image *ebiten.Image
	radius float64
	orbit, speedorbit float64
	rotation, speedrotation float64
}

type star struct {
	image *ebiten.Image
	rotation, speedrotation float64
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

	imageTerra, _, err = ebitenutil.NewImageFromFile("assets/p1.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	imageSolar, _, err = ebitenutil.NewImageFromFile("assets/Solar.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	imageGiove, _, err = ebitenutil.NewImageFromFile("assets/p2.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	imageMars, _, err = ebitenutil.NewImageFromFile("assets/p3.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	terra = planet{imageTerra, 200, 0, 10, 0, 150}
	giove = planet{imageGiove, 400, 0, 6, 0, 120}
	mars = planet{imageMars, 600, 0, 14, 0, 140}

	solar = star{imageSolar, 0, 150}
}

func movePlanet(globe *planet, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(screenWidth/2 - globe.radius*math.Cos(globe.orbit) - globe.radius/6, screenHeight/2 - globe.radius*math.Sin(globe.orbit) - globe.radius/6)
	i := int(globe.rotation) % frameCount
	sx, sy := i*frameWidth, frameOY
	screen.DrawImage(globe.image.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)

	globe.rotation += globe.speedrotation/1000
	globe.orbit += globe.speedorbit/1000
}

func moveStar(star *star, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(screenWidth/2 - frameWidth, screenHeight/2 - frameHeight)
	i := int(star.rotation) % frameCount
	sx, sy := i*frameWidth*2, frameOY
	screen.DrawImage(star.image.SubImage(image.Rect(sx, sy, sx+200, sy+200)).(*ebiten.Image), op)

	star.rotation += star.speedrotation/1000
}

func update(screen *ebiten.Image) error {

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	op.GeoM.Scale(2,2)
	screen.DrawImage(background, op)

	movePlanet(&terra, screen)
	movePlanet(&giove, screen)
	movePlanet(&mars, screen)
	moveStar(&solar, screen)

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 0.5, "Solar system"); err != nil {
		log.Fatal(err)
	}
}