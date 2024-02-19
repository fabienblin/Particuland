package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game implements ebiten.Game interface.
type GameEngine struct {
	Image *ebiten.Image
	Particles [][]*Particle
	Species []*ParticleSpecies
	// Qt Quadtree
	// MovedParticles [][]*Particle
}

func (g *GameEngine) InitSpecies(species ...*ParticleSpecies) {
	g.Particles = AllParticleFactory(species...)
	InitInteractions(species...)
	g.Species = species
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *GameEngine) Update() error {
    // Write your game's logical update.
	g.Image.Fill(color.Black)
	g.Particles = UpdateParticles()
	for _, species := range g.Particles {
		for _, p := range species {
			g.Image.Set(int(p.X), int(p.Y), p.Species.Color)
		}
	}

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *GameEngine) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.Image, &ebiten.DrawImageOptions{})
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *GameEngine) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return ImageWidth, ImageHeight
}
