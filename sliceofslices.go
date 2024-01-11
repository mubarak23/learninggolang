// slice of slices
// Slices can hold other slices, or a 2D slice.
// rows := [][]int{}
// append on the same slice pretty much everytime
//
// mySlice := []int{1, 2, 3}
// mySlice = append(mySlice, 4)

package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, 0)
	for i :=0; i < rows; i++ {
		row := make([]int, 0)

		for j :=0; j < cols; j++ {
				row = append(row, i *j)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

// dont edit below this line

func test(rows, cols int) {
	fmt.Printf("Creating %v x %v matrix...\n", rows, cols)
	matrix := createMatrix(rows, cols)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test(3, 3)
	test(5, 5)
	test(10, 10)
	test(15, 15)
}
