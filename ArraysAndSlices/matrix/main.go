package main

import "fmt"

func main() {
	test(3, 3)
	test(5, 5)
	test(10, 10)
	test(15, 15)
}

func test(rows, cols int) {
	fmt.Printf("Creating matrix of %v * %v\n", rows, cols)
	matrix := createMatrix(rows, cols)
	fmt.Println("Displaying matirx: ")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf(" %v ", matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println("======END OF REPORT=======")
}

func createMatrix(rows, cols int) [][]int {
	matrix := [][]int{}
	for i := 0; i < rows; i++ {
		row := []int{}
		for j := 0; j < cols; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}
	return matrix
}
