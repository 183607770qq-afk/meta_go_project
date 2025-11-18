package main

import (
	"fmt"
)

// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func main() {
	intArr := []int{8, 9, 9}
	for i := 1; i <= len(intArr); i++ {
		if intArr[len(intArr)-i] != 9 {
			intArr[len(intArr)-i] = intArr[len(intArr)-i] + 1
			fmt.Println(intArr)
			return
		} else {

			intArr[len(intArr)-i] = 0
			if i == len(intArr) {
				intArr = append([]int{1}, intArr...)
				fmt.Println(intArr)

				return
			}
			// intArr[len(intArr)-i]
		}
	}
	fmt.Println(intArr)
}
