package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func birthday(p *person) {
	p.age++
}

func (p *person) birthday() {
	p.age++
}

type stats struct {
	level             int
	endurance, health int
}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

type character struct {
	name  string
	stats stats
}

func reset(board *[8][8]rune) {
	board[0][0] = 'r'
	// ...
}

// Лишний указатель
func demolish(planets *map[string]string) {}

func reclassify(planets *[]string) {
	*planets = (*planets)[0:8]
}

func main() {
	answer := 42
	fmt.Println(&answer)

	address := &answer
	fmt.Printf("address это %T\n", address)
	fmt.Println(*address)

	canada := "Canada"

	var home *string
	fmt.Printf("home is a %T\n", home) // Выводит:  home is a *string

	home = &canada
	fmt.Println(*home) // Выводит: Canada
	var administrator *string

	scolese := "Christopher J. Scolese"
	administrator = &scolese
	fmt.Println(*administrator) // Выводит: Christopher J. Scolese

	bolden := "Charles F. Bolden"
	administrator = &bolden
	fmt.Println(*administrator) // Выводит: Charles F. Bolden

	bolden = "Charles Frank Bolden Jr."
	fmt.Println(*administrator) // Выводит: Charles Frank Bolden Jr.

	*administrator = "Maj. Gen. Charles Frank Bolden Jr."
	fmt.Println(bolden) // Выводит: Maj. Gen. Charles Frank Bolden Jr.

	major := administrator
	*major = "Major General Charles Frank Bolden Jr."
	fmt.Println(bolden) // Выводит: Major General Charles Frank Bolden Jr.

	fmt.Println(administrator == major) // Выводит: true

	charles := *major
	*major = "Charles Bolden"
	fmt.Println(charles) // Выводит: Major General Charles Frank Bolden Jr.
	fmt.Println(bolden)  // Выводит: Charles Bolden

	charles = "Charles Bolden"
	fmt.Println(charles == bolden)   // Выводит: true
	fmt.Println(&charles == &bolden) // Выводит: false

	// Указатели и структуры Go
	timmy := &person{
		name: "Timothy",
		age:  10,
	}

	timmy.superpower = "flying"
	fmt.Printf("%+v\n", timmy) // Выводит: &{name:Timothy superpower:flying age:10}

	// Указатели и массивы в Go
	superpowers := &[3]string{"flight", "invisibility", "super strength"}
	fmt.Println(superpowers[0])   // Выводит: flight
	fmt.Println(superpowers[1:2]) // Выводит: [invisibility]

	// Указатели в качестве параметров
	rebecca := person{
		name:       "Rebecca",
		superpower: "imagination",
		age:        14,
	}

	birthday(&rebecca)

	fmt.Printf("%+v\n", rebecca) // Выводит: {name:Rebecca superpower:imagination age:15}

	// Приемники указателя
	terry := &person{
		name: "Terry",
		age:  15,
	}
	terry.birthday()
	fmt.Printf("%+v\n", terry) // Выводит: &{name:Terry age:16}

	nathan := person{
		name: "Nathan",
		age:  17,
	}
	nathan.birthday()
	fmt.Printf("%+v\n", nathan) // Выводит: {name:Nathan age:18}

	// Если некоторым методам нужны приемники указателей, используйте из для всех методов типа

	// Внутренние указатели
	player := character{name: "Matthias"}
	levelUp(&player.stats)

	fmt.Printf("%+v\n", player.stats) // Выводит: {level:1 endurance:56 health:280}

	// мутации массивов
	var board [8][8]rune
	reset(&board)

	fmt.Printf("%c", board[0][0]) // Выводит: r

	// Карты в роли указателей

	// Срезы
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
		"Pluto",
	}
	reclassify(&planets)

	fmt.Println(planets) // Выводит: [Mercury Venus Earth Mars Jupiter Saturn Uranus Neptune]
}
