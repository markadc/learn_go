package main

import (
	. "demo/wauo"
	"sync"
	"time"
)

// 等待组。等待协程完成后，主程序再退出
func main() {
	ch := make(chan int)
	var wg sync.WaitGroup // 等待组

	wg.Add(2)
	go cons(ch, &wg)
	go prod(ch, &wg)
	wg.Wait() // 等待计数为0，再退出主程序
}

func prod(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // 计数-1
	defer close(ch)
	for i := 0; i < 5; i++ {
		ch <- i
		RandomSleep(0, 1)
	}
}

func cons(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // 计数-1
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				PyPrint("管道已关闭")
				return
			} else {
				PyPrint("v={}", v)
			}
		case <-time.After(3 * time.Second):
			PyPrint("发生超时")
		}
	}
}
