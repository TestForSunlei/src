package websever

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
* @program: src
*
* @description:一个简单的web客户端
*
* @author: sunlei
*
* @create: 2018-11-12 17:24
**/

func HttpClientFUNC() {
	res, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("http client err:", err)
		return
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ReadAll err :", err)
		return
	}
	fmt.Println(string(data))

}
