package main

import (
	"fmt"
	"sync"
	"time"
)

/*
设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
*/

func main() {
	var wg sync.WaitGroup
	// start := time.Now()
	tasks := []func(){a, b, c}
	for i, task := range tasks {
		wg.Add(1)
		go func(id int, task func()) {
			defer wg.Done()
			start := time.Now()
			task()
			fmt.Printf("任务%d 的执行时间 %v\n", i, time.Since(start))
		}(i, task)

	}
	wg.Wait()
}
func a() {
	fmt.Println("任务一执行")
	time.Sleep(500 * time.Millisecond)
}

func b() {
	fmt.Println("任务二执行")
	time.Sleep(800 * time.Millisecond)
}
func c() {
	fmt.Println("任务三执行")
	time.Sleep(600 * time.Millisecond)
}
