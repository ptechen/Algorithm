//题目描述
//给定一个二维数组，其每一行从左到右递增排序，从上到下也是递增排序。给定一个数，判断这个数是否在该二维数组中。
//
//Consider the following matrix:
//[
//[1,   4,  7, 11, 15],
//[2,   5,  8, 12, 19],
//[3,   6,  9, 16, 22],
//[10, 13, 14, 17, 24],
//[18, 21, 23, 26, 30]
//]
//
//Given target = 5, return true.
//Given target = 20, return false.
//复制到剪贴板错误已复制
//解题思路
//要求时间复杂度O（M + N），空间复杂度O（1）。其中M为行数，N为列数。
//
//该二维数组中的一个数，小于它的数一定在其左边，大于它的数一定在其下边。因此，从右上角开始查找，就可以根据目标和当前元素的大小关系来缩小查找区间，当前元素的查找区间为左下角的所有元素。

package main

import "fmt"

func find(array [][]int, target int) (key bool, err error) {
	if array == nil || len(array) == 0 || len(array[0]) == 0 {
		return key, fmt.Errorf("array is nil or length is 0")
	}
	rows := len(array)
	columns := len(array[0])
	r := 0
	c := columns - 1
	for r <= rows -1 && c >= 0 {
		if target == array[r][c] {
			return true, nil
		} else if target > array[r][c] {
			r ++
		} else {
				c --
		}
	}

	return  false, nil
}

func main()  {
	array := [][]int{
		[]int{1,   4,  7, 11, 15},
		[]int{2,   5,  8, 12, 19},
		[]int{3,   6,  9, 16, 22},
		[]int{10, 13, 14, 17, 24},
		[]int{18, 21, 23, 26, 30},
	}
	ex ,err := find(array, 14)
	if err != nil {
		panic(err)
	}
	fmt.Println(ex)
}