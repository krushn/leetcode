package main

import (
	"fmt"
)

func main() {

	nums := []int{1, 1, 2}

	response := removeDuplicates(nums)

	fmt.Printf("%v", response)
}

func removeDuplicates(nums []int) int {
	i := 0

	for j, _ := range nums {

		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}
