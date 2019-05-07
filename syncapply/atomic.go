package syncapply

import (
	"fmt"
	"sync"
	"sync/atomic"
)
/**
	原子增量
 */
func AtomicAdd() {
	var sum int32 = 0
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//atomic.AddUint32 提供了原子性增量操作
			atomic.AddInt32(&sum, -1)
		}()
	}

	wg.Wait()
	fmt.Println(sum)
}

/**
	CAS
	先比较变量的值是否等于给定旧值，等于旧值的情况下才赋予新值，最后返回新值是否设置成功。
 */
func AtomicCompareSwrap() {
	var sum uint32 = 99
	var wg sync.WaitGroup

	for i:=uint32(0);i<1000;i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//如果sum的值等于100，则sum的值加1
		atomic.CompareAndSwapUint32(&sum,100,sum+1)
		}()
	}

	wg.Wait()
	fmt.Println(sum)
}

/**
	原子的值存储操作总会成功，因为它并不关系其原来的值是什么
 */
func AtomicStore() {
	var sum uint32 = 0

	atomic.StoreUint32(&sum,300)

	fmt.Println(sum)
}

/**
	原子导出值
	当我们在读取value的过程中，其它的程序是有可能对value进行读|写操作的，
	所以我们读取的数据有可能只是更新了一半

 */
 func AtomicExport(){
	var v uint32 = 20
	//防止我们在读取v的时候，其它程序来对它进行更改操作
	atomic.LoadUint32(&v)
	fmt.Println(v)
 }

 /**
赋予变量新值，同时返回变量的旧值。
  */
func AtomicSwap() {
	var v uint32 = 40

	var v2 uint32 = 300
	fmt.Println("before"," ",v,"***",v2)

	old := atomic.SwapUint32(&v,v2)
	fmt.Println("after"," ",old,"#####",v,"******",v2)
}