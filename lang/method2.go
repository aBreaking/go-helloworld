package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {
	// make an Add2 function, give it a name p2, and call it:
	p2 := Add2(1)
	i, i2 := p2(3)
	fmt.Println("Call Add2  ", i, i2)

	var adder = adder()
	fmt.Println(adder(1))
	fmt.Println(adder(10))
	fmt.Println(adder(100))

	whereCode()

	elapse()
}

// Add2 /**
//6.9 应用闭包：将函数作为返回值
//https://fuckcloudnative.io/the-way-to-go/06.9.html**/
func Add2(x int) func(b int) (int, int) {
	return func(b int) (int, int) {
		return b + x, b * b
	}
}

/**
在多次调用中，变量 x 的值是被保留
https://fuckcloudnative.io/the-way-to-go/06.10.html
*/
func adder() func(int) int {
	var x int
	return func(i int) int {
		x += i
		return x
	}
}

/**
使用闭包调试
from : https://fuckcloudnative.io/the-way-to-go/06.10.html
*/
func whereCode() {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
	// some code
	where()
	// some more code
	where()
}

/**
6.11 计算函数执行时间
https://fuckcloudnative.io/the-way-to-go/06.11.html
*/
func elapse() {

	var start = time.Now()

	time.Sleep(time.Second)

	var end = time.Now()

	var delta = end.Sub(start)

	fmt.Println(delta)
}
