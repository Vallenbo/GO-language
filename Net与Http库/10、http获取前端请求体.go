package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) { // 1、获取请求
	fmt.Fprintln(w, "接收的请求地址是：", r.URL.Path)
	fmt.Fprintln(w, "接收的请求地址后的查询字符是：", r.URL.RawQuery)
	fmt.Fprintln(w, "请求头中的所有信息有：", r.Header)
	fmt.Fprintln(w, "请求头中Accept-Encoding信息是：", r.Header["Accept-Encoding"])
	lens := r.ContentLength //获取请求体内容长度
	body := make([]byte, lens)
	r.Body.Read(body) //将请求体读到Body中
	fmt.Fprintln(w, "请求体中的内容有：", string(body))
	r.ParseForm()
	fmt.Fprintln(w, "表单请求参数有：", r.Form)
	//fmt.Fprintln(w, "请求参数为：", r.PostForm) //PostForm 包含来自 PATCH、POST 或 PUT 正文参数的解析表单数据
	//fmt.Fprintln(w, "请求参数 username 的值为：", r.PostFormValue("username"))

	r.FormValue("user")
	fmt.Fprintln(w, "请求参数 username 的值为：", r.PostFormValue("username"))

	r.ParseMultipartForm(1024) // 将请求正文解析为多部分/表单数据。
}
func main() {
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":8080", nil)
}

func handler1(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024) //解析表单
	fileHeader := r.MultipartForm.File["photo"][0]
	file, err := fileHeader.Open()
	if err == nil {
		data, err := io.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}

}
