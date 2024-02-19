package main

import (
	"image/color"
	"log"
	"math/rand"

	"flag"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	G                       float64 = 1 // 6.67430e-11
	imageSize               int     = 50
	ImageHeight             int     = imageSize * 9
	ImageWidth              int     = imageSize * 16
	InertiaFactor           float64 = 0.5
	ParticleCollisionRadius float64 = 10
	DeltaTime               float64 = 1
	Friction                float64 = 0.2
)

var Game *GameEngine = &GameEngine{}
var rng *rand.Rand

func init() {
	var argSeed = flag.Int("seed", -1, "Seed for recreating a scenario.")
	flag.Parse()

	var seed int64 = int64(*argSeed)
	if *argSeed == -1 {
		seed = rand.Int63()
	}

	rng = rand.New(rand.NewSource(seed))
	log.Printf("Seed %d", seed)

	Game.Image = ebiten.NewImage(ImageWidth, ImageHeight)

	RedSpecies := &ParticleSpecies{Name: "Red", Color: color.RGBA{238, 36, 39, 255}, NbParticles: 100, Mass: 1, InteractionRadius: 60}
	BlueSpecies := &ParticleSpecies{Name: "Blue", Color: color.RGBA{139, 141, 255, 255}, NbParticles: 100, Mass: 1, InteractionRadius: 60}
	GreenSpecies := &ParticleSpecies{Name: "Green", Color: color.RGBA{50, 200, 50, 255}, NbParticles: 100, Mass: 1, InteractionRadius: 60}
	YellowSpecies := &ParticleSpecies{Name: "Yellow", Color: color.RGBA{202, 200, 0, 255}, NbParticles: 100, Mass: 1, InteractionRadius: 60}
	WhiteSpecies := &ParticleSpecies{Name: "White", Color: color.RGBA{255, 255, 255, 255}, NbParticles: 100, Mass: 1, InteractionRadius: 60}

	Game.InitSpecies(RedSpecies, BlueSpecies, GreenSpecies, YellowSpecies, WhiteSpecies)

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
