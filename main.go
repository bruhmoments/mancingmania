package main

import (
	"fmt"
)

func main() {
	// source := rand.NewSource(time.Now().UnixNano())
	// rng := rand.New(source)

	// Generate a random number of minutes between 1 and 15
	// randomNumbers := rng.Intn(15) + 1

	fmt.Println("Welcome to the Fishing Game!")
	delay(1000)

	day := 1
	playerGold := 100

	for {
		// fmt.Printf("\nYou have %d gold.\n", playerGold)
		showStats(day, playerGold)

		forecast := generateForecast()
		displayForecast(forecast)

		fmt.Println("\nDo you want to skip this day?")
		fmt.Println("1. Yes, skip the day")
		fmt.Println("2. No, continue fishing")
		var skipChoice int
		fmt.Print("Enter your choice (1-2): ")
		fmt.Scan(&skipChoice)

		if skipChoice == 1 {
			fmt.Println("You chose to skip the day.")
			day++
			continue
		} else if skipChoice != 2 {
			fmt.Println("Invalid choice. Skipping the day by default.")
			day++
			continue
		}

		pole := chooseFishingPole(&playerGold)

		baits := buyBaits(&playerGold)

		totalEarnings := goFishing(pole, baits, forecast)

		playerGold += totalEarnings

		if playerGold > 100 {
			fmt.Println("Congratulations! You finished with more than 100 gold. You win!")
		} else if playerGold <= 100 {
			fmt.Println("You finished with less than 100 gold. You lose!")
		} else {
			fmt.Println("You ended the day with 100 gold. It's a tie!")
		}

		fmt.Println("Do you want to play again? (y/n)")
		var choice string
		fmt.Scan(&choice)

		if choice != "y" {
			fmt.Println("Thanks for playing! Goodbye!")
			break
		}

		day++
	}
}

func showStats(day int, playerGold int) {
	fmt.Printf("\nDay %d, You have %d gold.\n", day, playerGold)
}
