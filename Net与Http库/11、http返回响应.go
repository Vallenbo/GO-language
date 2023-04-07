package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

type User struct {
	ID       int
	Username string
	Password string
}

func main() {
	http.HandleFunc("/testTemplate", handler)

	http.ListenAndServe(":8080", nil)
}

func handler2(w http.ResponseWriter, r *http.Request) { // 1、客户端响应 html格式
	w.Write([]byte("你的请求我已经收到"))
}

func handler3(w http.ResponseWriter, r *http.Request) { //1-1 客户端响应 json格式
	w.Header().Set("Content-Type", "application/json") //设置响应头中内容的类型
	user := User{
		ID:       1,
		Username: "admin",
		Password: "123456",
	}

	json, _ := json.Marshal(user) //将 user 转换为 json 格式
	w.Write(json)
}

func handler4(w http.ResponseWriter, r *http.Request) { //2、客户端重定向
	w.Header().Set("Location", "https:www.baidu.com") //以下操作必须要在 WriteHeader 之前进行
	w.WriteHeader(302)
}

func handler5(w http.ResponseWriter, r *http.Request) { //3、 客户端响应 返回template模板文件
	t, _ := template.ParseFiles("statics/hello.html") //解析模板文件
	//t, _ := template.ParseGlob("*.html") //通过该函数可以通过指定一个规则一次性传入多个模板文件
	//t := template.Must(template.ParseFiles("hello.html")) //解析模板时对错误进行处理

	t.Execute(w, "Hello World!") //执行模板
	//t.ExecuteTemplate(w, "statics/hello.html", "我要在 hello2.html 中显示")
}
