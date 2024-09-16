package main

func main() { {

}

func removeDuplicates(nums []int) int {

	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[i+1] && nums[i + 1] == nums[i+ 2] {
			// replace append with in-place
			nums = append(nums[:i + 1], nums[i+2:]...)
			i--
		}
	}

	return nums
}



