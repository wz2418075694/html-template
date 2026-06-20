package main

import (
	"html/template"
	"log"
	"net/http"
)

// 在Go的html/template包中，默认会对所有输出内容进行HTML转义，防止XSS攻击。
// xss跨站脚本攻击
// 如图safe函数就是不进行转移，标志为安全的数据
func xss(w http.ResponseWriter, r *http.Request) {
	//创建模板
	//加载模板
	//自定义函数要在加载模板前面进行定义，返回的template.HTML()是安全的意思，不需要转义
	tmpl, err := template.New("xss.html").Funcs(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	}).ParseFiles("./xss.html")
	if err != nil {
		log.Println("加载模板失败:", err)
		return
	}
	//渲染模板
	str1 := "<script>alert('你已经暴露了！')</script>"
	str2 := "<a href='https://www.baidu.com'>百度地址</a>"
	err = tmpl.Execute(w, map[string]string{
		"s1": str1,
		"s2": str2,
	})
	if err != nil {
		log.Println("渲染模板失败:", err)
		return
	}

}

func main() {
	http.HandleFunc("/xss", xss)
	log.Println("启动服务器")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("监听端口失败:", err)
		return
	}
}
