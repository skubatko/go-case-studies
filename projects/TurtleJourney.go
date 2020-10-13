package main

import (
	"fmt"
	"math/rand"
	"time"
)

type coordinate struct {
	x int
	y int
}

type turtle struct {
	name     string
	position coordinate
}

func (t *turtle) up() {
	t.position.y--
}

func (t *turtle) down() {
	t.position.y++
}

func (t *turtle) left() {
	t.position.x--
}

func (t *turtle) right() {
	t.position.x++
}

func (t turtle) String() string {
	return fmt.Sprintf("turtle %s located at %d,%d", t.name, t.position.x, t.position.y)
}

func main() {
	rand.Seed(time.Now().Unix())
	t := turtle{name: "Масяня", position: coordinate{x: 0, y: 0}}
	for i := 0; i < 100; i++ {
		switch rand.Intn(4) {
		case 0:
			t.up()
		case 1:
			t.down()
		case 2:
			t.left()
		case 3:
			t.right()
		default:
			// stay
		}
	}
	fmt.Println(t)
}
