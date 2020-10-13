package main

import "fmt"

type kelvinT float64
type celsiusT float64
type fahrenheitT float64

func kelvinToCelsiusPlus(k kelvinT) celsiusT { // функция kelvinToCelsius
	return celsiusT(k - 273.15)
}

func (k kelvinT) celsius() celsiusT { // метод celsius для типа kelvin
	return celsiusT(k - 273.15)
}

func (k kelvinT) fahrenheit() fahrenheitT {
	return k.celsius().fahrenheit()
}

func (f fahrenheitT) celsius() celsiusT {
	return celsiusT((f - 32.0) * 5.0 / 9.0)
}

func (f fahrenheitT) kelvin() kelvinT {
	return f.celsius().kelvin()
}

func (c celsiusT) fahrenheit() fahrenheitT {
	return fahrenheitT(c/5.0*9.0 + 32.0)
}

func (c celsiusT) kelvin() kelvinT {
	return kelvinT(c + 273.15)
}

func main() {
	var k kelvinT = 294.0

	fmt.Printf("%.2f\n", kelvinToCelsiusPlus(k))
	fmt.Printf("%.2f\n", k.celsius())
	fmt.Printf("%.2f\n", k.fahrenheit())
	fmt.Printf("%.2f\n", k.fahrenheit().celsius().kelvin())
}
