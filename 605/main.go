package main

import "fmt"

func main() {
	flowerbed := []int{1, 0, 0, 0, 0, 1}
	n := 2

	fmt.Println(canPlaceFlowers(flowerbed, n))
}

func canPlaceFlowers(flowerbed []int, n int) bool {

	nums := 0

	for i := 0; i < len(flowerbed); i++ {
		//make sure neighbour is empty
		if flowerbed[i] == 0 &&
			(i == 0 || flowerbed[i-1] == 0) &&
			(i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			nums++ //plant one
			i++    //skip one
		}
	}

	return nums >= n
}
