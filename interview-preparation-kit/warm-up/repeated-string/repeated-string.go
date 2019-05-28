package main

import "fmt"

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
	//Hold the number of appearance
	var repeatCounter int64
	//Count of a in a give word
	var countOfA int
	//Hold count of a to specific index for the last time of repeation
	var countOfATill int64
	//Hold number of repeation of specific string through a given number
	var repeatedString float64
	//Loop through string to count appearance of a
	for index, element := range s {
		currentString := string(element)
		if currentString == "a" {
			countOfA++
		}
		if currentString == "a" && int64(index) < n%int64(len(s)) {
			countOfATill++
		}

	}
	//Calculate number of repeation of specific string through a given number
	repeatedString = float64(n) / float64(len(s))
	//Calculate number of a of all repeations
	repeatCounter = int64(repeatedString) * int64(countOfA)
	repeatCounter += countOfATill
	return repeatCounter
}

func main() {
	fmt.Println(repeatedString("a", 100000000000))
}
