package main

import "fmt"

// Complete the formingMagicSquare function below.
func formingMagicSquare(s [][]int32) int32 {
	arrLen := len(s)
	initialzedArr := make([]int32, arrLen*arrLen)

	for _, element := range s {
		for _, innerElement := range element {
			// currentIndex := (index * arrLen) + innerIndex
			initialzedArr[innerElement-1]++
		}
	}
	fmt.Println(initialzedArr)
	return int32(1)
}

func main() {
	n := [][]int32{{1, 2, 4}, {1, 2, 3}, {1, 2, 3}}
	formingMagicSquare(n)
}
