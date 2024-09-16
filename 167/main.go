package main

import (
	"fmt"
)

func main() {
	target := 9
	nums := []int{2, 7, 11, 15}

	fmt.Println("output", twoSum(nums, target))
}

func twoSum(numbers []int, target int) []int {

	for i := 0; i < len(numbers); i++ {

		for j := i + 1; j < len(numbers); j++ {

			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}

	return []int{0, 0}
}
