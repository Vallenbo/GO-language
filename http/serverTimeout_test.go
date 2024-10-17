package http

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) { // server端，随机出现慢响应
	number := rand.Intn(2)
	if number == 0 {
		time.Sleep(time.Second * 10) // 耗时10秒的慢响应
		fmt.Fprintf(w, "slow response")
		return
	}
	fmt.Fprint(w, "quick response")
}

func Test_serverTimeout(*testing.T) {
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
