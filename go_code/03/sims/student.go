package main

import "fmt"

//	student 构造体
type student struct {
	id    int //	学号唯一
	class string
	name  string
}

//	newStudent为student的构造函数
func newStudent(id int, class, name string) *student {
	return &student{
		id:    id,
		class: class,
		name:  name,
	}
}

//	studentMgr构造体
type studentMgr struct {
	allStudents []*student
}

//	newStudentMgr 为 studentMgr的构造函数
func newStudentMgr() *studentMgr {
	return &studentMgr{
		allStudents: make([]*student, 0, 100),
	}
}

//添加学生信息
func (s *studentMgr) addStudent(newStu *student) {
	s.allStudents = append(s.allStudents, newStu)
}

//编辑学生信息
func (s *studentMgr) modefyStudent(newStu *student) {
	for k, v := range s.allStudents {
		if newStu.id == v.id { //	依据学号给对应的学生修改信息
			s.allStudents[k] = newStu
			return
		}
		fmt.Printf("目前信息中没有学号为%d的学生", newStu.id)
	}
}

//展示学生信息
func (s *studentMgr) showStudent() {
	for _, v := range s.allStudents {
		fmt.Printf("学号: %d 姓名: %s 班级: %s\n", v.id, v.name, v.class)
	}
}
