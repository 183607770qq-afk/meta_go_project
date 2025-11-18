package main

import "fmt"

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
*/
func main() {
	nums := []int{2, 3, 4, 5, 6}
	target := 7
	// for i, v := range nums {
	// 	for j := i + 1; j < len(nums); j++ {
	// 		if v+nums[j] == target {
	// 			arr := []int{i, j}
	// 			fmt.Println(arr)
	// 			return
	// 		}
	// 	}
	// }
	m := map[int]int{}
	for i, v := range nums {
		if p, ok := m[target-v]; ok {
			arr := []int{p, i}
			fmt.Println(arr)
			return
		}
		m[v] = i

	}
}
