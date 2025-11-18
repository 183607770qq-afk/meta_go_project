package main

import (
	"fmt"
)

func main() {
	//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
	fmt.Println("请输入")
	var inputLine string
	_, err := fmt.Scanln(&inputLine)
	if err != nil {
		fmt.Println("输入错误")
	}
	leninputLine := len(inputLine)
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < leninputLine; i++ {
		if m[inputLine[i]] > 0 {
			fmt.Println(len(stack))
			fmt.Println(string(stack[len(stack)-1]), "===", string(m[inputLine[i]]), i)
			if len(stack) == 0 || stack[len(stack)-1] != m[inputLine[i]] {
				fmt.Println("无效字符串")
				return
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, inputLine[i])
		}
	}

	if len(stack) == 0 {
		fmt.Println("字符串有效")
		return
	}

}
