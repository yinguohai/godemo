package sort

import "time"

/**
	睡眠排序法：通过为待排序的元素启动独立的线程，每个线程根据元素的value执行相对应的睡眠时间，然后通过信道一一传入到信道收集
				效率比较低，而且对于相对比较接近的数据进行排序会存在一定的误差率
 */
func SleepSort() {
	num := []int {9,5,8,3,2,7,3,0,1}
	ch := make(chan int)
	for _, i := range num {
		go func(v int) {
			time.Sleep(time.Duration(v) * time.Second)
			ch <- v
		}(i)
	}

	for _ = range num {
		tmp := <-ch

		print(tmp)
	}
}