package main

import "fmt"

/**
from : https://leetcode-cn.com/problems/reordered-power-of-2/
869. 重新排序得到 2 的幂
给定正整数 N ，我们按任何顺序（包括原始顺序）将数字重新排序，注意其前导数字不能为零。

如果我们可以通过上述方式得到 2 的幂，返回 true；否则，返回 false。



示例 1：

输入：1
输出：true
示例 2：

输入：10
输出：false
/示例 3：

输入：16
输出：true
示例 4：

输入：24
输出：false
示例 5：

输入：46
输出：true


提示：

1 <= N <= 10^9
*/

func main() {
	var n int = 215
	fmt.Println(reorderedPowerOf2(n))
}

var map2 = make(map[[10]int]bool)

/**
算法思路：提示给的思路是1 <= N <= 10^9，先把10^9所有2的幂数全部找出来。然后记录这些数的每个字数出现的字数。然后再将n进行匹配即可。
*/
func reorderedPowerOf2(n int) bool {

	for i := 1; i < 1e9; i <<= 1 {
		map2[c2(i)] = true
	}
	return map2[c2(n)]
}

func c2(n int) [10]int {

	s := [10]int{}

	for n > 0 {
		s[n%10]++
		n /= 10
	}

	return s

}
