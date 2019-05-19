//Equilibrium index of an array
//Input : A[] = {-7, 1, 5, 2, -4, 3, 0}
//Output : 3

package main

import "fmt"

func main() {
	//sumLeft will be the sum from [L1 .. R1] of PrefixSum array and sumRight will be the sum from [L2 .. R2] of PrefixSum array
	//
	// var sumLeft, sumRight int
	//Steps:
	//Step 1: Calculate the prefixsum array
	//Step 2: loop on array and calculate the sumleft to index of array and rightsum
	//step 3: if sumleft equal sumRight then return index
	A := []int{-7, 1, 5, 2, -4, 3, 0}
	lastIndex := len(A)
	prefixSum := make([]int, lastIndex)
	//Step 1: Start
	prefixSum[0] = A[0]
	for index := 1; index < lastIndex; index++ {
		prefixSum[index] = prefixSum[index-1] + A[index]
	}
	fmt.Println(prefixSum)
	//Step 1 : End
	//Step 2 : Start
	for index := 1; index < (lastIndex - 1); index++ {
		//leftSum will always at index of prefixSum
		leftSum := prefixSum[index-1]
		//righSum will be prefixSum of last index - [(index+1)-1 = index]
		rightSum := prefixSum[lastIndex-1] - prefixSum[index]
		//Step 3: Start
		if leftSum == rightSum {
			fmt.Println(index)
		}
		//Step 3: End
	}
	//Step 2: End
	fmt.Println(-1)

}
