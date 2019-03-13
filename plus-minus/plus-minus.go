package main

import "fmt"

func plusMinus(arr []int32) {
	var plusNumbers float64
	var minusNumbers float64
	var zeroNumbers float64
	arrSize := float64(len(arr))
	for _, element := range arr {
		if element > 0 {
			plusNumbers++
		} else if element < 0 {
			minusNumbers++
		} else {
			zeroNumbers++
		}
	}
	fmt.Printf("%.6f", plusNumbers/arrSize)
	fmt.Println()
	fmt.Printf("%.6f", minusNumbers/arrSize)
	fmt.Println()
	fmt.Printf("%.6f", zeroNumbers/arrSize)
	fmt.Println()
}

func main() {
	arr := []int32{-4, 3, -9, 0, 4, 1}
	plusMinus(arr)
}
