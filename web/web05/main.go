package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	hi := func(name string) (string, error) {
		// 需要在解析模板文件之前传入模板引擎
		return name + "好久不见！", nil
	}

	// 1. 定义模板文件
	t := template.New("xs.tmpl")

	// 在解析之前告诉模板引擎，现在多了一个叫hi的函数
	t.Funcs(template.FuncMap{
		"hi": hi,
	})

	// 2. 解析模板
	_, err := t.ParseFiles("./xs.tmpl")

	// 链式表达：t, err := template.New("xs.tmpl").ParseFiles("./xs.tmpl")
	if err != nil {
		fmt.Printf("create template failed, error:%v\n", err)
		return
	}

	// 3. 渲染模板
	name := "它是你"
	t.Execute(w, name)
}

func templ(wr http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, error:%v\n", err)
		return
	}
	name := "它是你"
	// 渲染模板
	t.Execute(wr, name)

}

func main() {
	http.HandleFunc("/hi", sayHello)
	http.HandleFunc("/tmpl", templ)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http serve failed, error:%v\n", err)
		return
	}
}
