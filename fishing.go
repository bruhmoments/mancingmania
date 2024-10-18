package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func goFishing(pole int, baits map[string]int, forecast Forecast) int {
	totalEarnings := 0

	for _, count := range baits {
		for i := 0; i < count; i++ {
			fishSize := getRandomFishSize(pole, forecast)
			fishColor := getRandomFishColor(forecast)

			earnings := getFishEarnings(fishColor, fishSize)
			totalEarnings += earnings

			fishAnimation(fishColor, fishSize, earnings)

			// Phew catch a break
			fmt.Println("Press enter to continue...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		}
	}

	fmt.Printf("You earned a total of %d gold today!\n", totalEarnings)
	return totalEarnings
}

func getRandomFishSize(pole int, forecast Forecast) string {
	switch pole {
	case SmallPole:
		if forecast.SmallFish > 0 {
			forecast.SmallFish--
			return "small"
		}
	case MediumPole:
		if forecast.MediumFish > 0 {
			forecast.MediumFish--
			return "medium"
		}
	case BigPole:
		if forecast.BigFish > 0 {
			forecast.BigFish--
			return "big"
		}
	}
	return "small"
}

func getRandomFishColor(f Forecast) string {
	r := rand.Intn(100) + 1
	if r <= f.RedPercent {
		return "red"
	} else if r <= f.RedPercent+f.BluePercent {
		return "blue"
	}
	return "green"
}

func getFishEarnings(color, size string) int {
	switch color {
	case "red":
		switch size {
		case "small":
			return rand.Intn(5) + 1
		case "medium":
			return rand.Intn(6) + 5
		case "big":
			return rand.Intn(6) + 10
		}
	case "blue":
		switch size {
		case "small":
			return rand.Intn(3) + 3
		case "medium":
			return rand.Intn(3) + 8
		case "big":
			return rand.Intn(3) + 13
		}
	case "green":
		switch size {
		case "small":
			return 5
		case "medium":
			return 10
		case "big":
			return 15
		}
	}
	return 0
}

// Animasi Mancing
func displayFishingScene(fishPosition int) {
	clearScreen()
	fmt.Println("|")
	fmt.Println("|")
	fmt.Println("~~~~~~~~~~~~~~")
	fmt.Println("|")
	fmt.Println("|")
	fmt.Print("|/")
	moveCursor(3+fishPosition, 6)
	fmt.Print("<><")
}

func fishAnimation(fishColor string, fishSize string, earnings int) {
	fishStartPosition := 10

	for i := fishStartPosition; i > 1; i-- {
		displayFishingScene(i)
		delay(100)
	}

	displayFishingScene(0)
	fmt.Println("\nStrike!")
	fmt.Println("Mancing mania mantap!")
	fmt.Printf("Caught a %s %s fish for %d gold.\n", fishColor, fishSize, earnings)
}
