package main

import (
	"fmt"
	"time"
)

/*
编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
*/

func main() {
	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 1 {
				fmt.Println("奇数=", i)
			}
			time.Sleep(1 * time.Second)

		}
	}()
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("偶数=", i)
		}
		time.Sleep(1 * time.Second)

	}

}
