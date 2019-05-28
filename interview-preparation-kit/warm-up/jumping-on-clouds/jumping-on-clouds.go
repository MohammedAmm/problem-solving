package main

import "fmt"

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32) int32 {
	//Hold number of steps
	var stepCount int32
	//Loop through array to check for available cumulus
	for index := 0; index < len(c)-1; index++ {
		//Check if a double jump can be done.
		if index+2 < len(c) && c[index+2] == 0 {
			index++
		}
		stepCount++
	}
	return int32(stepCount)
}

func main() {
	fmt.Println(jumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 0}))
}
