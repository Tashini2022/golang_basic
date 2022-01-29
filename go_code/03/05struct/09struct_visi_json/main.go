package main

import (
	"encoding/json"
	"fmt"
)

//	结构体字段的可见性与JSON序列化

//	如果一个go语言包中定义的标识符是大写，则表示其对外可见，小写则表示对外不可见
//	如果结构体内的字段首字母是小写，则其对外不可见，例如student 中的age字段，json序列化无法访问

type student struct {
	ID   int
	Name string
	age  int
}
type Class struct {
	Title    string    `json:"title"` //	为了能在序列化中显示小写字母所使用的tag，键值对用反引号包裹，不同键值对之间用空格分割
	Students []student `json:"student_list" db:"student" xml:"ss"`
}

//	student的构造函数
func newStudent(id int, name string) student {
	return student{
		ID:   id,
		Name: name,
	}
}

func main() {
	//	创建一个有30个作为的班级
	c1 := Class{
		Title:    "12121班",
		Students: make([]student, 0, 30),
	}

	//	构造10个学生,并传入班级c1
	for i := 1; i <= 10; i++ {
		stu := newStudent(i, fmt.Sprintf("stu%02d", i))
		c1.Students = append(c1.Students, stu)
	}
	fmt.Printf("c1=%#v ,type:%T\n", c1, c1)

	//	JSON序列化：GO语言中的数据==>JSON格式的字符串
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("a   json Marshal failed, err:", err)
		return
	}
	fmt.Printf("type: %T\n", data)
	fmt.Printf("%s\n", data)

	//	JSON反序列化：SON格式的字符串==>JGO语言中的数据
	jstr := `{"Title":"12121班","Students":[{"ID":1,"Name":"它是你"},{"ID":2,"Name":"stu02"},{"ID":3,"Name":"stu03"}]}` //	``反引号
	var c2 Class
	json.Unmarshal([]byte(jstr), &c2) //	注意使用字节类型数据，以及后面用内存地址接受数据
	fmt.Printf("type: %T\n", c2)
	fmt.Printf("%#v\n", c2)
}
