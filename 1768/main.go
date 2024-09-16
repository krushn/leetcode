package main

import (
	"fmt"
	"math"
)

func main() {
	word1 := "abcd"
	word2 := "pqr"

	fmt.Println(mergeAlternately(word1, word2))
}

func mergeAlternately(word1 string, word2 string) string {

	word1Length := len(word1)
	word2Length := len(word2)

	minLength := math.Min(float64(word1Length), float64(word2Length))

	result := ""

	for i := 0; i < int(minLength); i++ {
		result += string(word1[i])
		result += string(word2[i])
	}

	if word1Length > word2Length {
		result += word1[int(minLength):]
	} else if word2Length > word1Length {
		result += word2[int(minLength):]
	}

	return result
}
