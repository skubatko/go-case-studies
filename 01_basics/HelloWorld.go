package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Println("Hello, Sergey")
	fmt.Println("Hello, awesome")

	fmt.Printf("%v\n", 56_000_000+rand.Intn(401_000_000-56_000_000))
	fmt.Printf("%v\n", rand.Intn(12))

	fmt.Println("яблоко" > "банан")

	var year = 2100
	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)

	if leap {
		fmt.Println("Этот год високосный!")
	} else {
		fmt.Println("К сожалению, нет. Этот год не високосный.")
	}

	fmt.Println("Здесь вход в пещеру и путь на восток.")
	var command = "зайти внутрь"

	switch command { // Сравнивает case с command
	case "идти на восток":
		fmt.Println("Вы направляетесь к горе.")
	case "зайти в пещеру", "зайти внутрь": // Запятая разделяет список возможных значений
		fmt.Println("Вы находитесь в тускло освещенной пещере.")
	case "прочитать знак":
		fmt.Println("На знаке написано 'Несовершеннолетним вход запрещен'.")
	default:
		fmt.Println("Пока не совсем понятно.")
	}

	const col = 30
	// Clear the screen by printing \x0c.
	//bar := fmt.Sprintf("\x0c[%%-%vs]", col)
	bar := fmt.Sprintf("\033[H[%%-%vs]", col)
	for i := 0; i < col; i++ {
		fmt.Printf(bar, strings.Repeat("=", i)+">")
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf(bar+" Done!", strings.Repeat("=", col))
}
