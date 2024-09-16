package main

import "fmt"

func main() {
	s1, t1 := "abc", "ahbgdc"
	s2, t2 := "axc", "ahbgdc"

	fmt.Println(isSubsequence(s1, t1)) // Output: true
	fmt.Println(isSubsequence(s2, t2)) // Output: false
}

func isSubsequence(s string, t string) bool {
	sPointer, tPointer := 0, 0

	for sPointer < len(s) && tPointer < len(t) {
		if s[sPointer] == t[tPointer] {
			sPointer++
		}
		tPointer++
	}

	return sPointer == len(s)
}
