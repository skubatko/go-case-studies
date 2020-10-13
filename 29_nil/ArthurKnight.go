package main

import "fmt"

type item struct {
	name string
}

func (i item) String() string {
	return fmt.Sprintf("item: %v", i.name)
}

type character struct {
	name     string
	leftHand *item
}

func (c character) String() string {
	return fmt.Sprintf("character: %v", c.name)
}

func (c *character) pickup(i *item) {
	fmt.Printf("%v pickups %v\n", c, i)
	c.leftHand = i
}

func (c *character) give(to *character) {
	fmt.Printf("%v gives %v to %v\n", c, c.leftHand, to)
	to.leftHand = c.leftHand
}

func main() {
	arthur := character{name: "Arthur"}
	knight := character{name: "Knight"}

	arthur.pickup(&item{name: "Axe"})
	arthur.give(&knight)
}
