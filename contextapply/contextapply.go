package contextapply

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func Web() {
	go http.ListenAndServe(":8080",nil)
	//创建一个context具有超时时间属性的的上下文对象，
	//这种context的属性还可以有其它的表现形式
	// context.WithCancel(parent Context)
	//context.WithDeadline(parent Context,timeout time.Duration)
	//context.WithValue(parent Context , key interface{} , value interface{})
	ctx , _ := context.WithTimeout(context.Background(),(10*time.Second))
	go testA(ctx)
	select {}
}

func testA(ctx context.Context) {
	ctxA , _ := context.WithTimeout(ctx,(5*time.Second))
	ch := make(chan int)

	go testB(ctxA,ch)

	for {
		select {
		//如果父进程的ctx.Done 没有发送信号就不会强制接收本goroutine
		case <-ctx.Done():
			fmt.Println("父协程已经超时，需要结束本协程 ， testB 干干掉了")
			return
		case i := <-ch:
			fmt.Println("receiver :",i)
		default:
			time.Sleep(time.Second*1)
			fmt.Println("testA 监控中 ...")
		}
	}

}

func testB(ctx context.Context, ch chan int) {
	//模拟数据读取
	sumCh := make(chan int)
	go func(sumCh chan int) {
		sum := 10
		time.Sleep(3*time.Second)
		//10秒后发送一个信号
		sumCh <- sum
	}(sumCh)

	for {
		select {
		case <- ctx.Done():
			fmt.Println("父goroutine 已经超时， 需要结束本协程，  testB 干干掉了")
			<-sumCh
			return
		case i := <-sumCh:
			fmt.Println("send :",i)
			ch <- i
		default:
			time.Sleep(time.Second*1)
			fmt.Println("testB 监控中...")
		}
	}
}