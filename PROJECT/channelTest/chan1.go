package channelTest

import (
	"fmt"
	"time"
)

/**
 * @Author sunlei
 * @Description 写入channel
 * @Date 11:13 2018/11/12
 * @Param ch chan int
 * @return void
 **/
func write(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

/**
 * @Author sunlei
 * @Description 读入channel数据
 * @Date 11:12 2018/11/12
 * @Param ch chan int
 * @return void
 **/
func read(ch chan int) {
	for {
		var b int
		b = <-ch
		fmt.Println(b)
	}
}

func MainFUNC1() {
	intChan := make(chan int)
	go write(intChan)
	go read(intChan)

	time.Sleep(1 * time.Second)

}
