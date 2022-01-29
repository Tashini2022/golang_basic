package main

import (
	"fmt"
	"os"
)

//学院信息管理系统

//需求
//1.添加学员信息
//2.编辑学员信息
//3.展示所有学员信息
//4.退出系统

//操作流程
//1.给出系统菜单
func sysMenu() {
	fmt.Println("欢迎来到学员信息管理系统")
	fmt.Println("1.添加学员信息")
	fmt.Println("2.编辑学员信息")
	fmt.Println("3.展示所有学员信息")
	fmt.Println("4.退出系统")
}

//	获取用户输入的学生信息
func getInput() *student {
	var (
		id    int
		class string
		name  string
	)
	fmt.Println("请输入学生信息:")
	fmt.Print("请输入学生学号:")
	fmt.Scanf("%d\n", &id)
	fmt.Print("请输入学生姓名:")
	fmt.Scanf("%s\n", &name)
	fmt.Print("请输入学生班级:")
	fmt.Scanf("%s\n", &class)
	stu := newStudent(id, class, name)
	return stu
}

func main() {

	sm := newStudentMgr()
	for {
		sysMenu()

		//2.用户选择要执行的操作
		var input int
		fmt.Print("请输入您要操作的序号: ")
		fmt.Scanf("%d\n", &input) //	将控制台输入的数据赋值给input
		fmt.Println("当前用户选择的是: ", input)

		//3.执行用户操作
		switch input {
		case 1:
			//添加学生信息
			stu := getInput
			sm.addStudent(stu)
		case 2:
			//编辑学生信息
			stu := getInput
			sm.modefyStudent(stu)
		case 3:
			//展示学生信息
			sm.showStudent()
		case 4:
			//退出程序
			os.Exit(0)
		default:
			fmt.Println("err! 请输入正确选项")

		}
	}
}
