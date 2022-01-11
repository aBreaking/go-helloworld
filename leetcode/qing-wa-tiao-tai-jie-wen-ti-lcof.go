package main

import "fmt"

/**

n 级的台阶总共有多少种跳法。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

示例 1：

输入：n = 2
输出：2
示例 2：

输入：n = 7
输出：21
示例 3：

输入：n = 0
输出：1
提示：

0 <= n <= 100

来源：力扣（LeetCode）
链接：

https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
https://leetcode-cn.com/problems/climbing-stairs/solution/pa-lou-ti-by-leetcode-solution/
*/
func main() {
	fmt.Println(numWays(100))
}

/**
f(n) 表示 跳上n级台接需要的次数，那么最后一跳有两种情况：
剩一个台阶，那么在此之前共有f(n-1)种跳法；
剩两个台阶，那么在此之前共有f(n-2)种跳法；

所以：f(n) = f(n-1)+f(n-2)  => f(n+1) = f(n) + f(n-1)
这是一个很明显的斐波那契数列，只不过初始值不一样
f(0) = 1 ; f(1) = 1 ; f(2) = 2

f(n)是由 f(n-1) 和f(n-2) 状态转移而来，所以只需要定义三个变量即可


*/
func numWays(n int) int {
	fnp1, fn, fns1 := 0, 1, 1 //f(n+1) = f(n) + f(n-1)
	for i := 1; i < n; i++ {  // f1 f0的初始值都已经给出，所以i的下标是从1开始
		fnp1 = (fn + fns1) % 1000000007
		fns1 = fn
		fn = fnp1
	}
	return fn

}
