package main

import (
	"fmt"
)

/**
693. 交替位二进制数
给定一个正整数，检查它的二进制表示是否总是 0、1 交替出现：换句话说，就是二进制表示中相邻两位的数字永不相同。



示例 1：

输入：n = 5
输出：true
解释：5 的二进制表示是：101
示例 2：

输入：n = 7
输出：false
解释：7 的二进制表示是：111.
示例 3：

输入：n = 11
输出：false
解释：11 的二进制表示是：1011.

*/
func main() {

	fmt.Println(hasAlternatingBits(2))
	fmt.Println(hasAlternatingBits(5))
	fmt.Println(hasAlternatingBits(6))
	fmt.Println(hasAlternatingBits(4))
	fmt.Println(hasAlternatingBits(10))
	fmt.Println(hasAlternatingBits(7))
	fmt.Println(hasAlternatingBits(11))
	fmt.Println(hasAlternatingBits(42))
}

func hasAlternatingBits(n int) bool {

	for i := 0; i <= n*2; i = i<<2 + 2 {

		if i == n || i == n*2 {
			return true
		}
	}
	return false
}
