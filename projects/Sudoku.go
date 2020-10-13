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
	ErrRow    = errors.New("invalid row")
	ErrCol    = errors.New("invalid column")
	ErrPart   = errors.New("invalid part")
)

// Grid является сеткой Судоку
type Grid [rows][columns]int8

func NewSudoku(array [9][9]int8) Grid {
	return array
}

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

func (g Grid) validInRow(row, col int, value int8) bool {
	for c := 0; c < 9; c++ {
		if c != col && g[row][c] == value {
			return false
		}
	}

	return true
}

func (g Grid) validInCol(row, col int, value int8) bool {
	for r := 0; r < 9; r++ {
		if r != row && g[r][col] == value {
			return false
		}
	}

	return true
}

func (g Grid) validInPart(row, col int, value int8) bool {
	rMin := row / 3 * 3
	cMin := col / 3 * 3
	for r := rMin; r < rMin+3; r++ {
		for c := cMin % 3; c < cMin+3; c++ {
			if r != row && c != col && g[r][c] == value {
				return false
			}
		}
	}

	return true
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

	if !g.validInRow(row, column, digit) {
		errs = append(errs, ErrRow)
	}

	if !g.validInCol(row, column, digit) {
		errs = append(errs, ErrCol)
	}

	if !g.validInPart(row, column, digit) {
		errs = append(errs, ErrPart)
	}

	if len(errs) > 0 {
		return errs
	}

	g[row][column] = digit
	return nil
}

func main() {
	s := NewSudoku([rows][columns]int8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})

	err := s.Set(5, 5, 6)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, row := range s {
		fmt.Println(row)
	}
}
