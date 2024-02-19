package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	G                       float64 = 1 // 6.67430e-11
	imageSize               int     = 30
	ImageHeight             int     = imageSize * 9
	ImageWidth              int     = imageSize * 16
	InertiaFactor           float64 = 0.5
	ParticleCollisionRadius float64 = 10 + (float64(imageSize) / 10)
	DeltaTime               float64 = 0.8
	Friction                float64 = 1.01
)

var Game *GameEngine = &GameEngine{}

func init() {
	seed := rand.Int63()
	rand.New(rand.NewSource(seed))
	log.Printf("Seed %d", seed)

	Game.Image = ebiten.NewImage(ImageWidth, ImageHeight)

	RedSpecies := &ParticleSpecies{Name: "Red", Color: color.RGBA{238, 36, 39, 255}, NbParticles: 100, Mass: 1, InteractionRadius: 100}
	BlueSpecies := &ParticleSpecies{Name: "Blue", Color: color.RGBA{139, 141, 255, 255}, NbParticles: 100, Mass: 1, InteractionRadius: 100}
	GreenSpecies := &ParticleSpecies{Name: "Green", Color: color.RGBA{78, 202, 58, 255}, NbParticles: 50, Mass: 1, InteractionRadius: 100}

	Game.InitSpecies(RedSpecies, BlueSpecies, GreenSpecies)

	// SetInteraction(GreenSpecies, GreenSpecies, -0.40)
	// SetInteraction(GreenSpecies, RedSpecies, 0.22)
	// SetInteraction(GreenSpecies, BlueSpecies, 0.40)

	// SetInteraction(RedSpecies, GreenSpecies, -0.93)
	// SetInteraction(RedSpecies, RedSpecies, -0.28)
	// SetInteraction(RedSpecies, BlueSpecies, -0.73)

	// SetInteraction(BlueSpecies, GreenSpecies, 0.63)
	// SetInteraction(BlueSpecies, RedSpecies, -0.86)
	// SetInteraction(BlueSpecies, BlueSpecies, 0.27)
}

func main() {
	PrintInteractions()

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(ImageWidth, ImageHeight)
	ebiten.SetWindowTitle("Particuland")
	ebiten.SetFullscreen(true)
	ebiten.SetTPS(25)

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(Game); err != nil {
		log.Fatal(err)
	}
}
