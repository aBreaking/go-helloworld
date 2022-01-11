package main

import (
	"fmt"
)

func main() {

	var nums = []int{3, 6, 2, 8, 1}
	fmt.Println(nums)
	swap(nums, 1, 4)
	fmt.Println(nums)

}

/**
快速排序
*/

func FastSort(nums *[]int, start int, end int) {

	for start <= end {

	}

}

func swap(nums []int, i int, j int) {
	var p = nums
	var temp = p[i]
	p[i] = p[j]
	p[j] = temp
}
