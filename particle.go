package main

import (
	"image/color"
	"math"
	"math/rand"
)

// Id is auto generated
type ParticleSpecies struct {
	Id                int
	Name              string
	Color             color.RGBA
	NbParticles       int
	InteractionRadius float64
	Mass              float64
	CollisionRadius   float64
}

type Particle struct {
	Species   *ParticleSpecies
	X         float64
	Y         float64
	VelocityX float64
	VelocityY float64
}

func GetParticleInteraction(A *Particle, B *Particle) float64 {
	return Interactions[A.Species.Id][B.Species.Id]
}

func ParticleFactory(species *ParticleSpecies) []*Particle {
	particleLst := []*Particle{}
	for i := 0; i < species.NbParticles; i++ {
		newParticle := &Particle{
			Species:   species,
			X:         rand.Float64() * float64(ImageWidth-1),
			Y:         rand.Float64() * float64(ImageHeight-1),
			VelocityX: 0,
			VelocityY: 0,
		}
		particleLst = append(particleLst, newParticle)
	}

	return particleLst
}

func AllParticleFactory(species ...*ParticleSpecies) [][]*Particle {
	allParticles := [][]*Particle{}
	for i, s := range species {
		s.Id = i
		speciesParticles := ParticleFactory(s)
		allParticles = append(allParticles, speciesParticles)
		s.CollisionRadius = ParticleCollisionRadius
	}

	return allParticles
}

func (p *Particle) RangeSearch() [][]*Particle {
	result := [][]*Particle{}

	for _, species := range Game.Particles {
		resultSpecies := []*Particle{}
		for _, particle := range species {
			if p != particle {
				resultSpecies = append(resultSpecies, particle)
			}
		}
		result = append(result, resultSpecies)
	}

	return result
}

func UpdateParticles() [][]*Particle {
	updatedParticles := [][]*Particle{}
	for i, species := range Game.Particles {
		updatedParticles = append(updatedParticles, []*Particle{})
		for j, particle := range species {
			updatedParticles[i] = append(updatedParticles[i], &Particle{})
			particle.VelocityX /= Friction
			particle.VelocityY /= Friction
			*updatedParticles[i][j] = *particle
		}
	}

	// Force calculation phase
	for i, species := range Game.Particles {
		for j, subject := range species {
			var fx, fy float64

			// Accumulate forces from all neighbors
			for _, species := range Game.Particles {
				for _, neighbour := range species {
					if subject == neighbour {
						continue
					}

					distance, dx, dy := distance(subject, neighbour)

					g := G
					if distance < ParticleCollisionRadius {
						g = -g
					} else if distance > subject.Species.InteractionRadius {
						g = 0
					}

					force := (GetParticleInteraction(subject, neighbour) * g * subject.Species.Mass * neighbour.Species.Mass) / (distance * distance)
					fx += force * dx
					fy += force * dy
				}
			}

			// Update phase
			update := updatedParticles[i][j]

			update.VelocityX += fx * InertiaFactor * DeltaTime
			update.VelocityY += fy * InertiaFactor * DeltaTime

			update.X += update.VelocityX * DeltaTime
			update.Y += update.VelocityY * DeltaTime

			handleBorderCollision(update)
		}
	}

	return updatedParticles
}

func distance(A, B *Particle) (float64, float64, float64) {
	dx := B.X - A.X
	dy := B.Y - A.Y

	return math.Sqrt(dx*dx + dy*dy), dx, dy
}

func handleBorderCollision(p *Particle) {
	if p.X < 0 {
		p.X = -p.X
		p.VelocityX = -p.VelocityX
	} else if p.X > float64(ImageWidth) {
		p.X = float64(ImageWidth) - (p.X - float64(ImageWidth))
		p.VelocityX = -p.VelocityX
	}

	if p.Y < 0 {
		p.Y = -p.Y
		p.VelocityY = -p.VelocityY
	} else if p.Y > float64(ImageHeight) {
		p.Y = float64(ImageHeight) - (p.Y - float64(ImageHeight))
		p.VelocityY = -p.VelocityY
	}
}
