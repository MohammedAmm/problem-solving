package main

import "fmt"

//Swap element
func swap(x, y *int32) {
	*x, *y = *y, *x
}

//Insertion Sort
func insertionSort(arr []int32) []int32 {
	for unSortedIndex := 1; unSortedIndex < len(arr); unSortedIndex++ {
		sortedIndex := unSortedIndex - 1
		for arr[sortedIndex] > arr[unSortedIndex] {
			swap(&arr[sortedIndex], &arr[unSortedIndex])
			unSortedIndex--
			sortedIndex--
			if sortedIndex < 0 {
				break
			}
		}
	}
	return (arr)
}

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {

	//Sorting the array
	insertionSort(arr)

	//Declare result arr[minSum, maxSum]
	var result [2]int64

	//Sum only from second element to before the lastone and store to min
	for index := 1; index < len(arr)-1; index++ {
		result[0] += int64(arr[index])
	}

	//Assign min to max
	result[1] = result[0]

	//Add first element to min
	result[0] += int64(arr[0])

	//Add last element to max
	result[1] += int64(arr[len(arr)-1])
	fmt.Print(result[0], result[1])
}

func main() {
	arr := []int32{7, 8974, 69, 2, 221}
	miniMaxSum(arr)
}
