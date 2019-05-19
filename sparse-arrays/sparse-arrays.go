package main

import "fmt"

// Complete the matchingStrings function below.
func matchingStrings(strings []string, queries []string) []int32 {
	//Initialize array with zeros (The length equal to the length of quaries array)
	result := make([]int32, len(queries))
	//Loop through quaries to find match
	for index, query := range queries {
		//Loop through strings to match with the query
		for _, currentString := range strings {
			//check if the query equals current string
			//Increase the index of matched query in result array by
			if query == currentString {
				result[index]++
			}
		}
	}
	return result
}

func main() {
	strings := []string{
		"abcde",
		"sdaklfj",
		"asdjf",
		"na",
		"basdn",
		"sdaklfj",
		"asdjf",
		"na",
		"asdjf",
		"na",
		"basdn",
		"sdaklfj",
		"asdjf",
	}
	queries := []string{
		"abcde",
		"sdaklfj",
		"asdjf",
		"na",
		"basdn",
	}
	fmt.Println(matchingStrings(strings, queries))

}
