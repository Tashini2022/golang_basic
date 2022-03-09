package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t, err := template.ParseFiles("./templ/base.tmpl", "./templ/index.tmpl")
	if err != nil {
		fmt.Printf("paese template failed, error:%v\n", err)
		return
	}
	msg := "它是你"
	// 渲染模板
	t.ExecuteTemplate(w, "index.tmpl", msg)
}

func main() {
	http.HandleFunc("/index", index)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("http serve failed, error:%v\n", err)
		return
	}

}
