package main

import "fmt"

func swap(a, b *int32) {
	tempa := *a
	*a = *b
	*b = tempa
}

func reverseArray(a []int32) []int32 {
	arraySize := len(a)
	for index := 0; index < arraySize/2; index++ {
		swap(&a[index], &a[arraySize-index-1])
	}
	return (a)

}
func main() {
	arr := []int32{0, 1, 2, 3, 4}

	fmt.Println(reverseArray(arr))
}
