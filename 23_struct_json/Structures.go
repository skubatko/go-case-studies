package main

import "fmt"

type location struct {
	lat, long float64
}

func main() {
	var curiosity struct {
		lat  float64
		long float64
	}

	// Присваивание значений полям структуры
	curiosity.lat = -4.5895
	curiosity.long = 137.4417

	fmt.Println(curiosity.lat, curiosity.long) // Выводит: -4.5895 137.4417
	fmt.Println(curiosity)                     // Выводит: {-4.5895 137.4417}

	var spirit location // Повторное использование типа location
	spirit.lat = -14.5684
	spirit.long = 175.472636

	var opportunity location // Повторное использование типа location
	opportunity.lat = -1.9462
	opportunity.long = 354.4734

	fmt.Println(spirit, opportunity) // Выводит: {-14.5684 175.472636} {-1.9462 354.4734}

	opportunity = location{lat: -1.9462, long: 354.4734}
	fmt.Println(opportunity) // Выводит: {-1.9462 354.4734}

	insight := location{lat: 4.5, long: 135.9}
	fmt.Println(insight) // Выводит: {4.5 135.9}

	spirit = location{-14.5684, 175.472636}
	fmt.Println(spirit) // Выводит: {-14.5684 175.472636}

	// символ %v со знаком плюса + для вывода названий полей
	curiosity = location{-4.5895, 137.4417}
	fmt.Printf("%v\n", curiosity)  // Выводит: {-4.5895 137.4417}
	fmt.Printf("%+v\n", curiosity) // Выводит: {lat:-4.5895 long:137.4417}

	//  Переменная curiosity инициализируется с копией значений, находящихся в bradbury
	bradbury := location{-4.5895, 137.4417}
	curiosity = bradbury

	curiosity.long += 0.0106 // Направляется на восток к бухте Йеллоунайф

	fmt.Println(bradbury, curiosity) // Выводит: {-4.5895 137.4417} {-4.5895 137.4523}

	// Пример среза структур
	type locationPlus struct {
		name string
		lat  float64
		long float64
	}

	locations := []locationPlus{
		{name: "Bradbury Landing", lat: -4.5895, long: 137.4417},
		{name: "Columbia Memorial Station", lat: -14.5684, long: 175.472636},
		{name: "Challenger Memorial Station", lat: -1.9462, long: 354.4734},
	}
	fmt.Println(locations)
}
