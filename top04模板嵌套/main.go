package main

import (
	"html/template"
	"log"
	"net/http"
)

// 模板嵌套
func demo(w http.ResponseWriter, r *http.Request) {
	//定义模板

	//解析模板
	//模板书写的顺序是:被嵌套的模板应该书写在后面，由于son.html被嵌套，所以它书写在后面
	tmpl, err := template.ParseFiles("./father.html", "./son.html")
	if err != nil {
		log.Println("解析模板失败：", err)
		return
	}
	//渲染模板
	name := "张三"
	err = tmpl.Execute(w, name)
	if err != nil {
		log.Println("渲染模板失败：", err)
		return
	}
}

func main() {
	http.HandleFunc("/demo", demo)
	log.Println("启动服务器！")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("监听端口失败：", err)
		return
	}
}
