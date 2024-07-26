package main

import (
	. "demo/wauo"
	"sync"
)

var (
	num int
	mu  sync.Mutex
	wg  sync.WaitGroup
)

// 使用互斥锁
func main() {
	PyPrint("num={}", num)

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go task()
	}
	wg.Wait()

	PyPrint("num={}", num)

}

func task() {
	// 计数-1
	defer wg.Done()

	// 协程可能出现异常
	defer func() {
		if r := recover(); r != nil {
			PyPrint("异常={}", r)
		}
	}()

	// 具体任务
	for i := 0; i < 10000; i++ {
		if iv := Randint(1, 10000); iv == 404 {
			panic("404")
		}

		mu.Lock()
		num++
		mu.Unlock()
	}
}
