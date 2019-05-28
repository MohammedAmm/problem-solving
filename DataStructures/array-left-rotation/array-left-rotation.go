package main

import "fmt"

func shiftLeftArray(arr []int32, shiftNumber int) []int32 {
	index := 0
	for index < shiftNumber {
		temp := arr[0]
		for arrIndex := 1; arrIndex < len(arr); arrIndex++ {
			arr[arrIndex-1] = arr[arrIndex]
		}
		arr[len(arr)-1] = temp
		index++
	}
	return arr
}
func main() {
	arr := []int32{1, 2, 3, 4, 5}
	d := 4
	for _, element := range shiftLeftArray(arr, d) {
		fmt.Print(element, " ")
	}

}
