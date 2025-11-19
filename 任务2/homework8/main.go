package main

import (
	"fmt"
	"time"
)

/*
实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
*/
func main() {
	c := make(chan int, 100)
	go func() {
		i := 0
		for {
			i++
			c <- i
			fmt.Printf("存入数据=%d\n", i)
			time.Sleep(time.Millisecond * 100)
		}
	}()
	go func() {
		for {
			num := <-c
			fmt.Printf("取出数据=%d\n", num)
			time.Sleep(time.Millisecond * 500)

		}
	}()
	time.Sleep(time.Second * 60)

}
