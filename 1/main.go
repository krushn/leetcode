package main

import (
	"fmt"
	"slices"
)

func merge(nums1 []int, m int, nums2 []int, n int) {

	a := nums1[0:m]

	nums1 = append(a, nums2...)

	slices.Sort(nums1)

	fmt.Printf("%v", nums1)
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	merge(nums1, m, nums2, n)
}
