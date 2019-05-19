package main

import (
	"fmt"
)

// func executeQuery(arrayOfValues []int64, query []int32) []int64 {
// 	for index := query[0]; index <= query[1]; index++ {
// 		arrayOfValues[index-1] += int64(query[2])
// 	}
// 	return arrayOfValues
// }

// // Complete the arrayManipulation function below.
// func arrayManipulation(n int32, queries [][]int32) int64 {
// 	var result int64
// 	intN := int64(n)
// 	//Make an array and initialize it with zeros
// 	arrayOfValues := make([]int64, intN)
// 	for _, query := range queries {
// 		executeQuery(arrayOfValues, query)
// 	}
// 	//Sort array without initialize new array
// 	insertionSort(arrayOfValues)
// 	//Max will be always the last element
// 	result = arrayOfValues[intN-1]

// 	return result
// }

// //Swap element
// func swap(x, y *int64) {
// 	*x, *y = *y, *x
// }

// //Insertion Sort
// func insertionSort(arr []int64) []int64 {
// 	for unSortedIndex := 1; unSortedIndex < len(arr); unSortedIndex++ {
// 		sortedIndex := unSortedIndex - 1
// 		for arr[sortedIndex] > arr[unSortedIndex] {
// 			swap(&arr[sortedIndex], &arr[unSortedIndex])
// 			unSortedIndex--
// 			sortedIndex--
// 			if sortedIndex < 0 {
// 				break
// 			}
// 		}
// 	}
// 	return (arr)
// }

// //Quicksort ...
// func Quicksort(v []int64) {
// 	if len(v) < 20 {
// 		insertionSort(v)
// 		return
// 	}
// 	p := Pivot(v)
// 	low, high := Partition(v, p)
// 	Quicksort(v[:low])
// 	Quicksort(v[high:])
// }

// //Pivot ...
// func Pivot(v []int64) int64 {
// 	n := len(v)
// 	return Median(v[rand.Intn(n)],
// 		v[rand.Intn(n)],
// 		v[rand.Intn(n)])
// }

// //Median ...
// func Median(a, b, c int64) int64 {
// 	if a < b {
// 		switch {
// 		case b < c:
// 			return b
// 		case a < c:
// 			return c
// 		default:
// 			return a
// 		}
// 	}
// 	switch {
// 	case a < c:
// 		return a
// 	case b < c:
// 		return c
// 	default:
// 		return b
// 	}
// }

// //Partition ...
// func Partition(v []int64, p int64) (low, high int) {
// 	low, high = 0, len(v)
// 	for mid := 0; mid < high; {
// 		// Invariant:
// 		//  - v[:low] < p
// 		//  - v[low:mid] = p
// 		//  - v[mid:high] are unknown
// 		//  - v[high:] > p
// 		//
// 		//         < p       = p        unknown       > p
// 		//     -----------------------------------------------
// 		// v: |          |          |a            |           |
// 		//     -----------------------------------------------
// 		//                ^          ^             ^
// 		//               low        mid           high
// 		switch a := v[mid]; {
// 		case a < p:
// 			v[mid] = v[low]
// 			v[low] = a
// 			low++
// 			mid++
// 		case a == p:
// 			mid++
// 		default: // a > p
// 			v[mid] = v[high-1]
// 			v[high-1] = a
// 			high--
// 		}
// 	}
// 	return
// }

func executeQuery(arrayOfValues []int64, query []int32) []int64 {
	arrayOfValues[query[0]-1] += int64(query[2])
	if int64(query[1]) < int64(len(arrayOfValues)) {
		arrayOfValues[query[1]] -= int64(query[2])
	}
	return arrayOfValues
}

// Complete the arrayManipulation function below.
func arrayManipulation(n int32, queries [][]int32) int64 {
	var max, currentMax int64
	intN := int64(n)
	//Make an array and initialize it with zeros
	arrayOfValues := make([]int64, intN)
	for _, query := range queries {
		executeQuery(arrayOfValues, query)
	}
	for _, value := range arrayOfValues {
		currentMax += value
		if currentMax > max {
			max = currentMax
		}
	}

	return max
}
func main() {
	n := 5
	queries := [][]int32{{1, 2, 5}, {2, 4, 6}, {3, 5, 4}}
	fmt.Println(arrayManipulation(int32(n), queries))

}
