package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	t := template.New("hello.tmpl")
	// 自定义一个夸人的模板函数
	kua := func(arg string) (string, error) {
		return arg + "真帅", nil
	}
	t.Funcs(template.FuncMap{
		"kua": kua,
	})
	_, err := t.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	name := "小王子"
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", f1)
	err := http.ListenAndServe("0.0.0.0:9099", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}
