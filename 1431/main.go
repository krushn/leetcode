package main

import (
	"fmt"
)

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3

	fmt.Println(kidsWithCandies(candies, extraCandies))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {

	// Find the maximum number of candies any kid currently has
	maxCandies := 0
	for _, candy := range candies {
		if candy > maxCandies {
			maxCandies = candy
		}
	}

	// Initialize the result array
	result := make([]bool, len(candies))

	// Determine if each kid can have the greatest number of candies
	for i, candy := range candies {
		if candy+extraCandies >= maxCandies {
			result[i] = true
		} else {
			result[i] = false
		}
	}

	return result
}
