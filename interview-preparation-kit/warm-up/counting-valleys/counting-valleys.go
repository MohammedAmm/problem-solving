package main

import "fmt"

// Complete the countingValleys function below.
func countingValleys(n int32, s string) int32 {
	var numberOfValleys int32
	var currentMove string
	result := 0
	for _, element := range s {
		currentMove = string(element)
		oldResult := result
		if currentMove == "U" {
			result++
		} else {
			result--
		}
		if oldResult == 0 && result < 0 {
			numberOfValleys++
		}

	}
	return numberOfValleys
}

func main() {
	fmt.Println(countingValleys(8, "UDDDUDUU"))
}
