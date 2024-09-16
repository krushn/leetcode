package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {

	result := isHappy(19)

	println(result)
}

func isHappy(target int) bool {
	return getNumber(target)
}

func getNumber(num int) bool {

	time.Sleep(1 * time.Second)

	println("getNumber called with %v", num)

	s := fmt.Sprintf("%d", num)

	println("string", s, len(s))

	sum := 0

	strings.Map(func(r rune) rune {
		sum += int(math.Pow(float64(r), 2))
		return r
	}, s)

	//for i := 0; i < len(s); i++ {
	//println(s[i], float64(s[i]), math.Pow(float64(s[i]), 2))
	//	println("inv numv", i, strings.Fields())
	//sum += int(math.Pow(float64(s[i]), 2))
	//}

	println(sum)

	//return sum > 0

	/*if sum == 1 {
		return true
	} else if sum < 10 {
		return false
	} else {
		return getNumber(sum)
	}*/

	return false
}
