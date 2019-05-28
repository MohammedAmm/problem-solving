package main

import "fmt"

//Check if number is greater than zero return zero , if equal zero return zero
// func zeroOrOne(number int32) int32 {
// 	if number > 0 {
// 		return 1
// 	}
// 	return 0
// }

//XOR Function
// func XOR(a, b int32) int {
// 	//Check if x,lastAnswer is greater than zero

// 	return (int(a) || int(b)) && !(int(a) && int(b))
// }

// Complete the dynamicArray function below.
func dynamicArray(n int32, queries [][]int32) []int32 {
	//Array that contains last answers
	var lastAnswers []int32
	var lastAnswer int32
	s := make([][]int32, int(n))
	for i := 0; i < int(n); i++ {
		s[i] = make([]int32, 0)
	}
	//Iterate for each query
	for _, query := range queries {
		//Assign query type to executedQuery
		executedQuery := query[0]
		x, y := query[1], query[2]

		//Make Exclusive or for result x,lastAnswer from above line
		index := (int(x) ^ int(lastAnswer)) % int(n)
		// index := logicalResult % n
		switch executedQuery {
		case 1:
			{
				s[index] = append(s[index], y)
			}
		case 2:
			{
				lastAnswer = s[index][int(y)%len(s[index])]
				lastAnswers = append(lastAnswers, lastAnswer)
			}
		}
	}
	return lastAnswers
}

func main() {
	arr := [][]int32{{1, 0, 5},
		{1, 1, 7},
		{1, 0, 3},
		{2, 1, 0},
		{2, 1, 1}}
	fmt.Println(dynamicArray(2, arr))

}
