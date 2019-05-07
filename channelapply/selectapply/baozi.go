package selectapply

import (
	"fmt"
	"time"
)

/**
	生产者,共10个人,一个人包一个包子
 */
func producer(c chan string) {
	for i:=0;i<10;i++ {
		go func(c chan string,i int) {
			msg := fmt.Sprintf("生产的包子编号：%v",i)
			time.Sleep(time.Second*5) //TODO
			fmt.Println(msg)
			//往通道中写数据，需要放到最后面。防止main 提取结束
			//这是为什么呢？^_^~~~
			c <- msg
		}(c,i)
	}
}

/**
	消费者，共10个人，一个消费一个包子
 */
func customer(baozi string,log chan bool) {
	fmt.Println( "*买了包子*",baozi)
	log <- true
}

/**
	卖包子
 */
func Shop() {
	//流水线3个
	wokerLine := make(chan string ,3)

	//消费记录
	logChan := make(chan bool,2)
	//统计数据
	total := 0
	//开启生产包子
	producer(wokerLine)
	p("生产任务分配完了，坐等包子吃")
	//使用select监听是否产生了包子，有则消费
	for {
		select {
		case b := <-wokerLine:
			//每生产了一个包子则立马卖给一个顾客
			go customer(b,logChan)
		case <-logChan:
			total += 1
			if total == 10 {
				fmt.Println("今天的包子卖完了，明天再来吧")
				return
				//break
			}
		}
		//select {}
	}
}