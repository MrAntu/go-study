package main

import (
	"bookstore0612/controller"
	"net/http"
	"text/template"
)

// IndexHandler 去首页
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t := template.Must(template.ParseFiles("views/index.html"))
	//执行
	t.Execute(w, "")
}

func main() {
	//设置处理静态资源，如css和js文件
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	//直接去html页面
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	http.HandleFunc("/main", IndexHandler)

	//去登录
	http.HandleFunc("/login", controller.Login)
	//去注册
	http.HandleFunc("/regist", controller.Regist)
	http.ListenAndServe(":8080", nil)
}
