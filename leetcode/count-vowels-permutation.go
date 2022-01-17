package main

import "fmt"

/**
1220. 统计元音字母序列的数目
给你一个整数 n，请你帮忙统计一下我们可以按下述规则形成多少个长度为 n 的字符串：

字符串中的每个字符都应当是小写元音字母（'a', 'e', 'i', 'o', 'u'）
每个元音 'a' 后面都只能跟着 'e'
每个元音 'e' 后面只能跟着 'a' 或者是 'i'
每个元音 'i' 后面 不能 再跟着另一个 'i'
每个元音 'o' 后面只能跟着 'i' 或者是 'u'
每个元音 'u' 后面只能跟着 'a'
由于答案可能会很大，所以请你返回 模 10^9 + 7 之后的结果。



示例 1：

输入：n = 1
输出：5
解释：所有可能的字符串分别是："a", "e", "i" , "o" 和 "u"。
示例 2：

输入：n = 2
输出：10
解释：所有可能的字符串分别是："ae", "ea", "ei", "ia", "ie", "io", "iu", "oi", "ou" 和 "ua"。
示例 3：

输入：n = 5
输出：68
*/
func main() {
	fmt.Println(countVowelPermutation(144))

}

/**
算法描述略
*/
func countVowelPermutation(n int) int {
	var mod int = 1e9 + 7
	f := map[string]int{
		"a": 1,
		"e": 1,
		"i": 1,
		"o": 1,
		"u": 1,
	}
	for j := 1; j < n; j++ {
		a := f["a"]
		e := f["e"]
		i := f["i"]
		o := f["o"]
		u := f["u"]
		f["a"] = (e + i + u) % mod
		f["e"] = (a + i) % mod
		f["i"] = (e + o) % mod
		f["o"] = i % mod
		f["u"] = (i + o) % mod
	}
	sum := 0
	for _, v := range f {
		sum = (sum + v) % mod
	}
	return sum
}
