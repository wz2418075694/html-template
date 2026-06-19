package main

import (
	"html/template"
	"log"
	"net/http"
)

func demo(w http.ResponseWriter, r *http.Request) {

	//定义模板

	//解析模板
	tmpl, err := template.New("delims.html").
		Delims("{[", "]}").
		ParseFiles("./delims.html")
	if err != nil {
		log.Println("渲染模板失败：", err)
		return
	}
	//渲染模板
	name := "王召"
	err = tmpl.Execute(w, name)
	if err != nil {
		log.Println("渲染模板失败：", err)
		return
	}

}
func main() {
	http.HandleFunc("/delims", demo)
	log.Println("启动服务器")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("监听端口失败：", err)
		return
	}
}
