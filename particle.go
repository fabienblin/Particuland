package main

import (
	"image/color"
	"math"
)

// Id is auto generated
type ParticleSpecies struct {
	Id                int
	Name              string `mapstructure:"Name"`
	Color             color.RGBA
	NbParticles       int     `mapstructure:"NbParticles"`
	InteractionRadius float64 `mapstructure:"InteractionRadius"`
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
			X:         rng.Float64() * float64(ImageWidth-1),
			Y:         rng.Float64() * float64(ImageHeight-1),
			VelocityX: 0,
			VelocityY: 0,
		}
		particleLst = append(particleLst, newParticle)
	}

	return particleLst
}

func AllParticleFactory(species []*ParticleSpecies) [][]*Particle {
	allParticles := [][]*Particle{}
	for i, s := range species {
		s.Id = i
		speciesParticles := ParticleFactory(s)
		allParticles = append(allParticles, speciesParticles)
	}

	return allParticles
}

func UpdateParticles() [][]*Particle {
	updatedParticles := [][]*Particle{}
	for i, species := range Game.Particles {
		updatedParticles = append(updatedParticles, []*Particle{})
		for j, particle := range species {
			updatedParticles[i] = append(updatedParticles[i], &Particle{})
			particle.VelocityX -= config.Friction * particle.VelocityX
			particle.VelocityY -= config.Friction * particle.VelocityY
			*updatedParticles[i][j] = *particle
		}
	}

	// Force calculation phase
	for i := range Game.Particles {
		for j, subject := range Game.Particles[i] {
			var fx, fy float64

			// Accumulate forces from all neighbors
			for k := range Game.Particles {
				for _, neighbour := range Game.Particles[k] {
					if subject == neighbour {
						continue
					}

					distance, dx, dy := distance(subject, neighbour)

					force := 0.0
					if distance < config.RepelDistance {
						force = -math.Abs(GetParticleInteraction(subject, neighbour)) / (distance * distance)
					} else if distance > subject.Species.InteractionRadius {
						continue
					} else {
						force = GetParticleInteraction(subject, neighbour) / (distance * distance)
					}

					fx += force * dx
					fy += force * dy
				}
			}

			// Update phase
			update := updatedParticles[i][j]

			update.VelocityX += fx * config.DeltaTime
			update.VelocityY += fy * config.DeltaTime

			update.X += update.VelocityX
			update.Y += update.VelocityY

			handleBorderCollision(update)
		}
	}

	return updatedParticles
}

func distance(A, B *Particle) (float64, float64, float64) {
    dx := B.X - A.X
    dy := B.Y - A.Y

    if dx > float64(ImageWidth)/2 {
        dx -= float64(ImageWidth)
    } else if dx < -float64(ImageWidth)/2 {
        dx += float64(ImageWidth)
    }

    if dy > float64(ImageHeight)/2 {
        dy -= float64(ImageHeight)
    } else if dy < -float64(ImageHeight)/2 {
        dy += float64(ImageHeight)
    }

    return math.Sqrt(dx*dx + dy*dy), dx, dy
}

func handleBorderCollision(p *Particle) {
    if p.X < 0 {
        p.X += float64(ImageWidth)
    } else if p.X > float64(ImageWidth) {
        p.X -= float64(ImageWidth)
    }

    if p.Y < 0 {
        p.Y += float64(ImageHeight)
    } else if p.Y > float64(ImageHeight) {
        p.Y -= float64(ImageHeight)
    }
}
