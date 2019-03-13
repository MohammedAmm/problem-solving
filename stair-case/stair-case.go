package main

import "fmt"

func staircase(n int32) {
	for index := 0; index < int(n); index++ {
		for index2 := 0; index2 < int(n)-(index+1); index2++ {
			fmt.Print(" ")
		}
		for index3 := 0; index3 < index+1; index3++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

}

func main() {
	var size int32 = 8
	staircase(size)
}
