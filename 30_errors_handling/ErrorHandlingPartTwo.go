package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const rows, columns = 9, 9

var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit  = errors.New("invalid digit")
)

// Grid является сеткой Судоку
type Grid [rows][columns]int8

func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if column < 0 || column >= columns {
		return false
	}
	return true
}

func validDigit(digit int8) bool {
	return 0 < digit && digit < 10
}

type SudokuError []error

// Error возвращает одну или несколько ошибок через запятые
func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error()) // Конвертирует ошибки в строки
	}
	return strings.Join(s, ", ")
}

func (g *Grid) Set(row, column int, digit int8) error { // Возвращает тип ошибки
	var errs SudokuError
	if !inBounds(row, column) {
		errs = append(errs, ErrBounds)
	}
	if !validDigit(digit) {
		errs = append(errs, ErrDigit)
	}
	if len(errs) > 0 {
		return errs
	}

	g[row][column] = digit
	return nil // Возвращает nil
}

func main() {
	var g Grid
	err := g.Set(3, 0, 5)
	if err != nil {
		fmt.Printf("An error occurred: %v.\n", err)
		os.Exit(1)
	}

	err = g.Set(0, 0, 15)
	if err != nil {
		switch err {
		case ErrBounds, ErrDigit:
			fmt.Println("Les erreurs de paramètres hors limites.")
		default:
			fmt.Println(err)
		}
		os.Exit(1)
	}
}
