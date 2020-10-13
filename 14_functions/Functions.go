package main

import "fmt"

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0
}

func kelvinToFahrenheit(k float64) float64 {
	return celsiusToFahrenheit(kelvinToCelsius(k))
}

func main() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Println(kelvin, "° K is ", celsius, "° C")
	fmt.Printf("%d ° K is %.2f ° C\n", 233, kelvinToCelsius(233))
	fmt.Printf("%d ° C is %.2f ° F\n", 0, celsiusToFahrenheit(0))
	fmt.Printf("%d ° K is %.2f ° F\n", 0, kelvinToFahrenheit(0))
}
