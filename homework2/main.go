package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Println("请输入整数")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("输入错误")
		return
	}

	if input < 0 {
		return
	}
	inputTem := input
	var reversal int
	for input > 0 {
		reversal = reversal*10 + input%10
		input /= 10
	}
	if inputTem == reversal {
		fmt.Println("是回文数")
	} else {
		fmt.Println("不是回文数")
	}

}
