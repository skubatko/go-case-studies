package main

import (
	"fmt"
	"math"
	"sort"
)

type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

type location struct {
	Module string     `json:"module"`
	Place  string     `json:"place"`
	Lat    coordinate `json:"latitude"`
	Long   coordinate `json:"longitude"`
}

type world struct {
	radius float64
}

// rad конвертирует градусы в радианы.
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

// вычисление расстояния через Сферическую теорему косинусов.
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.Lat.decimal()))
	s2, c2 := math.Sincos(rad(p2.Lat.decimal()))
	clong := math.Cos(rad(p1.Long.decimal() - p2.Long.decimal()))
	return w.radius * math.Acos(s1*s2+c1*c2*clong) // Использует поле радиуса world
}

type way struct {
	start location
	end   location
}

func main() {
	locations := []location{
		{Module: "Spirit", Place: "Columbia Memorial Station", Lat: coordinate{14, 34, 6.2, 'S'}, Long: coordinate{175, 28, 21.5, 'E'}},
		{Module: "Opportunity", Place: "Challenger Memorial Station", Lat: coordinate{1, 56, 46.3, 'S'}, Long: coordinate{354, 28, 24.2, 'E'}},
		{Module: "Curiosity", Place: "Bradbury Landing", Lat: coordinate{4, 35, 22.2, 'S'}, Long: coordinate{137, 26, 30.1, 'E'}},
		{Module: "InSight", Place: "Elysium Planitia", Lat: coordinate{4, 30, 0.0, 'N'}, Long: coordinate{135, 54, 0, 'E'}},
	}

	for _, l := range locations {
		fmt.Printf("%15s %35s %15.2f %15.2f\n", l.Module, l.Place, l.Lat.decimal(), l.Long.decimal())
	}

	mars := world{radius: 3389.5}

	distances := make(map[float64]way)
	for _, start := range locations {
		for _, end := range locations {
			if start != end {
				distance := mars.distance(start, end)
				distances[distance] = way{start: start, end: end}
			}
		}
	}

	sorted := make([]float64, 0, len(distances))
	for d, _ := range distances {
		sorted = append(sorted, d)
	}
	sort.Float64s(sorted)

	way := distances[sorted[0]]
	fmt.Printf("min distance is %.2f between %s and %s\n", sorted[0], way.start.Module, way.end.Module)
	way = distances[sorted[len(distances)-1]]
	fmt.Printf("max distance is %.2f between %s and %s\n", sorted[len(distances)-1], way.start.Module, way.end.Module)

	earth := world{6371.0}
	London := location{Module: "London", Place: "England", Lat: coordinate{51, 30, 0, 'N'}, Long: coordinate{0, 8, 0, 'W'}}
	Paris := location{Module: "Paris", Place: "France", Lat: coordinate{48, 51, 0, 'N'}, Long: coordinate{2, 21, 0, 'E'}}

	fmt.Printf("distance between %s and %s is %.2f\n", London.Module, Paris.Module, earth.distance(London, Paris))

	MountSharp := location{Module: "Mount Sharp", Place: "Mars", Lat: coordinate{5, 4, 48, 'S'}, Long: coordinate{137, 51, 0, 'E'}}
	OlympusMons := location{Module: "Olympus Mons", Place: "Mars", Lat: coordinate{18, 39, 0, 'N'}, Long: coordinate{226, 12, 0, 'E'}}
	fmt.Printf("distance between %s and %s is %.2f\n", MountSharp.Module, OlympusMons.Module, mars.distance(MountSharp, OlympusMons))
}
