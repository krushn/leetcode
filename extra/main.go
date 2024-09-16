package main

import "fmt"

type Person struct {
	Left bool
	Name string
}

func main() {
	a := []Person{Person{Left: true, Name: "Abdul"}, Person{Left: true, Name: "Karim"}, Person{Left: false, Name: "Choksi"}}

	var left []interface{}
	var right []interface{}

	for i := 0; i < len(a); i++ {

		if a[i].Left {
			left = append(left, a[i])
		} else {
			right = append(right, a[i])
		}
	}

	fmt.Println(left, right)
}
