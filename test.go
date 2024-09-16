package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	var a uint64 = 10
	var b uint64 = 20

	var c int64 = int64(math.Ceil(float64(a) / float64(b)))

	log.Println(c)

	fmt.Printf("%v", a/b)

}
