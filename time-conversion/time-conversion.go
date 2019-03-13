package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertString(s string) string {
	splitedDateArr := strings.Split(s, ":")
	first, last := splitedDateArr[0], splitedDateArr[2][2:4]
	firstInt, err := strconv.Atoi(first)
	if err == nil && last == "PM" && firstInt < 12 {
		firstInt += 12
	} else if last == "AM" && firstInt == 12 {
		firstInt -= 12
	}
	firstString := strconv.Itoa(firstInt)
	if firstInt < 10 {
		firstString = "0" + firstString
	}
	splitedDateArr[0] = firstString
	splitedDateArr[2] = splitedDateArr[2][0:2]
	return (strings.Join(splitedDateArr, ":"))

}
func main() {
	fmt.Println(convertString("12:05:45AM"))

}
