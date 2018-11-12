package channelTest

import (
	"fmt"
	"time"
)

/**
* @program: src
*
* @description:对channnel进行select操作
*
* @author: sunlei
*
* @create: 2018-11-12 15:10
**/

func getInput(channel1 chan int, channel2 chan int, channel3 chan int, channel4 chan int) {
	for {
		select {
		case i := <-channel1:
			fmt.Println("channel1: ", i)
		case i := <-channel2:
			fmt.Println("channel2: ", i)
		case i := <-channel3:
			fmt.Println("channel3: ", i)
		case i := <-channel4:
			fmt.Println("channel4: ", i)
		}
	}
}

func MainFUNC3() {
	var channel1 chan int = make(chan int, 10)
	var channel2 chan int = make(chan int, 10)
	var channel3 chan int = make(chan int, 10)
	var channel4 chan int = make(chan int, 10)

	go getInput(channel1, channel2, channel3, channel4)

	for i := 0; i < 100; i++ {
		channel1 <- i
		channel2 <- 1000 + i
		channel3 <- 10000 + i
		channel4 <- 100000 + i
		fmt.Println("第", i, "次")
		//time.Sleep(time.Second)
	}
	time.Sleep(time.Second)
}
