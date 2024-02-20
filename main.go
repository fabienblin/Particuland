package main

import (
	"flag"
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Config struct {
	G         float64           `mapstructure:"G"`
	ImageSize int               `mapstructure:"imageSize"`
	DeltaTime float64           `mapstructure:"deltaTime"`
	Friction  float64           `mapstructure:"friction"`
	Species   []ParticleSpecies `mapstructure:"species"`
}

var (
	ImageHeight int = 9
	ImageWidth  int = 16
	config      Config
	Game        *GameEngine = &GameEngine{}
	rng         *rand.Rand
)

func init() {
	var configFile = flag.String("config", "./config.json", "The config file.")

	var argSeed = flag.Int("seed", -1, "Seed for recreating a scenario.")
	flag.Parse()

	var seed int64 = int64(*argSeed)
	if *argSeed == -1 {
		seed = rand.Int63()
	}

	readConfig(*configFile)

	rng = rand.New(rand.NewSource(seed))
	log.Printf("Seed %d", seed)

	Game.Image = ebiten.NewImage(ImageWidth, ImageHeight)

	speciesArray := getSpeciesFromConfig()
	Game.InitSpecies(speciesArray)
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

func readConfig(configFile string) {
	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}

	ImageHeight *= config.ImageSize
	ImageWidth *= config.ImageSize
}

func getSpeciesFromConfig() []*ParticleSpecies {
	var speciesArray []*ParticleSpecies

	// Retrieve the array of species from the config
	speciesConfigArray := viper.Get("species").([]interface{})

	// Convert each species configuration to ParticleSpecies
	for _, speciesConfig := range speciesConfigArray {
		var species *ParticleSpecies
		if err := mapstructure.Decode(speciesConfig, &species); err != nil {
			log.Fatalf("Error decoding species config: %v", err)
		}

		color := color.RGBA{species.Color.R, species.Color.G, species.Color.B, 255}
		particleSpecies := &ParticleSpecies{
			Name:              species.Name,
			Color:             color,
			NbParticles:       species.NbParticles,
			Mass:              species.Mass,
			InteractionRadius: species.InteractionRadius,
		}

		speciesArray = append(speciesArray, particleSpecies)
	}

	return speciesArray
}
