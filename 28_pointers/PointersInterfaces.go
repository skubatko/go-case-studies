package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l *laser) talk() string {
	return strings.Repeat("pew ", int(*l))
}

func main() {
	// Указатели и интерфейсы
	shout(martian{})  // Выводит: NACK NACK
	shout(&martian{}) // Выводит: NACK NACK

	pew := laser(2)
	shout(&pew) // Выводит: PEW PEW
}
