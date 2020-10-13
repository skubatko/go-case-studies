package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	temperature := map[string]int{
		"Земля": 15, // Пары ключ-значения являются композитными литералами для карт
		"Марс":  -65,
	}

	temp := temperature["Земля"]
	fmt.Printf("Средняя температура на поверхности Земли составляет %v° C.\n", temp) // Выводит: Средняя температура на поверхности Земли составляет 15° C.

	temperature["Земля"] = 16 // Небольшое изменение климата
	temperature["Венера"] = 464

	fmt.Println(temperature) // Выводит: map[Венера:464 Земля:16 Марс:-65]

	moon := temperature["Луна"]
	fmt.Println(moon) // Выводит: 0

	// Переменная moon должна содержать значение ключа "Луна" или же нулевое значение.
	// При наличии ключа значение дополнительной переменной ok будет равно true,  в противном случае — false.
	if moon, ok := temperature["Луна"]; ok { // Синтаксис comma, ok
		fmt.Printf("Средняя температура на поверхности Луны составляет %v° C.\n", moon)
	} else {
		fmt.Println("Где Луна?") // Выводит: Где Луна?
	}

	planets := map[string]string{
		"Земля": "Сектор ZZ9",
		"Марс":  "Сектор ZZ9",
	}

	planetsMarkII := planets
	planets["Земля"] = "упс"

	fmt.Println(planets)       // Выводит: map[Земля:упс Марс:Сектор ZZ9]
	fmt.Println(planetsMarkII) // Выводит: map[Земля:упс Марс:Сектор ZZ9]

	delete(planets, "Земля")   // Земля удаляется из карты
	fmt.Println(planetsMarkII) // Выводит: map[Марс:Сектор ZZ9]

	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	frequency := make(map[float64]int)

	for _, t := range temperatures { // Итерирует через срез (индекс, значение)
		frequency[t]++
	}

	for t, num := range frequency { // Итерирует через карту (ключ, значение)
		fmt.Printf("%+.2f встречается %d раз(а) \n", t, num)
	}

	temperatures = []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	groups := make(map[float64][]float64) // Карта с ключами float64 и значениями []float64

	for _, t := range temperatures {
		g := math.Trunc(t/10) * 10 // Округляет температуры вниз -20, -30 и так далее
		groups[g] = append(groups[g], t)
	}

	for group, temperatures := range groups {
		fmt.Printf("%v: %v\n", group, temperatures)
	}

	// Set
	temperatures = []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	set := make(map[float64]bool) // Создание карты с булевыми значениями
	for _, t := range temperatures {
		set[t] = true
	}

	if set[-28.0] {
		fmt.Println("часть множества") // Выводит: часть множества
	}

	fmt.Println(set) // Выводит: Prints map[-31:true -29:true -23:true -33:true -28:true 32:true]

	// Sort
	unique := make([]float64, 0, len(set))
	for t := range set {
		unique = append(unique, t)
	}
	sort.Float64s(unique)

	fmt.Println(unique) // Выводит: [-33 -31 -29 -28 -23 32]
}
