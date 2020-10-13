package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var piggyBank float64

	for i := 0; i < 11; i++ {
		piggyBank += 0.10
	}

	fmt.Printf("%v\n", piggyBank)

	var sum float64
	for sum < 20.0 {
		switch rand.Intn(3) {
		case 0:
			sum += 0.05
		case 1:
			sum += 0.10
		case 2:
			sum += 0.25
		}
		fmt.Printf("$ %5.2f\n", sum)
	}

	var bank = 0
	for bank < 2500 {
		switch rand.Intn(3) {
		case 0:
			bank += 5
		case 1:
			bank += 10
		case 2:
			bank += 25
		}
		fmt.Printf("$ %d.%02d\n", bank/100, bank%100)
	}

	const distance = 236_000_000_000_000_000
	const lightSpeed = 299792
	const secondsPerDay = 86400
	const daysPerYear = 365

	const years = distance / lightSpeed / secondsPerDay / daysPerYear

	fmt.Println("Расстояние в световых годах до Карликовой галактики в Большом Псе составляет:", years)
}
