package websever

import (
	"fmt"
	"net/http"
)

/**
* @program: src
*
* @description:一个最简单的web服务器程序
*
* @author: sunlei
*
* @create: 2018-11-12 17:10
**/

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")
	fmt.Fprint(w, "hello")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle login")
	fmt.Fprint(w, "login")
}

func SimpleWebServer() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		fmt.Println("http error")
	}
}
