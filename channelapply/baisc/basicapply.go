package baisc

import (
	"fmt"
	"time"
)

/**
	协程以 go 【滚】 来声明的
 */
func Demo() {
	c1 := make(chan int)

	go func() {
		time.Sleep(time.Second*1)
		fmt.Println("管道赋值了")
		c1 <- 2
		//此处会进行阻塞，知道管道ch中的数据被接收后才能正确执行
		fmt.Println("管道已经赋值了")
	}()

	fmt.Println("等待中")
	time.Sleep(time.Second*5)
	//会进行阻塞，直到ch中的数据被接收
	fmt.Println("data is :",<-c1)
}