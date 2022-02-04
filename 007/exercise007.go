package main

import "fmt"

func createMatrix(x, y int) [][]int {
	matrix := make([][]int, x)

	for i := 0; i < x; i++ {
		row := make([]int, y)
		for j := 0; j < y; j++ {
			row[j] = i * j
		}
		matrix[i] = row
	}
	return matrix
}

func main() {

	ans := createMatrix(3, 5)
	fmt.Println(ans)

}
