package main

import (
	"fmt"
	"math/rand"
)

type animaler interface {
	String() string
	eat() string
	move() string
}

type tiger struct {
	name string
}

func (t tiger) String() string {
	return t.name
}

func (t tiger) eat() string {
	return "rrrrr ..."
}

func (t tiger) move() string {
	return "прыг-скок ..."
}

type falcon struct {
	name string
}

func (f falcon) String() string {
	return f.name
}

func (f falcon) eat() string {
	return "клю-клю ..."
}

func (f falcon) move() string {
	return "фьють ..."
}

func main() {
	animals := []animaler{
		tiger{name: "Киса Варя"},
		falcon{name: "Сокол Петя"},
	}

	for day := 1; day <= 3; day++ {
		for hour := 1; hour <= 24; hour++ {
			if hour > 7 && hour < 19 {
				animal := animals[rand.Intn(len(animals))]
				action := rand.Intn(2)
				if action == 0 {
					fmt.Printf("%v action: %s\n", animal, animal.eat())
				} else {
					fmt.Printf("%v action: %s\n", animal, animal.move())
				}
			}
		}
	}
}
