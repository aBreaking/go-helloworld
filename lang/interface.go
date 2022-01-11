package main

import "fmt"

func main() {
	x := new(C)
	x.c = 3
	doA(x)

}

type A interface {
	A1() int
	A2() int
}

type B struct {
	b int
}

func (b B) A1() int {
	return b.b * b.b
}

func (b B) A2() int {
	return b.b + b.b
}

type C struct {
	c int
}

func (c C) A1() int {
	return c.c + c.c
}

func (c C) A2() int {
	return c.c * c.c
}

func doA(a A) {
	fmt.Println(a.A1())
}
