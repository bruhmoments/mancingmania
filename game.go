package main

import (
	"fmt"
	"math/rand"
)

const (
	SmallPole  = 5
	MediumPole = 10
	BigPole    = 15
)

type Forecast struct {
	SmallFish    int
	MediumFish   int
	BigFish      int
	RedPercent   int
	BluePercent  int
	GreenPercent int
}

func generateForecast() Forecast {
	smallFish := rand.Intn(10) + 1
	mediumFish := rand.Intn(5) + 1
	bigFish := rand.Intn(5) + 1

	redPercent := rand.Intn(50) + 1
	bluePercent := rand.Intn(100-redPercent) + 1
	greenPercent := 100 - redPercent - bluePercent

	return Forecast{smallFish, mediumFish, bigFish, redPercent, bluePercent, greenPercent}
}

func displayForecast(f Forecast) {
	fmt.Printf("\nToday's forecast:\nSmall fish: %d\nMedium fish: %d\nBig fish: %d\n", f.SmallFish, f.MediumFish, f.BigFish)
	fmt.Printf("Fish color distribution - Red: %d%%, Blue: %d%%, Green: %d%%\n", f.RedPercent, f.BluePercent, f.GreenPercent)
}

func chooseFishingPole(playerGold *int) int {
	fmt.Println("\nChoose a fishing pole:")
	fmt.Println("1. Small Fishing Pole - 5 gold")
	fmt.Println("2. Medium Fishing Pole - 10 gold")
	fmt.Println("3. Big Fishing Pole - 15 gold")
	var choice int
	for {
		fmt.Print("Enter your choice (1-3): ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			if *playerGold >= SmallPole {
				*playerGold -= SmallPole
				return SmallPole
			}
		case 2:
			if *playerGold >= MediumPole {
				*playerGold -= MediumPole
				return MediumPole
			}
		case 3:
			if *playerGold >= BigPole {
				*playerGold -= BigPole
				return BigPole
			}
		}
		fmt.Println("Can't select this choice. Please select again.")
	}
}

func buyBaits(playerGold *int) map[string]int {
	baits := map[string]int{"red": 0, "blue": 0, "green": 0}
	fmt.Println("\nBuy baits:")
	for *playerGold > 0 {
		fmt.Printf("\nYou have %d gold left. Choose bait to buy:\n", *playerGold)
		fmt.Println("1. Red bait - 1 gold")
		fmt.Println("2. Blue bait - 2 gold")
		fmt.Println("3. Green bait - 3 gold")
		fmt.Println("4. Done buying")
		var choice int
		fmt.Print("\nEnter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			if *playerGold >= 1 {
				baits["red"]++
				*playerGold -= 1
			}
		case 2:
			if *playerGold >= 2 {
				baits["blue"]++
				*playerGold -= 2
			}
		case 3:
			if *playerGold >= 3 {
				baits["green"]++
				*playerGold -= 3
			}
		case 4:
			return baits
		default:
			fmt.Println("\nInvalid choice. Please select again.")
		}
	}
	return baits
}
