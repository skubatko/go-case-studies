package main

import (
	"fmt"
	"sort"
	"strings"
)

func hyperspace(worlds []string) { // Данный аргумент является срезом, а не массивом
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

func main() {
	planets := []string{" Венера   ", "Земля  ", " Марс"} // Планеты, разделенные друг от друга пробелами
	hyperspace(planets)

	fmt.Println(strings.Join(planets, "")) // Выводит: ВенераЗемляМарс

	planets = []string{
		"Меркурий", "Венера", "Земля", "Марс",
		"Юпитер", "Сатурн", "Уран", "Нептун",
	}

	//sort.StringSlice(planets).Sort() // Сортирует planets в алфавитном порядке = sort.Strings(planets)
	sort.Strings(planets)

	sort.Strings(planets)
	fmt.Println(planets) // Выводит: [Венера Земля Марс Меркурий Нептун Сатурн Уран Юпитер]
}
