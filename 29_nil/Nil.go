package main

import (
	"fmt"
	"sort"
)

type person struct {
	age int
}

func (p *person) birthday() {
	if p == nil {
		return
	}
	p.age++ // разыменование указателя nil
}

func sortStrings(s []string, less func(i, j int) bool) {
	if less == nil {
		less = func(i, j int) bool { return s[i] < s[j] }
	}
	sort.Slice(s, less)
}

func mirepoix(ingredients []string) []string {
	return append(ingredients, "onion", "carrot", "celery")
}

type number struct {
	value int
	valid bool
}

func newNumber(v int) number {
	return number{value: v, valid: true}
}

func (n number) String() string {
	if !n.valid {
		return "not set"
	}
	return fmt.Sprintf("%d", n.value)
}

func main() {
	var nowhere *int
	fmt.Println(nowhere) // Выводит: <nil>
	// fmt.Println(*nowhere) // Выводит: nil pointer dereference
	if nowhere != nil {
		fmt.Println(*nowhere)
	}

	var nobody *person
	fmt.Println(nobody) // Выводит: <nil>

	nobody.birthday()

	var fn func(a, b int) int
	fmt.Println(fn == nil) // Выводит: true

	food := []string{"onion", "carrot", "celery"}
	sortStrings(food, nil)
	fmt.Println(food) // Выводит: [carrot celery onion]
	sortStrings(food, func(i, j int) bool { return len(food[i]) < len(food[j]) })
	fmt.Println(food) // Выводит: [carrot celery onion]

	// Срезы nil
	var soup []string
	fmt.Println(soup == nil) // Выводит: true

	for _, ingredient := range soup {
		fmt.Println(ingredient)
	}

	fmt.Println(len(soup)) // Выводит: 0

	soup = append(soup, "onion", "carrot", "celery")
	fmt.Println(soup) // Выводит: [onion carrot celery]

	soup = mirepoix(nil)
	fmt.Println(soup) // Выводит: [onion carrot celery]

	// Карты nil
	var soupMap map[string]int
	fmt.Println(soup == nil) // Выводит: true

	measurement, ok := soupMap["onion"]
	if ok {
		fmt.Println(measurement)
	}

	for ingredient, measurement := range soupMap {
		fmt.Println(ingredient, measurement)
	}

	// Интерфейсы nil
	var v interface{}
	fmt.Printf("%T %v %v\n", v, v, v == nil) // Выводит: <nil> <nil> true

	var p *int
	v = p
	fmt.Printf("%T %v %v\n", v, v, v == nil) // Выводит: *int <nil> false

	fmt.Printf("%#v\n", v) // Выводит: (*int)(nil)

	// Альтернатива nil
	n := newNumber(42)
	fmt.Println(n) // Выводит: 42

	e := number{}
	fmt.Println(e) // Выводит: not set
}
