package main

import "fmt"

func main() {
	u := MyUser{"a", 10}

	z := Zhangsan{u, "101"}
	z.PrintName()
}

type MyUser struct {
	Name string
	age  int
}

func (u MyUser) PrintName() {
	fmt.Println(u.Name)
}

type Zhangsan struct {
	MyUser // 相当于继承了User

	address string
}
