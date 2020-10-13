package main

import "fmt"

type celsius float64
type kelvin float64

func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}

func main() {
	fmt.Printf("%d° C is %.2f° K\n", 127, celsiusToKelvin(127))
}
