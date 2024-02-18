package main

import (
	"image/color"
	"log"

	// "math"
	// "time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	G                       float64 = 1 // 6.67430e-11
	imageSize               int     = 70
	ImageHeight             int     = imageSize * 9
	ImageWidth              int     = imageSize * 16
	InertiaFactor           float64 = 0.5
	ParticleCollisionRadius float64 = 5
	DeltaTime               float64 = 1
)
type GameEngine struct {
	Particles [][]*Particle
	Species   []*ParticleSpecies
}

func (g *GameEngine) InitSpecies(species ...*ParticleSpecies) {
	g.Particles = AllParticleFactory(species...)
	InitInteractions(species...)
	g.Species = species
}

var Game *GameEngine = &GameEngine{}

func init() {
	// Game.Image = ebiten.NewImage(ImageWidth, ImageHeight)

	RedSpecies := &ParticleSpecies{Name: "Red", Color: color.RGBA{238, 36, 39, 255}, NbParticles: 200, Mass: 1}
	BlueSpecies := &ParticleSpecies{Name: "Blue", Color: color.RGBA{139, 141, 255, 255}, NbParticles: 200, Mass: 1}
	GreenSpecies := &ParticleSpecies{Name: "Green", Color: color.RGBA{78, 202, 58, 255}, NbParticles: 200, Mass: 1}

	Game.InitSpecies(RedSpecies, BlueSpecies, GreenSpecies)

	SetInteraction(GreenSpecies, GreenSpecies, 1)
	SetInteraction(GreenSpecies, RedSpecies, 0)
	SetInteraction(GreenSpecies, BlueSpecies, 0.5)

	SetInteraction(RedSpecies, GreenSpecies, -0.5)
	SetInteraction(RedSpecies, RedSpecies, 1)
	SetInteraction(RedSpecies, BlueSpecies, 0)

	SetInteraction(BlueSpecies, GreenSpecies, -0.5)
	SetInteraction(BlueSpecies, RedSpecies, 0)
	SetInteraction(BlueSpecies, BlueSpecies, -1)
}

func main() {
	PrintInteractions()
	// Specify the window size as you like. Here, a doubled size is specified.
	// ebiten.SetWindowSize(ImageWidth, ImageHeight)
	// ebiten.SetWindowTitle("Particuland")
	// ebiten.SetFullscreen(true)
	// ebiten.SetTPS(25)
	// Call ebiten.RunGame to start your game loop.
	// if err := ebiten.RunGame(Game); err != nil {
	// 	log.Fatal(err)
	// }

	//

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		log.Println("Error initializing SDL:", err)
		return
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Gravity Simulator", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(ImageWidth), int32(ImageHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		log.Println("Error creating window:", err)
		return
	}
	defer window.Destroy()

	// window.SetFullscreen(sdl.WINDOW_FULLSCREEN)

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Println("Error creating renderer:", err)
		return
	}
	defer renderer.Destroy()

	for {
		// startTime := time.Now()
		Game.Particles = UpdateParticles()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		for _, species := range Game.Particles {
			for _, particle := range species {
				r, g, b, a := particle.Species.Color.RGBA()
				renderer.SetDrawColor(uint8(r), uint8(g), uint8(b), uint8(a))
				renderer.DrawPoint(int32(particle.X), int32(particle.Y))
			}
		}

		renderer.Present()

		// elapsed := time.Since(startTime)
		// time.Sleep(time.Duration(int64(math.Max(0, DeltaTime*1000))-elapsed.Milliseconds()) * time.Millisecond)
	}
}
