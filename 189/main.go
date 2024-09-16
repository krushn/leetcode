package main

/*
func main() {

		nums := []int{1, 2, 3, 4, 5, 6, 7}
		k := 3

		rotate(nums, k)
	}

func rotate(nums []int, k int) {

	length := len(nums)

	for i := 0; i < length; i++ {

		index := i - k

		if index < 0 {
			index += length

			tmp := nums[i]
			nums[i] = nums[index]
			nums[index] = tmp
		}
		//i++
	}

	fmt.Print(nums)

}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	for _, n := range nums {
		go func() {
			fmt.Println(n)
		}()
	}
}

func swap(a *int, b *int) (int, int) {
	tmp := *a
	*a = *b
	*b = tmp

	//return a, b
	return *a, *b
}*/

func main() {
	//a := 10
	//b := 20
	//fmt.Println(swap(&a, &b))

	numbs := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	aLastIndex := len(numbs) - k
	a := numbs[:aLastIndex]
	b := numbs[aLastIndex:]

	result := append(b, a...)

	return result
}
