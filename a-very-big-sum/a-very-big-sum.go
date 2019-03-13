package main

import (
	"fmt"
)

// Complete the aVeryBigSum function below.
func aVeryBigSum(ar []int64) int64 {
	var result int64
	for _, element := range ar {
		result += element
	}
	return (result)
}

func main() {
	a := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	fmt.Println(aVeryBigSum(a))
}
