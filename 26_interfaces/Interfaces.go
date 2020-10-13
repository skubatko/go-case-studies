package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type talker interface {
	talk() string
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type starship struct {
	laser
}

type rover struct{}

func (r rover) talk() string {
	return "whir whir"
}

type stardater interface {
	YearDay() int
	Hour() int
}

// stardate возвращает выдуманное измерение времени.
func stardate(t stardater) float64 {
	doy := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + doy + h
}

type sol int

func (s sol) YearDay() int {
	return int(s % 668) // Марсианской год состоит из 668 дней
}

func (s sol) Hour() int {
	return 0 // Неизвестный час
}

// location с широтой и долготой в десятичных градусах
type location struct {
	lat, long float64
}

// String форматирует location с широтой и долготой
func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

type coordinate struct {
	d, m, s float64
	h       rune
}

// String форматирует координаты DMS
func (c coordinate) String() string {
	return fmt.Sprintf("%v°%v'%.1f\" %c", c.d, c.m, c.s, c.h)
}

// location с широтой и долготой в десятичных градусах
type locationPlus struct {
	lat, long coordinate
}

// String форматирует location с широтой и долготой
func (l locationPlus) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

// decimal конвертирует координаты d/m/s в десятичные градусы
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func (c coordinate) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		DD  float64 `json:"decimal"`
		DMS string  `json:"dms"`
		D   float64 `json:"degrees"`
		M   float64 `json:"minutes"`
		S   float64 `json:"seconds"`
		H   string  `json:"hemisphere"`
	}{
		DD:  c.decimal(),
		DMS: c.String(),
		D:   c.d,
		M:   c.m,
		S:   c.s,
		H:   string(c.h),
	})
}

// location с широтой и долготой в десятичных градусах
type locationExtended struct {
	Name string     `json:"name"`
	Lat  coordinate `json:"latitude"`
	Long coordinate `json:"longitude"`
}

func main() {
	// В отличие от Java, в Go martian и laser напрямую не объявляют, что они имплементируют интерфейс
	// Существуют правило в отношении именования типов интерфейса с суффиком -er

	var t talker

	t = martian{}
	fmt.Println(t.talk()) // Выводит: nack nack

	t = laser(3)
	fmt.Println(t.talk()) // Выводит: pew pew pew

	shout(t)
	shout(martian{}) // Выводит: NACK NACK
	shout(laser(2))  // Выводит: PEW PEW

	s := starship{laser(3)}

	fmt.Println(s.talk()) // Выводит: pew pew pew
	shout(s)              // Выводит: PEW PEW PEW

	shout(rover{})

	day := time.Date(2012, 8, 6, 5, 17, 0, 0, time.UTC)
	fmt.Printf("%.1f Curiosity has landed\n", stardate(day)) // Выводит: 1219.2 Curiosity has landed

	sol := sol(1422)
	fmt.Printf("%.1f Happy birthday\n", stardate(sol)) // Выводит: 1086.0 Happy birthday

	// Если тип предоставляет метод String то Println, Sprintf и другие функции отображения содержимого будут используют его
	curiosity := location{-4.5895, 137.4417}
	fmt.Println(curiosity) // Выводит: -4.5895, 137.4417

	elysium := locationPlus{
		lat:  coordinate{4, 30, 0.0, 'N'},
		long: coordinate{135, 54, 0.0, 'E'},
	}
	fmt.Println("Elysium Planitia is at", elysium) // Выводит: Elysium Planitia is at 4°30’0.0” N, 135°54’0.0” E

	elysiumExt := locationExtended{
		Name: "Elysium Planitia",
		Lat:  coordinate{4, 30, 0.0, 'N'},
		Long: coordinate{135, 54, 0.0, 'E'},
	}

	bytes, err := json.MarshalIndent(elysiumExt, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(bytes))
}
