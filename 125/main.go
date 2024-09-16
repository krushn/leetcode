package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	s := "Marge, let's \"[went].\" I await {news} telegram."
	//"race a car"
	//"A man, a plan, a canal: Panama"
	println(isPalindrome(s))
}
func isPalindrome(s string) bool {

	// This pattern matches any character that is not a letter, number, or space
	reg, err := regexp.Compile(`[^a-zA-Z0-9\s]+`)
	if err != nil {
		fmt.Println(err)
		return false
	}

	// Replace all special characters with an empty string
	s = reg.ReplaceAllString(s, "")
	s = strings.ReplaceAll(s, " ", "")

	s = strings.ToLower(s)

	j := len(s) - 1
	for i := 0; i < len(s); i++ {

		if s[i] != s[j] {
			return false
		}

		j--
	}

	return true
}
