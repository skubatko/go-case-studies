package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const distance = 62_100_000

	fmt.Printf("%-20v%10v%14v%7v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("===================================================")

	for i := 0; i < 10; i++ {
		spaceLineType := rand.Intn(4) + 1
		tripType := rand.Intn(2) + 1

		spaceLine := "Space Adventures"
		switch spaceLineType {
		case 2:
			spaceLine = "SpaceX"
		case 3:
			spaceLine = "Virgin Galactic"
		default:
			spaceLine = "Space Adventures"
		}

		velocity := rand.Intn(30-16+1) + 16
		duration := distance / velocity / 24 / 60 / 60

		price := 20 + velocity

		trip := "One-way"
		if tripType == 2 {
			trip = "Round-trip"
			price *= 2
		}

		fmt.Printf("%-20v%10v%15v $ %3v\n", spaceLine, duration, trip, price)
	}
}
