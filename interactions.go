package main

import (
	"fmt"
)

var Interactions [][]float64

// Create table of particle interactions, a square of
func InitInteractions(species ...*ParticleSpecies) {
	Interactions = make([][]float64, len(species))
	for i := range species {
		Interactions[i] = make([]float64, len(species))
	}

	for _, sA := range species {
		for _, sB := range species {
			SetInteraction(sA, sB, RandomInteraction())
		}
	}
}

func SetInteraction(A *ParticleSpecies, B *ParticleSpecies, force float64) {
	Interactions[A.Id][B.Id] = force
}

func GetInteraction(A *ParticleSpecies, B *ParticleSpecies) float64 {
	return Interactions[A.Id][B.Id]
}

const MaxInteraction float64 = 10

func RandomInteraction() float64 {
	return ((rng.Float64() * 2) - 1)
}

func PrintInteractions() {
	// Print column headers (species names)
	fmt.Printf("%-15s", "") // Empty cell for top-left corner
	for i := 0; i < len(Game.Particles); i++ {
		fmt.Printf("%-15s", fmt.Sprintf(Game.Particles[i][0].Species.Name))
	}
	fmt.Println()

	// Print rows with species names and interaction values
	for i, row := range Interactions {
		fmt.Printf("%-15s", fmt.Sprintf(Game.Particles[i][0].Species.Name))
		for _, val := range row {
			fmt.Printf("%-15.2f", val)
		}
		fmt.Println()
	}
}
