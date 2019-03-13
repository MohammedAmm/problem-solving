package main

import "fmt"

// Complete the birthdayCakeCandles function below.
func birthdayCakeCandles(ar []int32) int32 {
	var numberOfCandles int32
	var tallestCandle int32
	for _, element := range ar {
		if tallestCandle < element {
			tallestCandle = element
		}
	}

	for _, element := range ar {
		if tallestCandle == element {
			numberOfCandles++
		}
	}
	return numberOfCandles
}

func main() {
	arr := []int32{3, 1, 3, 3}
	fmt.Println(birthdayCakeCandles(arr))
}
