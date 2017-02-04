package main

import (
	"fmt"
	"math/rand"
	"models"
)

const NUM_PROTISTS = 100
const NUM_CYCLES = 1000
const MAX_PARAM = 100

var protists [NUM_PROTISTS]*models.Protist

// var conditions = []func(*models.Protist, int)bool{ isHealthAbove, isFoodAbove, isCold, isHot }
// var actions = []func(*models.Protist, int){ actionCover, actionUncover, actionEat }

func generateProtist(num int, environment *models.Environment) *models.Protist {
	newPro := &models.Protist{
		Id:          num,
		Health:      100,
		Food:        100,
		Days_lived:  0,
		Covered:     false,
		Alive:       true,
		Environment: environment,
	}
	newPro.GenerateDNA()
	newPro.GenerateActionFromDNA()
	return newPro
}

func generateEnvironment() *models.Environment {
	startingTemperature := rand.Intn(101)
	startingNumDead := 0
	environment := &models.Environment{
		Temperature: startingTemperature,
		GoodWeather: 0,
		BadWeather:  0,
		NumDead:     startingNumDead,
	}
	return environment
}

func main() {
	rand.Seed(10)

	environment := generateEnvironment()

	// create a ton of random protists
	for i := 0; i < NUM_PROTISTS; i++ {
		newPro := generateProtist(i, environment)
		protists[i] = newPro
	}

	for cycle := 0; cycle < NUM_CYCLES; cycle++ {
		fmt.Println("\n\nDay", cycle+1)
		environment.UpdateEnvironment()
		for _, p := range protists {
			p.DoCycle()
		}
		fmt.Println("\nStill alive: ")
		for _, p := range protists {
			if p.Alive {
				fmt.Print(" models.Protist ", p.Id, ", ")
			}
		}
		if environment.NumDead >= NUM_PROTISTS {
			cycle = NUM_CYCLES
		}
	}
	fmt.Println("\nDays of Bad weather: ", environment.BadWeather)
	fmt.Println("Days of Good weather: ", environment.GoodWeather)
	for i, p := range protists {
		fmt.Println("models.Protist", i, "\tDNA: ", p.Dna, "\tDays lived: ", p.Days_lived)
	}

	// print details of last surviving protists
}
