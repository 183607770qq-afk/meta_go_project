package main

import "fmt"

/*
实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
*/
func main() {
	arr := []int{2, 3, 4, 5, 6}
	receiveArr(&arr)
	fmt.Println(arr)
}
func receiveArr(receiveArr *[]int) {
	for i, v := range *receiveArr {
		// fmt.Println(i, v)
		(*receiveArr)[i] = v * 2
	}
	// fmt.Println(*receiveArr)
}
