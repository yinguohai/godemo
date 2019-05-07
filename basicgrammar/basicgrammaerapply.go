package basicgrammar

import "fmt"

/**
	golang 中 声明变量的方式有两种
	a. 初始化变量
	b. 短声明变量

 */
func InitDemo() {
	//最基础的声明方式
	var x int
	x = 32
	fmt.Println(x)

	//联合声明,一次可以声明多个不同类型的变量
	var  (
		y , z string
		q int
	)
	y = "nan"
	z = "boy"
	q = 20
	fmt.Println(y ,z,q)

	//简短声明, 注意，在同一个作用域类，被var 声明过的变量不能再使用简短声明
	bb := "hello"
	fmt.Println(bb)
	//简短声明特殊：示例
	if x >0 {
		//此处的y使用的则是if语句块中的y ，而不再是19行var中声明的y了
		y := "99"
		fmt.Println("y in if :",y)
	}
	fmt.Println("y in func",y)
}
