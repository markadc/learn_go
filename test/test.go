package main

import (
	. "demo/wauo"
	"fmt"
	rand2 "math/rand/v2"
)

func main() {
	s := make([]int, 0)
	s2 := append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	for i, v := range s2 {
		PyPrint("索引={} 值={}", i, v)
	}
	fmt.Println(s)
	fmt.Println(s2)
	PyPrint("预期{}实际{}", 200, 429)

	for i := 0; i < 10; i++ {
		fmt.Print(rand2.IntN(100), ", ")

	}
	fmt.Println()
	for i := 0; i < 10; i++ {
		fmt.Print(Randint(1, 100), ", ")
	}

}
