package main

func main() {
	a := []int{1, 2, 2, 5, 7, 7, 7}

	majorityElement(a)
}

func majorityElement(nums []int) int {

	var count map[int]int = map[int]int{}

	for _, v := range nums {
		count[v]++
	}

	var maxCount int = 0
	var result int = 0

	for i, count := range count {
		if count > maxCount {
			maxCount = count
			result = i
		}
	}

	return result
}
