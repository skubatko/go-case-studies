package main

import "fmt"

type celsius float64
type fahrenheit float64

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}
func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit(c/5.0*9.0 + 32.0)
}

func drawSeparator() {
	fmt.Println("=======================")
}

func drawHeader(head1 string, head2 string) {
	drawSeparator()
	fmt.Printf("| %8s | %8s |\n", head1, head2)
	drawSeparator()
}

func drawTable(dataFn func() (float64, float64)) {
	row1, row2 := dataFn()
	fmt.Printf("| %8.1f | %8.1f |\n", row1, row2)
}

func main() {
	drawHeader("°C", "°F")
	for i := 40; i <= 100; i += 5 {
		drawTable(func() (float64, float64) {
			return float64(i), float64(celsius(i).fahrenheit())
		})
	}
	drawSeparator()

	drawHeader("°F", "°C")
	for i := 40; i <= 100; i += 5 {
		drawTable(func() (float64, float64) {
			return float64(i), float64(fahrenheit(i).celsius())
		})
	}
	drawSeparator()
}
