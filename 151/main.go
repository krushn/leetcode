package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	s := "a good   example"
	fmt.Println(reverseWords(s))
}

func reverseWords(s string) string {
	/*s = strings.Trim(s, " ")
	s = strings.ReplaceAll(s, "  ", " ")
	a := strings.Split(s, " ")
	slices.Reverse(a)
	return strings.Join(a, " ")
	*/

	arrString := strings.Split(s, " ")
	slices.Reverse(arrString)

	var arrResult []string

	//for i := len(s) - 1; i >= 0; i-- {
	for _, word := range arrString {
		trimedWord := strings.ReplaceAll(word, " ", "")
		if trimedWord != "" {
			arrResult = append(arrResult, word)
		}
	}

	return strings.Join(arrResult, " ")
}
