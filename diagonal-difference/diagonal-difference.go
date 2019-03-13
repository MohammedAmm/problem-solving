package main

import (
	"fmt"
)

// Complete the diagonalDifference function below.
func diagonalDifference(arr [][]int32) int32 {
	var result int32
	var diganonalOneSum int32
	var diganonalTwoSum int32
	for index := 0; index < len(arr); index++ {
		arrTwoIndex := len(arr) - index - 1
		diganonalOneSum += arr[index][index]
		diganonalTwoSum += arr[index][arrTwoIndex]
	}
	result = diganonalOneSum - diganonalTwoSum
	//check if result negative number to get the absolute value
	if result < 0 {
		// Change sign of negative number to postive
		return (result * -1)
	}
	return (result)
}

func main() {
	arr := [][]int32{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
	fmt.Println(diagonalDifference(arr))
}
