package main

import (
	"fmt"
)

type User struct {
	name string
	age  int
}

func main() {
	ch := make(chan int)
	go myChan(5, ch)
	go myChan(6, ch)
	x, y := <-ch, <-ch
	fmt.Println(x + y)
}
func myChan(s int, ch chan int) {
	ch <- s
}

func goMap() {

	myMap := make(map[string]int) //使用make的方式来初始化一个空map
	myMap["zhangsan"] = 18
	myMap["lisi"] = 20
	fmt.Println(myMap["zhangsan"]) //18

	delete(myMap, "zhangsan")
	fmt.Println(myMap["zhangsan"]) //0
}

func goSlice() {
	myArray := []int{1, 2, 3, 4, 5, 6, 7}
	mySlice1 := myArray[2:5]                            //表示截取从下标 2 到下标5的元素，即[2,5)
	fmt.Println(mySlice1, len(mySlice1), cap(mySlice1)) // [3 4 5] 3 5

	var mySlice2 = make([]int, 10)                      //可用用make来进行声明切片
	mySlice2 = myArray[:5]                              //表示截取从开头 到5 的元素
	fmt.Println(mySlice2, len(mySlice2), cap(mySlice2)) //[1 2 3 4 5] 5 7

	mySlice3 := myArray[2:]                             //表示接却从2 到结尾的元素
	fmt.Println(mySlice3, len(mySlice3), cap(mySlice3)) //[3 4 5 6 7] 5 5

	array1 := []int{1, 2, 3, 4, 5}
	array1 = append(array1, 6, 7, 8)
	fmt.Println(array1) // [1 2 3 4 5 6 7 8]

	array2 := []int{10, 11, 12}
	copy(array1, array2) // 把array2的元素拷贝到array1里面
	fmt.Println(array1)  //[10 11 12 4 5 6 7 8]
}

func goStruct() {
	zhangsan := User{"zhangsan", 18}
	lisi := User{name: "lisi"}
	lisi.age = 19
	zhangsan.name = "zhangsan"
}

func complexMethod(a, b int) (map[int]string, string, error) {
	return nil, "", nil
}

func simpleMethod(a int) map[int]string {
	return nil
}

func goFor() {
	myArray := [3]int{1, 2, 3}
	for index, value := range myArray {
		fmt.Println(index, value)
	}

	for _, value := range myArray {
		fmt.Println(value)
	}

	for i := 0; i < 10; i++ {

	}

	for true {
	}
}
