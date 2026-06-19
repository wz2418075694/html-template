package main

import (
	"html/template"
	"log"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

func hello1(w http.ResponseWriter, r *http.Request) {
	//使用继承创建模板
	//加载模板
	tmpl, err := template.ParseFiles("./base.html", "./hello1.html")
	if err != nil {
		log.Println("加载模板失败：", err)
		return
	}
	//渲染模板
	wangzhao := Person{
		Name: "王召",
		Age:  18,
	}
	err = tmpl.ExecuteTemplate(w, "hello1.html", wangzhao)
	if err != nil {
		log.Println("渲染模板失败：", err)
		return
	}

}

func hello2(w http.ResponseWriter, r *http.Request) {
	//使用继承创建模板
	//加载模板
	tmpl, err := template.ParseFiles("./base.html", "./hello2.html")
	if err != nil {
		log.Println("加载模板失败：", err)
		return
	}
	//渲染模板
	zhangsan := Person{
		Name: "张三",
		Age:  20,
	}
	//因为加载了两个模板，所以要指定具体渲染的模板
	err = tmpl.ExecuteTemplate(w, "hello2.html", zhangsan)
	if err != nil {
		log.Println("渲染模板失败：", err)
		return
	}

}
func main() {
	http.HandleFunc("/hello1", hello1)
	http.HandleFunc("/hello2", hello2)
	log.Println("启动服务器")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("监听端口失败：", err)
		return
	}
}
