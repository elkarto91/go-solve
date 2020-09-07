package main

import (
	"fmt"
)

func main() {

	matrix := [][]int{{1, 1, 9}, {0, 2, 7}, {4, 5, 7}}
	fmt.Println("length ", len(matrix))
	fmt.Println("matrix ", matrix)

	x, y := zeroMat(matrix)
	fmt.Println("Length now ", len(x))
	if y {
		fmt.Println("Rotated matrix is", x)
	}

}

func zeroMat(matrix [][]int) ([][]int, bool) {

	n := len(matrix)
	m := len(matrix[0])
	fmt.Println("Length of matrix is", n)

	if n == 0 || n != len(matrix[0]) {
		return matrix, false
	}

	//find zeroes
	zeroRow := make([]int, n)
	zeroCol := make([]int, m)

	fmt.Println("Row array now is ", zeroRow)
	fmt.Println("Col array now is ", zeroCol)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				zeroRow[i] = -1
				zeroCol[j] = -1
			}
		}
	}

	fmt.Println("Row array now is ", zeroRow)
	fmt.Println("Col array now is ", zeroCol)

	for i := 0; i < len(zeroRow); i++ {
		if zeroRow[i] == -1 {
			matrix = nullifyRow(matrix, i)
		}
	}
	for i := 0; i < len(zeroCol); i++ {
		if zeroCol[i] == -1 {
			matrix = nullifyCol(matrix, i)
		}
	}
	fmt.Println("Matrix now is ", matrix)
	return matrix, true
}

func nullifyRow(matrix [][]int, i int) [][]int {

	for k := i; k < len(matrix); k++ {
		matrix[i][k] = 0
	}
	return matrix
}
func nullifyCol(matrix [][]int, i int) [][]int {

	for k := i; k < len(matrix[0]); k++ {
		matrix[k][i] = 0
	}
	return matrix
}
