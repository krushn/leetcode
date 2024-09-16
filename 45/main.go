package main

import "fmt"

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4})) // Output: 2
	fmt.Println(jump([]int{2, 3, 0, 1, 4})) // Output: 2
}

func jump(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	jumps := 0
	currentEnd := 0
	farthest := 0

	for i := 0; i < n-1; i++ {
		farthest = max(farthest, i+nums[i])
		
		if i == currentEnd {
			jumps++
			currentEnd = farthest
		}

		if currentEnd >= n-1 {
			return jumps
		}
	}

	return jumps
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
