package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("*.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	err = tmpl.ExecuteTemplate(w, "index.tmpl", nil)
	if err != nil {
		fmt.Println("render template failed, err:", err)
		return
	}
}

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe("0.0.0.0:9099", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}
