package channelTest

import (
	"fmt"
	"time"
)

/**
* @program: src
*
* @description:channel通信的一个例子
*
* @author: sunlei
*
* @create: 2018-11-12 10:50
**/

/**
 * @Author sunlei
 * @Description 判断channel中的数是不是素数,是的话放入管道
 * @Date 11:35 2018/11/12
 * @Param
 * @return
 **/
func calc(taskChan chan int, output chan int) {
	for v := range taskChan {
		flag := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			output <- v
		}
	}
}

func MainFUNC2() {
	var input chan int = make(chan int, 100)
	var output chan int = make(chan int, 100)

	//因为input创建时缓冲区就只有100,不能直接插入大于100的数据,否则会阻塞
	for i := 0; i < 100; i++ {
		input <- i
	}
	close(input)

	for i := 0; i < 8; i++ {
		go calc(input, output)
	}

	//一个协程运行的立执行函数
	go func() {
		for v := range output {
			fmt.Println(v)
		}
	}()

	/*这种写法可以用ok来测试管道是否已经关闭
	for{
		b,ok <-chanInt
		if ok{
			break
		}
		fmt.Println(b)
	}
	*/
	time.Sleep(10 * time.Second)
}
