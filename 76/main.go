package main

import (
	"fmt"
	"math"
)

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"

	fmt.Printf("output", minWindow(s, t))
}

func minWindow(s string, t string) string {

	j := 0
	matchStart := math.Inf(1)

	for i := 0; i < len(s); i++ {
		println("i", i)

		if s[i] == t[j] {

			println("matching char at ", i)

			matchStart = math.Min(matchStart, float64(i))

			if j == len(t)-1 {

				//startIndex := matchStart == 0 ? 0: matchStart - 1:
				//return s[startIndex:startIndex + len(t) + 1]

				println("matched all char", matchStart, i, j)

				return s[int(matchStart) : int(matchStart)+len(t)]
			}

			j++
			continue
		} else {

			println("mismatch char at /resetting ", i)

			j = 0
			matchStart = math.Inf(1)
		}
	}

	return ""
}
