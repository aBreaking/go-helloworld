package main

import (
	"fmt"
)

func main() {
	fmt.Println(getX2(3))
	var a, b = getX2AndX3(3)
	x, y := getX2AndX3_2(3)
	fmt.Println(a, b, x, y)

	fmt.Println(minX(2, 5, 6, 3, 2))
	slice := []int{5, 6, 3, 7}
	//如果参数被存储在一个 slice 类型的变量 slice 中，则可以通过 slice... 的形式来传递参数，调用变参函数。
	fmt.Println(minX(2, slice...))

	deferFunc()

	callback(2, add)
}

/**
函数参数与返回值
https://fuckcloudnative.io/the-way-to-go/06.2.html
*/
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

/**
将函数作为参数
https://fuckcloudnative.io/the-way-to-go/06.7.html
*/
func callback(a int, f func(int, int) int) {
	var v = f(a, 3)
	fmt.Println("callback result is", v)
}

func add(a, b int) int {
	return a + b
}

/**
6.3 传递变长参数
from : https://fuckcloudnative.io/the-way-to-go/06.3.html
*/
func minX(x int, s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, i := range s {
		if i < min {
			min = i
		}
	}
	return x * min
}

type Options struct {
	p1 int
	p2 string
}

/**
如果一个变长参数的类型没有被指定，则可以使用默认的空接口 interface{}，这样就可以接受任何类型的参数（详见第 11.9 节）。该方案不仅可以用于长度未知的参数，还可以用于任何不确定类型的参数。一般而言我们会使用一个 for-range 循环以及 switch 结构对每个参数的类型进行判断：
*/
func typecheck(values ...interface{}) []string {
	var s = make([]string, len(values))
	for _, value := range values {
		switch value.(type) {
		case int:
			s = append(s, "int")
		case string:
			s = append(s, "string")
		default:
			s = append(s, "string")

		}
	}
	return s
}

/**
defer的用法， 详见defer.go
*/
func deferFunc() {
	defer fun2()
	fmt.Println("start func")
	fmt.Println("end func")
}

func fun2() {
	fmt.Println("this is fun2 ")
}

/**
匿名函数
*/
func h() {
	var a = func(b int) int {
		return b * b
	}
	var av = a(3)
	fmt.Println(av)

	/**
	闭包
	*/
	func(x, y int) {
		var c = av + x + y
		fmt.Printf("c is %d", c)
	}(3, a(2))
}
