package main

import (
	"html/template"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {

	//自定义一个功能函数
	hello := func(name string) (string, error) {
		return name + "欢迎你！", nil
	}

	//定义模板
	//利用New()创建了一个名字叫f.html的模板对象
	tmpl := template.New("f.html")
	//定义的函数必须在解析模板之前，不然tmpl.ParseFiles()找不到定义的函数
	//Html中调用helloYou，就是执行go里面的hello函数，将他们两个进行映射
	tmpl.Funcs(template.FuncMap{
		"helloYou": hello,
	})

	//解析模板
	_, err := tmpl.ParseFiles("./f.html")
	if err != nil {
		log.Println("解析模板对象失败:", err)
		return
	}
	//渲染模板
	name := "王召"
	err = tmpl.Execute(w, name)
	if err != nil {
		log.Println("渲染模板失败：", err)
	}
}
func main() {
	http.HandleFunc("/hello", hello)
	log.Println("启动服务器！")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("监听端口失败：", err)
		return
	}
}
