//题目描述
//在一个长度为n的数组里的所有数字都在0到n-1的范围内。数组中某些数字是重复的，但不知道有几个数字是重复的，也不知道每个数字重复几次。请找出数组中任意一个重复的数字。
//解题思路
//要求时间复杂度O（N），空间复杂度O（1）。因此不能使用排序的方法，也不能使用额外的标记数组。
//
//对于这种数组元素在[0，n-1]范围内的问题，可以将值为我的元素调整到第i个位置上进行求解。
//
//以（2,3,1,0,2,5）为例，遍历到位置4时，该位置上的数为2，但是第2个位置上已经有一个2的值了，因此可以知道2重复：
package main

import "fmt"

func findFirstRepeatingNum(array []int) (num int, err error) {
	if array == nil || len(array) == 0 {
		return 0, fmt.Errorf("array is nil or length is 0")
	}
	for i := 0; i < len(array); i ++{
		for array[i] != i {
			if array[i] == array[array[i]]  {
				return array[array[i]], err
			}
			array[i], array[array[i]] = array[array[i]], array[i]
		}
	}
	return 0, fmt.Errorf("无重复数字")
}

func main()  {
	array := []int{2,3,1,0,2,5}
	num, err := findFirstRepeatingNum(array)
	if err != nil {
		panic(err)
	}
	fmt.Println(num)
}