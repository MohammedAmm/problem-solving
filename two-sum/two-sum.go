package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result [2]int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == target-nums[i+1] {
			result[0] = i
			result[1] = i + 1
			return result[:]
		}
	}
	return (nil)
}
func main() {
	var input = []int{2, 7, 11, 15}
	fmt.Println(twoSum(input, 9))
}
