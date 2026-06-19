package main

import (
	"html/template"
	"log"
	"net/http"
)

// 1.已经定义好了模板，hello.tmpl
func helle(w http.ResponseWriter, r *http.Request) {

	//2.解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		log.Println("解析模板文件失败:", err)
		return
	}
	//3.渲染模板文件
	name := "王召在中国"
	//io.Writer,其实是接口的类型，代表任何能接受字节流的地方
	//http.ResponseWriter类型其实就是HTTP响应对象

	err = tmpl.Execute(w, name)
	if err != nil {
		log.Println("渲染模板失败：", err)
		return
	}

}
func main() {
	//写一个httpServer.
	http.HandleFunc("/", helle)
	log.Println("服务端启动！")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("服务器启动失败：", err)
		return
	}

}
