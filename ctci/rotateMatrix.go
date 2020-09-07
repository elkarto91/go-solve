package main

import "fmt"

func main() {

	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	fmt.Println("length ", len(matrix))
	fmt.Println("matrix ", matrix)

	x, y := rotate(matrix)
	fmt.Println("Length now ", len(x))
	if y {
		fmt.Println("Rotated matrix is", x)
	}

}

func rotate(matrix [][]int) ([][]int, bool) {

	n := len(matrix)
	fmt.Println("Length of matrix is", n)

	if n == 0 || n != len(matrix[0]) {
		return matrix, false
	}

	for layer := 0; layer < n/2; layer++ {

		fmt.Println("N/2", n/2)

		first := layer
		last := n - 1 - layer

		for i := first; i < last; i++ {
			offset := i - first
			top := matrix[first][i]

			fmt.Println("Layer ", layer, " First ", first, " Last ", last, " Offset ", offset)
			fmt.Println("Top -> ", top)

			fmt.Println(matrix[first][i], " <- ", matrix[last-offset][first])
			matrix[first][i] = matrix[last-offset][first]

			fmt.Println(matrix[last-offset][first], " <- ", matrix[last][last-offset])
			matrix[last-offset][first] = matrix[last][last-offset]

			fmt.Println(matrix[last][last-offset], " <- ", matrix[i][last])
			matrix[last][last-offset] = matrix[i][last]

			fmt.Println(matrix[i][last], " <- ", top)
			matrix[i][last] = top

			fmt.Println("Matrix at end of cycle -> ", matrix)

		}
	}
	return matrix, true
}
