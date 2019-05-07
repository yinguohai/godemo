package selectapply

import (
	"fmt"
	"time"
)

func addX(c chan int,s chan string) {
	for i:=0;i<10;i++ {
		time.Sleep(time.Second)
		c <-i
	}
	s <- "a"
}

func addY(c chan int,s chan string) {
	for i:=0;i<20;i++ {
		time.Sleep(time.Second)
		c <-i
	}
	s <- "b"
}

func SelectBaisc() {
	c1 := make(chan int)
	c2 := make(chan int)

	sigle := make(chan string,2)

	go addX(c1 , sigle)

	go addY(c2,sigle)

	var nums []string

	for {
		select {
		case t1 := <-c1:
			fmt.Println("addX :", t1)
		case t2 := <-c2:
			fmt.Println("addY :", t2)
		case t3 := <-sigle:
			p("接收到了信号")
			nums = append(nums,t3)
			//下面这种写法，nums会进行重新赋值，而不会累加
			//nums := append(nums,t3)
			p("nums :",nums,"****length****",len(nums))
			if len(nums) == 2{
				goto TAG
			}
		}
	}

	TAG:return

}

func CombineChannel(ch1,ch2 chan int) <-chan int {
	out := make(chan int,3)
	go func() {
		defer close(out)
		for {
			select {
				case v1 , ok := <-ch1:
					if !ok {
						ch1 = nil
						continue
					}
					out <- v1
				case v2 , ok := <-ch2:
					if !ok {
						ch2 = nil
						continue
					}
					out <- v2

			}

			if ch1 == nil && ch2 == nil {
				break
			}
		}
	}()
	return out
}

/**
	断点函数
 */
func p(param ...interface{}) {
	fmt.Println(param...)
	time.Sleep(time.Second*2)
}