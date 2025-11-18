package main

import "fmt"

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，
将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
*/
func main() {
	nums := [][]int{
		{2, 3},
		{1, 2},
		{7, 9},
		{4, 6},
		{2, 7},
	}

	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n-i; j++ {
			if nums[j-1][0] > nums[j][0] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
	// for i, v := range nums {
	// 		if v[0] < nums[i+1][0] {
	// 			tem = v
	// 			v = nums[j]
	// 			nums[j] = tem
	// 		}

	// }
	arr := [][]int{
		{nums[0][0], nums[n-1][1]},
	}

	fmt.Println(arr)

}
