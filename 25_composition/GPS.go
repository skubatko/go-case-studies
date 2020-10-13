package main

import (
	"fmt"
	"math"
)

type location struct {
	name      string
	lat, long float64
}

type world struct {
	radius float64
}

type gps struct {
	current     location
	destination location
	world
}

func (l location) description() string {
	return fmt.Sprint(l.name, " ", l.lat, " ", l.long)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong) // Использует поле радиуса world
}

func (g gps) distance() float64 {
	return g.world.distance(g.current, g.destination)
}

func (g gps) message() string {
	return fmt.Sprintf("%.2f", g.distance())
}

type rover struct {
	gps
}

func main() {
	mars := world{radius: 3389.5}
	bradbury := location{name: "Bradbury Landing", lat: -4.5895, long: 137.4417}
	elysium := location{name: "Elysium Planitia", lat: 4.5, long: 135.9}
	gps := gps{
		world:       mars,
		current:     bradbury,
		destination: elysium,
	}
	curiosity := rover{gps}
	fmt.Println(curiosity.message())
}
