package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.New("index.tmpl").
		Delims("{[", "]}"). // 修改标识符
		ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, error:%v\n", err)
		return
	}
	str := "它是你"
	// 渲染模板
	t.Execute(w, str)
}

func xss(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.New("xss.tmpl"). // 自定义函数
						Funcs(template.FuncMap{
			"safe": func(s string) template.HTML {
				return template.HTML(s)
			},
		}).
		ParseFiles("./xss.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, error:%v\n", err)
		return
	}
	str1 := "<script>alert(123)</script>"
	str2 := "<a href='http://liwenzhou.com'>学习golang!</a>"
	// 渲染模板
	t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("http serv failed, error:%v\n", err)
		return
	}
}
