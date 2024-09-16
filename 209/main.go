package main

import (
	"fmt"
	"math"
)

func main() {
	target := 11
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println("output", minSubArrayLen(target, nums))
}

func minSubArrayLen(target int, nums []int) int {

	sum := 0
	lengthInCurrentWindow := math.Inf(+1)

	start := 0

	for i := 0; i < len(nums); i++ {

		sum += nums[i]

		for sum >= target {

			lengthInCurrentWindow = math.Min(lengthInCurrentWindow, float64(i-start+1))

			sum -= nums[start]
			start++
		}
	}

	if lengthInCurrentWindow == math.Inf(+1) {
		return 0
	}

	return int(lengthInCurrentWindow)
}
