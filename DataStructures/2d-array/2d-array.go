package main

import "fmt"

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {
	//Variable to store max
	var maxResult int32 = -63
	rows := len(arr)
	limit := rows / 2

	//Iterate for each row
	for i := 0; i <= limit; i++ {
		//Iterate for each column
		for j := 0; j <= limit; j++ {
			//For each element sum hourglassSum
			result := arr[i+1][j+1]
			for k := (0 + j); k <= (2 + j); k++ {
				result += arr[i][k] + arr[i+2][k]
			}
			//Check if first sum greater than max then assign it to max
			if result > maxResult {
				maxResult = result
			}
		}

	}
	return maxResult
}

func main() {
	arr := [][]int32{{-9, -9, -9, 1, 1, 1},
		{0, -9, 0, 4, 3, 2},
		{-9, -9, 11, 12, 12, 12},
		{8, 8, 0, 11, 12, 9},
		{0, 8, 11, 12, 12, 12},
		{8, 8, 8, 9, 9, 9}}
	fmt.Println(hourglassSum(arr))
}
