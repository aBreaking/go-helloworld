package main

import "fmt"

func main() {
	fmt.Println(getX2(3))
	var a, b = getX2AndX3(3)
	x, y := getX2AndX3_2(3)

	fmt.Println(a, b, x, y)
}

func getX2(input int) int {
	return input * 2
}

func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	return
}
