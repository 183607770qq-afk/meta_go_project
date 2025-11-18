package main

import "fmt"

func main() {

	arr := []int{12, 32, 23, 32, 12}
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	for k, v := range m {
		fmt.Println(k, "===", v)
	}

}
