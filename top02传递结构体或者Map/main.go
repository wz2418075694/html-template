package main

import (
	"html/template"
	"log"
	"net/http"
)

type Person struct {
	Name   string
	Age    int
	Gender string
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	tmpl, err := template.ParseFiles("./hello.html")
	if err != nil {
		log.Println("解析模板失败：", err)
		return
	}
	wangzhao := Person{
		Name:   "王召",
		Age:    66,
		Gender: "男",
	}

	lisi := map[string]interface{}{
		"Name":   "李四",
		"Age":    20,
		"Gender": "女",
	}

	//切片
	hobbywangzhao := []string{
		"抽烟",
		"喝酒",
		"烫头",
	}

	//渲染模板
	err2 := tmpl.Execute(w, map[string]interface{}{
		"p1":    wangzhao,
		"p2":    lisi,
		"hobby": hobbywangzhao,
	})
	if err2 != nil {
		log.Println("渲染模板失败：", err2)
	}
}

// 写一个httpServer.
func main() {
	http.HandleFunc("/Hello", sayHello)
	//监听地址
	log.Println("启动服务器！")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Println("监听接口失败:", err)
		return
	}
}
