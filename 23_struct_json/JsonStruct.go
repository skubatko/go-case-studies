package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type location struct {
		Lat, Long float64 // Поля должны начинаться с большой буквы
	}

	curiosity := location{-4.5895, 137.4417}

	bytes, err := json.Marshal(curiosity)
	exitOnError(err)

	// ключи JSON совпадают с названиями полей структуры location
	fmt.Println(string(bytes)) // Выводит: {“Lat”:-4.5895,“Long”:137.4417}

	// использование тегов struct
	type locationTeg struct {
		// Теги в struct меняют вывод
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}

	curiosityTeg := locationTeg{-4.5895, 137.4417}

	bytes, err = json.Marshal(curiosityTeg)
	exitOnError(err)

	fmt.Println(string(bytes)) // Выводит: {“latitude”:-4.5895,“longitude”:137.4417}
}

// exitOnError выводит любые ошибки и выходит.
func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
