package main

import "fmt"

type report struct {
	sol         int
	temperature temperature // Поле temperature является структурой типа temperature
	location    location
}

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type celsius float64

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

type sol int

type embeddedReport struct {
	sol
	temperature // Тип temperature встроен в report
	location
}

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}

func main() {
	bradbury := location{-4.5895, 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	bradburyReport := report{sol: 15, temperature: t, location: bradbury}

	fmt.Printf("%+v\n", bradburyReport) // Выводит: {sol:15 temperature:{high:-1 low:-78} location:{lat:-4.5895 long:137.4417}}

	fmt.Printf("a balmy %v° C\n", bradburyReport.temperature.high) // Выводит: a balmy -1° C

	t = temperature{high: -1.0, low: -78.0}
	fmt.Printf("average %v° C\n", t.average()) // Выводит: -39.5° C

	reportPlus := embeddedReport{
		sol:         15,
		location:    location{-4.5895, 137.4417},
		temperature: temperature{high: -1.0, low: -78.0},
	}

	fmt.Printf("average %v° C\n", reportPlus.average())             // Выводит: average -39.5° C
	fmt.Printf("average %vº C\n", reportPlus.temperature.average()) // Выводит: average -39.5° C
	fmt.Printf("%v° C\n", reportPlus.high)                          // Выводит: -1° C
	reportPlus.high = 32
	fmt.Printf("%v° C\n", reportPlus.temperature.high) // Выводит: 32° C

	embeddedReport := embeddedReport{sol: 15}

	fmt.Println(embeddedReport.sol.days(1446)) // Выводит: 1431
	fmt.Println(embeddedReport.days(1446))     // Выводит: 1431
}
