package main

import "fmt"

func main() {

	var nums = []int{3, 1, 5, 2, 4, 5, 6}
	var ret = 0
	for i := 0; i < len(nums); i++ {
		ret ^= nums[i] ^ (i)
	}
	fmt.Println(ret)

}

func swap(x, y int) (int, int) {
	x = x ^ y
	y = x ^ y
	x = x ^ y
	return x, y
}
