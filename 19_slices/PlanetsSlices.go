package main

import "fmt"

type Planets []string

func (planets Planets) terraform() {
	for i := range planets {
		planets[i] = "Новый " + planets[i]
	}
}

// dump для длины, вместимости и содержимого среза
func dump(label string, slice []string) {
	fmt.Printf("%v: длина %v, вместимость %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	planets := []string{
		"Меркурий", "Венера", "Земля", "Марс",
		"Юпитер", "Сатурн", "Уран", "Нептун",
	}
	Planets(planets[3:4]).terraform()
	Planets(planets[6:]).terraform()
	fmt.Println(planets)

	dwarfs := []string{"Церера", "Плутон", "Хаумеа", "Макемаке", "Эрида"}

	dwarfs = append(dwarfs, "Оркус")
	dwarfs = append(dwarfs, "Салация", "Квавар", "Седна")
	fmt.Println(dwarfs) // Выводит: [Церера Плутон Хаумеа Макемаке Эрида Оркус]
	fmt.Println(len(dwarfs))

	dwarfs = []string{"Церера", "Плутон", "Хаумеа", "Макемаке", "Эрида"}
	dump("dwarfs", dwarfs)           // Выводит: dwarfs: длина 5, вместимость 5 [Церера Плутон Хаумеа Макемаке Эрида]
	dump("dwarfs[1:2]", dwarfs[1:2]) // Выводит: dwarfs[1:2]: длина 1, вместимость 4 [Плутон]

	dwarfs1 := []string{"Церера", "Плутон", "Хаумеа", "Макемаке", "Эрида"} // Длина 5, вместимость 5
	dump("dwarfs1", dwarfs1)
	dwarfs2 := append(dwarfs1, "Оркус") // Длина 6, вместимость 10
	dump("dwarfs2", dwarfs2)
	dwarfs3 := append(dwarfs2, "Салация", "Квавар", "Седна") // Длина 9, вместимость 10
	dump("dwarfs3", dwarfs3)

	planets = []string{
		"Меркурий", "Венера", "Земля", "Марс",
		"Юпитер", "Сатурн", "Уран", "Нептун",
	}
	fmt.Println(planets) // Выводит: [Меркурий Венера Земля Марс Юпитер Сатурн Уран Нептун]
	dump("planets", planets)

	terrestrial := planets[0:4:4] // Длина 4, вместимость 4
	dump("terrestrial", terrestrial)

	worlds := append(terrestrial, "Церера")
	fmt.Println(worlds)
	dump("worlds", worlds)

	// Если третий индекс не уточняется, вместимость terrestrial равна 8.
	// Добавление элемента "Церера" не перемещает новый массив, а вместо этого переписывает "Юпитер"
	terrestrial = planets[0:4] // Длина 4, вместимость 8
	dump("terrestrial", terrestrial)

	worlds = append(terrestrial, "Церера")
	dump("worlds", worlds)

	fmt.Println(planets) // Выводит: [Меркурий Венера Земля Марс Церера Сатурн Уран Нептун]
	dump("planets", planets)

	dwarfs = make([]string, 0, 10)
	dump("dwarfs", dwarfs)
	dwarfs = append(dwarfs, "Церера", "Плутно", "Хаумеа", "Макемаке", "Эрида")
	dump("dwarfs", dwarfs)

	dwarfs = make([]string, 10)
	dump("dwarfs", dwarfs)
	dwarfs = append(dwarfs, "Церера", "Плутно", "Хаумеа", "Макемаке", "Эрида")
	dump("dwarfs", dwarfs)
}
