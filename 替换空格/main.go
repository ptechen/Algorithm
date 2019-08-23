//题目描述
//将一个字符串中的空格替换成“％20”。
//
//Input:
//"A B"
//
//Output:
//"A%20B"
//复制到剪贴板错误已复制
//解题思路
//在字符串尾部填充任意字符，使得字符串的长度等于替换之后的长度。因为一个空格要替换成三个字符（％20），因此当遍历到一个空格时，需要在尾部填充两个任意字符。
//
//令P1指向字符串原来的末尾位置，P2指向字符串现在的末尾位置.P1和P2从后向前遍历，当P1遍历到一个空格时，就需要令P2指向的位置依次填充02％（注意是逆序的），否则就填充上P1指向字符的值。
//
//从后向前遍是为了在改变P2所指向的内容时，不会影响到P1遍历原来字符串的内容。
//

package main

import "fmt"

func main()  {
	str := "AB AB"
	byteStr := []byte(str)
	p1 := len(byteStr) - 1
	byteStr = append(byteStr, []byte("  ")...)
	p2 := len(byteStr) - 1
	for p1 >= 0 && p2 > p1 {
		c := byteStr[p1]
		p1 --
		if c == ' ' {

			byteStr[p2] = '0'
			p2 --

			byteStr[p2] = '2'
			p2 --

			byteStr[p2] = '%'
			p2 --

		} else {
			byteStr[p2] = c
			p2 --
		}
	}
	fmt.Println(string(byteStr))
}