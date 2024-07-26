package main

import (
	"demo/wauo"
	"fmt"
)

func main() {
	fmt.Print("t1  ==>  ")
	t1()
	fmt.Print("t2  ==>  ")
	t2()

}

func t1() {
	ch := make(chan int, 1)
	close(ch)
	v, ok := <-ch
	if ok == false {
		wauo.PyPrint("管道已被关闭，并且管道里没有值了。零值={}", v)
	} else {
		wauo.PyPrint("拿到值={}", v)
	}
}

func t2() {
	ch := make(chan int, 1)
	ch <- 200
	close(ch)
	v, ok := <-ch
	if ok == false {
		wauo.PyPrint("管道已被关闭，并且管道里没有值了。零值={}", v)
	} else {
		wauo.PyPrint("拿到值={}", v)
	}
}
