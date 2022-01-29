package main

import "fmt"

type Animal struct {
	name string
}
type Dog struct {
	feet    string
	*Animal //	匿名嵌套(嵌套类型为指针)
}

func (a *Animal) move() { //	接受者为Animal指针的方法
	fmt.Printf("%s会自己跑～\n", a.name)
}
func (d Dog) wang() {
	fmt.Printf("%s会汪汪汪～\n", d.name)
}
func main() {
	d1 := Dog{
		feet: "4条",
		Animal: &Animal{
			name: "旺财",
		},
	}
	d1.move()
	d1.wang()
}
