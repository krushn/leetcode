package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{3, 2, 2, 3}
	val := 3

	response := removeElement(nums, val)

	fmt.Printf("%v", response)
}

func removeElement(nums []int, val int) int {

	count := len(nums)

	for i, v := range nums {
		if v == val {
			nums[i] = 105
			count--
		}
	}

	sort.Ints(nums)

	return count
}
