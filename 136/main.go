package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
}

func singleNumber(nums []int) int {
	mask := 0

	for _, num := range nums {
		mask ^= num
	}

	return mask
}
