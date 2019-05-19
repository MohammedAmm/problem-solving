package main

import "fmt"

func gradingStudents(grades []int32) []int32 {
	/*
	 * Write your code here.
	 */
	//loop on array
	for index, element := range grades {
		//Get the remaining of division by 5 ex.13 % 5 = 3
		remain := (element % 5)
		//Check if element greater than forty
		//Check if remain is 3 or 4 to round the number
		if element > 37 && remain >= 3 {
			//First reduce remain from number to round it
			grades[index] = (element - remain) + 5
		}
	}

	return grades
}

func main() {
	fmt.Println(gradingStudents([]int32{38, 55, 60, 98}))
}
