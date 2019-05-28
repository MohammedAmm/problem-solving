package main

import "fmt"

//Check if array contain value or not
//Return with index
func contains(array []int32, item int32) int {
	for index, currentItem := range array {
		if currentItem == item {
			return index
		}
	}
	return -1
}

// Complete the sockMerchant function below.
func sockMerchant(n int32, ar []int32) int32 {
	arr := []int32{0}
	result := 0
	for _, element := range ar {
		elementIndex := contains(arr, element)
		//If exists it means found new pair so delete it from slice
		//if doesn't exist push item to array to later check
		if elementIndex != -1 {
			result++
			arr[elementIndex] = 0
		} else {
			arr = append(arr, element)
		}
	}
	return int32(result)
}
func main() {
	fmt.Println(sockMerchant(7, []int32{1, 2, 1, 2, 1, 3, 3, 3, 3}))
}
