package main

import (
	"encoding/json"
	"fmt"
)

/**
struct json tag主要在struct与json数据转换的过程(Marshal/Unmarshal)中使用。

json的tag格式如下：

Key type  `json:"name,opt1,opt2,opts..."`
1
说明：

变量必须是可导出的（Key首字母必须大写），否则会被忽略处理。
没有json tag或者tag中name省略（但不能少了","），默认使用字段名。
name要注意命名的有效性。
opt1、opt2等项为可选项，必须使用有限的几个限定的opt的一个或组合，如"omitempty"、"string"，使用非限定的opt会发生错误。

*/
func main() {
	s := Student{
		StudentId:      1,
		StudentName:    "zhangsan",
		StudentTeacher: "lisi",
	}
	jsonStr, _ := json.Marshal(s)
	fmt.Println(string(jsonStr))

	so := StudentWithOption{
		s,
		Class{"7class"},
		"3年2班",
	}
	soJson, _ := json.Marshal(so)
	fmt.Println(string(soJson))
}

type Student struct {
	StudentId      int    `json:"sid"`
	StudentName    string `json:"sname"`
	StudentTeacher string `json:"-"` //表示转为json时忽略该字段
	S2             string //默认就以S2作为key
}

type Class struct {
	Name string `json: className`
}

type StudentWithOption struct {
	Student      `json:"student"` //student作为key，里面内容作为json {student:{},.... }
	Class        `json:",inner"`  // 直接把class里的内容作为key value { className:"" }
	StudentClass string           `json:"class,omitempty"` // 该字段没有值才会忽略该key
}
