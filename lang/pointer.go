package main

import "fmt"

func main() {
	user := User{"zhangsan", 10}
	changePointer(&user)
	fmt.Println(user)
}

func changePointer(user *User) {
	user.age = 20
	user.name = "lisi"
}

func changeValue(user User) {
	user.age = 20
	user.name = "lisi"
}
