package main

import (
	"fmt"
	"sort"
)

/**
给你一个整数数组 nums 和一个整数 k ，按以下方法修改该数组：

选择某个下标 i 并将 nums[i] 替换为 -nums[i] 。
重复这个过程恰好 k 次。可以多次选择同一个下标 i 。

以这种方式修改数组后，返回数组 可能的最大和 。


示例 1：

输入：nums = [4,2,3], k = 1
输出：5
解释：选择下标 1 ，nums 变为 [4,-2,3] 。
示例 2：

输入：nums = [3,-1,0,2], k = 3
输出：6
解释：选择下标 (1, 2, 2) ，nums 变为 [3,1,0,2] 。
示例 3：

输入：nums = [2,-3,-1,5,-4], k = 2
输出：13
解释：选择下标 (1, 4) ，nums 变为 [2,3,-1,5,4] 。


提示：

1 <= nums.length <= 104
-100 <= nums[i] <= 100
1 <= k <= 104

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximize-sum-of-array-after-k-negations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	var nums = []int{3, -1, 0, 2}
	var k = 3
	fmt.Println(largestSumAfterKNegations(nums, k))
}

/**
思路一：  nums排序，分析前k的元素.
 1. 如果负数开头，那么直接先把 负数抵消；
 2. 抵消后，再对nums排序，
 3. 如果k还有剩余，那么说明 此时nums以0或正数开头，那么开头的数自相抵消即可。
*/
func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	sum := 0
	for i, n := range nums {
		if n < 0 && k > 0 {
			nums[i] = n * -1
			k--
		}
		sum += nums[i]
	}

	sort.Ints(nums)

	if k%2 != 0 {
		sum -= 2 * nums[0]
	}
	return sum
}
