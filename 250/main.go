package main

func main() {
	//a := []int{2, 7, 11, 15} //3, 2, 4
	a := []int{3, 3}

	result := twoSum(a, 6)

	println(result[0], result[1])
}

func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {

		for j := i + 1; j < len(nums); j++ {

			if nums[i]+nums[j] == target {
				return []int{
					i,
					j,
				}
			}
		}
	}

	return []int{}
}
