package main

import (
	"fmt"
)

func validareSudoku(sudoku [9][9]int, a[3][9]int) bool {

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			b := sudoku[i][j] - 1
			t := 1 << b

			if a[0][i]&t != 0 {
				errorType(sudoku, 0, i)
				return false
			}
			a[0][i] = a[0][i] | t

			if a[1][j]&t != 0 {
				errorType(sudoku, 1, j)
				return false
			}
			a[1][j] = a[1][j] | t

			k := i/3*3 + j/3
			if a[2][k]&t != 0 {
				errorType(sudoku, 2, k)
				return false
			}
			a[2][k] = a[2][k] | t
		}
	}

	return true
}


func errorType(sudoku [9][9]int, errorType int, position int) {

	fmt.Println()

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			if (errorType == 0 && i == position) || (errorType == 1 && j == position) || (errorType == 2 && i/3*3+j/3 == position) {
				fmt.Print( sudoku[i][j])
			} else {
				fmt.Print( sudoku[i][j])
			}

			if (j+1)%3 == 0 && j+1 != 9 {
				fmt.Print( "  ")
			} else {
				fmt.Print(" ")
			}

		}
		if (i+1)%3 == 0 && i+1 != 9 {
			fmt.Print( "  \n\n")
		} else {
			fmt.Print("\n")
		}
	}

	fmt.Println()

}

func main() {

	//configuratie valida
	var sudokuvalid = [9][9]int{
		{9, 3, 4, 5, 6, 8, 1, 2, 7},
		{8, 2, 6, 7, 1, 4, 5, 9, 3},
		{1, 5, 7, 9, 2, 3, 4, 6, 8},
		{2, 7, 8, 1, 5, 9, 3, 4, 6},
		{6, 4, 1, 3, 8, 7, 2, 5, 9},
		{3, 9, 5, 6, 4, 2, 7, 8, 1},
		{5, 6, 3, 4, 9, 1, 8, 7, 2},
		{7, 8, 9, 2, 3, 5, 6, 1, 4},
		{4, 1, 2, 8, 7, 6, 9, 3, 5}}

	//// greseala la linia 7
	//sudokuvalid = [9][9]int{
	//	{9, 3, 4, 5, 6, 8, 1, 2, 7},
	//	{8, 2, 6, 7, 1, 4, 5, 9, 3},
	//	{1, 5, 7, 9, 2, 3, 4, 6, 8},
	//	{2, 7, 8, 1, 5, 9, 3, 4, 6},
	//	{6, 4, 1, 3, 8, 7, 2, 5, 9},
	//	{3, 9, 5, 6, 4, 2, 7, 8, 1},
	//	{5, 6, 3, 4, 9, 1, 5, 7, 2},
	//	{7, 8, 9, 2, 3, 5, 6, 1, 4},
	//	{4, 1, 2, 8, 7, 6, 9, 3, 5}}



	var a = [3][9]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			fmt.Print( sudokuvalid[i][j])

			if (j+1)%3 == 0 && j+1 != 9 {
				fmt.Print("  ")
			}else{
				fmt.Print(" ")
			}
		}
			if (i+1)%3 == 0 && i+1 != 9 {
				fmt.Print("  \n\n")
			} else {
				fmt.Print("\n")
			}
	}
			fmt.Print()

	if validareSudoku(sudokuvalid, a) {
		fmt.Println("Sudoku valid!")
	} else {
		fmt.Println( "Sudoku invalida!")
	}

	fmt.Print()
}

