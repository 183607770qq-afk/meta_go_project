package main

import (
	"fmt"
)

// 查找字符串数组中的最长公共前缀
func main() {
	const size = 3
	arr := [size]string{}
	fmt.Printf("请输入 %d 个字符\n", size)
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	arr0 := arr[0]
	for i, _ := range arr0 {
		for _, s := range arr {
			if i == len(s) || arr0[i] != s[i] {
				fmt.Println("111", arr0[:i])
				return
			}
		}
	}
	fmt.Println("222", arr0)

	// fmt.Println("\n输入的数组:", arr)	// fmt.Println(arr[0][0], "\n\n", byte(arr[0][0]))
}
