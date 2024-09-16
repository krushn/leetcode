package main

import "fmt"

func main() {
	nums := []int{3, 2, 1, 0, 4} // //
	fmt.Println(canJump(nums))
}

func canJump(nums []int) bool {
	farthest := 0
	for i := 0; i < len(nums); i++ {
		if i > farthest {
			return false
		}
		farthest = max(farthest, i+nums[i])
		if farthest >= len(nums)-1 {
			return true
		}
	}
	return farthest >= len(nums)-1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
