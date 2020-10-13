package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64
type sensor func() kelvin

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 0 // TO-DO: внедрить реальный сенсор
}

func measureTemperature(samples int, sensor func() kelvin) { // measureTemperature принимает функцию в качестве второго параметра
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%v° K\n", k)
		time.Sleep(time.Second)
	}
}

func measureTemperatureTyped(samples int, s sensor) {
	measureTemperature(samples, s)
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin { // Объявляет и возвращает анонимную функцию
		return s() + offset
	}
}

func main() {
	var sensor func() kelvin

	sensor = fakeSensor // Присваивает функцию переменной

	fmt.Println(sensor())

	sensor = realSensor
	fmt.Println(sensor())

	measureTemperature(3, fakeSensor) // Передает название функции другой функции
	measureTemperature(3, realSensor)

	var f = func() { // Присваивает анонимную функцию переменной
		fmt.Println("Dress up for the masquerade.")
	}

	f() // Выводит: Dress up for the masquerade.

	fVar := func(message string) { // Присваивает анонимную функцию переменной
		fmt.Println(message)
	}
	fVar("Go to the party.") // Выводит: Go to the party.

	func() { // Объявляет анонимную функцию
		fmt.Println("Functions anonymous")
	}() // Вызов функции

	sensor = calibrate(realSensor, 5)
	fmt.Println(sensor()) // Выводит: 5

	//замыкание сохраняет ссылку на ближайшие переменные
	var k kelvin = 294.0

	sensor = func() kelvin {
		return k
	}
	fmt.Println(sensor()) // Выводит: 294

	k++
	fmt.Println(sensor()) // Выводит: 295
}
