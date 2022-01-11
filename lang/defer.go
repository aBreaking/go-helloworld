package main

import (
	"fmt"
	"io"
)

/**
关键字 defer 允许我们推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数（为什么要在返回之后才执行这些语句？因为 return 语句同样可以包含一些操作，而不是单纯地返回某个值）。

关键字 defer 的用法类似于面向对象编程语言 Java 和 C# 的 finally 语句块，它一般用于释放某些已分配的资源。

一些使用场景：比如关闭数据源链接、文件io关闭、实现代码执行追踪等等

from : https://fuckcloudnative.io/the-way-to-go/06.4.html
*/
func main() {
	n, err := d("3")
	fmt.Println(n)
	fmt.Println(err)
}

func function1() {
	defer function2()
	fmt.Printf("In function1 at the top\n")
	fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Printf("Function2: Deferred until the end of the calling function!")
}

/**
使用 defer 的语句同样可以接受参数，下面这个例子就会在执行 defer 语句时打印 0：
*/
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

/**
当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）：
defer放在循环里是不推荐这样做的
*/
func b() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

/**
一个基础但十分实用的实现代码执行追踪的方案就是在进入和离开某个函数打印相关的消息，即可以提炼为
*/
func x() {
	fmt.Println("enter x ")
	defer fmt.Println("leave x")

	fmt.Println("in x")
}

func y() {
	fmt.Println("enter y")
	defer fmt.Println("leave y")
	fmt.Println("in y")
	x()
}

/**
使用 defer 语句来记录函数的参数与返回值
*/
func d(s string) (n int, err error) {
	defer func() {
		/**
		n, err 都是在函数返回之后拿到的结果
		*/
		fmt.Printf("func1(%q) = %d, %v", s, n, err)

	}()
	return 7, io.EOF
}
