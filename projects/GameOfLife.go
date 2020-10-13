package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func NewUniverse() Universe {
	universe := make([][]bool, height, height)
	for i := 0; i < height; i++ {
		universe[i] = make([]bool, width, width)
	}
	return universe
}

func (u Universe) Seed() {
	numberOfAlive := width * height / 4
	row := rand.Intn(height)
	cell := rand.Intn(width)
	for i := 0; i < numberOfAlive; i++ {
		for u[row][cell] {
			row = rand.Intn(height)
			cell = rand.Intn(width)
		}
		u[row][cell] = true
	}
}

func (u Universe) Show() {
	for _, row := range u {
		for _, isAlive := range row {
			if isAlive {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}

func (u Universe) Alive(x, y int) bool {
	return u[(y+height)%height][(x+width)%width]
}

func (u Universe) Neighbors(x, y int) int {
	var count int
	for row := -1; row <= 1; row++ {
		for cell := -1; cell <= 1; cell++ {
			if u.Alive(x+cell, y+row) && cell != x && row != y {
				count++
			}
		}
	}
	return count
}

func (u Universe) Next(x, y int) bool {
	neighbors := u.Neighbors(x, y)
	if u.Alive(x, y) {
		if neighbors < 2 || neighbors > 3 {
			return false
		}
	} else {
		if neighbors != 3 {
			return false
		}
	}
	return true
}

func Step(a, b Universe) {
	for row := 0; row < height; row++ {
		for cell := 0; cell < width; cell++ {
			b[row][cell] = a.Next(cell, row)
		}
	}
}

func main() {
	a := NewUniverse()
	b := NewUniverse()
	a.Seed()
	for true {
		//fmt.Print("\x0c")   // очищение экрана в Go Playground
		fmt.Print("\033[H") // очищение экрана в MacOS
		a.Show()
		Step(a, b)
		a, b = b, a
		time.Sleep(time.Second)
	}
}
