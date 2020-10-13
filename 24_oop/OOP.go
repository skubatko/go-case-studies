package main

import (
	"fmt"
	"math"
)

// координаты в градусах, минутах, секундах для сферы N/S/E/W
type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	lat, long float64
}

type world struct {
	radius float64
}

// decimal конвертирует координаты d/m/s в десятичные градусы.
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

// newLocation из координат d/m/s широты и долготы.
func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

// rad конвертирует градусы в радианы.
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

// вычисление расстояния через Сферическую теорему косинусов.
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong) // Использует поле радиуса world
}

func main() {
	// Кратер Брэдбери: 4°35'22.2" S, 137°26'30.1" E
	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}

	fmt.Println(lat.decimal(), long.decimal()) // Выводит: -4.5895 137.4417

	// Функции конструктора
	curiosity := newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.12, 'E'})
	fmt.Println(curiosity) // Выводит: {-4.5895 137.4417}

	// Классы в Golang
	var mars = world{radius: 3389.5}

	spirit := location{-14.5684, 175.472636}
	opportunity := location{-1.9462, 354.4734}

	dist := mars.distance(spirit, opportunity) // Использует метод distance для mars
	fmt.Printf("%.2f km\n", dist)              // Выводит: 9669.71 km
}
