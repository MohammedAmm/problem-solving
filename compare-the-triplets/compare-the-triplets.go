package main

import (
	"fmt"
)

// Complete the compareTriplets function below.
func compareTriplets(a []int32, b []int32) []int32 {
	result := []int32{0, 0}
	for index, element := range a {
		if element > b[index] {
			result[0] += 1
		} else if element < b[index] {
			result[1] += 1
		}
	}
	return (result)
}

func main() {
	a := []int32{2, 2, 3}
	b := []int32{11, 2, 3}
	fmt.Println(compareTriplets(a, b))
}
