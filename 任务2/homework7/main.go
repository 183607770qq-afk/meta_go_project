package main

import (
	"fmt"
	"time"
)

/*
编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
*/
func main() {
	c := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			c <- i
			time.Sleep(1 * time.Second)
		}
	}()
	for i := 1; i <= 10; i++ {
		i := <-c
		fmt.Printf("接收的参数=%d\n", i)
	}
}
