package http

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

func f1(w http.ResponseWriter, r *http.Request) { //处理函数
	b, err := os.ReadFile(`F:\1\a.txt`)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(b)
}

func getHandler(w http.ResponseWriter, r *http.Request) { // GET请求响应函数
	defer r.Body.Close()
	data := r.URL.Query()
	fmt.Println(data.Get("name"))
	fmt.Println(data.Get("age"))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}
func postHandler(w http.ResponseWriter, r *http.Request) { // POST请求响应函数
	defer r.Body.Close()
	// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
	r.ParseForm()
	fmt.Println(r.PostForm) // 打印form数据
	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))
	// 2. 请求类型是application/json时从r.Body读取数据
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read request.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Hello 沙河11！")
	if err != nil {
		println(err)
		return
	}
	print("1231")
}

func Test_httpServer(*testing.T) {
	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}

	//s := &http.Server{  // 自定义Server
	//	Addr:           ":8080",
	//	Handler:        myHandler,
	//	ReadTimeout:    10 * time.Second,
	//	WriteTimeout:   10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//log.Fatal(s.ListenAndServe())
}
