package ioapply

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)
var (
	firstName , lastName string
	age int
)
/**
	读取
 */
func IoReader() {
	fmt.Println("Please enter your full name:")
	fmt.Scanln(&firstName,&lastName)

	fmt.Println(firstName,lastName)
}



/**
	func Scanln(a ...interface{}) (n int, err error)
	Scanln 用于扫描 os.Stdin 中的数据，并将数据以空格为分割符进行分割，然后填写到参数列表 a 中
	当扫描过程中遇到 '\n' 或者参数列表 a 被全部填写完毕， 则停止扫描
 */
func Scanln() {
	fmt.Println("Please enter your full name")
	//获取控制台输入的数据
	fmt.Scanln(&firstName,&lastName)
	fmt.Println(firstName,lastName)
}

/**
	func Scanf(format string, a ...interface{}) (n int, err error)
	Scanf 用于扫描 os.Stdin 中的数据，并根据 format 指定的格式 ，将扫描出的数据填写到参数列表 a 中
 */
func Scanf() {
	//	>#go run main.go
	//	name = huahua
	//输入的字符必须和format 的一抹一样
	fmt.Scanf("%s = %s",&firstName,&lastName)

	fmt.Println(firstName,"====>",lastName)
}

/**
	func Sscanf(str string, format string, a ...interface{}) (n int, err error)
	用于扫描字符串str中的数据，并根据format指定的格式将扫描出的数据填写到参数列表a中
 */
func Sscanf() {
	//提取“Golang”  和 “4” 这两个数据
	s := "我的名字叫 Golang ，今年 4 岁"
	fmt.Sscanf(s, "我的名字叫 %s ，今年 %d 岁", &firstName, &age)
	fmt.Printf("%s %d", firstName, age)
}
/**
	func Fscanf(r ioapply.Reader, format string, a ...interface{}) (n int, err error)
	用于扫描r中的数据，并根据format指定的格式将扫描的数据填写到参数列表 a 中 。当 r 中的数据被全部扫描完毕或扫描长度超出 format 指定的长度时
	则停止扫描（换行符会被当作空格处理）
*/
func Fscanf() {
	//声明要给io类型过来的字符串流
	s := strings.NewReader("我的名字叫Golang , 今年 5 岁了")
	fmt.Println(s)
	f := "我的名字叫%s , 今年 %d 岁了"
	fmt.Fscanf(s,f,&lastName,&age)
	fmt.Println(lastName,"===>",age)
}

/**
	golang 对命令行的解析提供了flag包，按照功能可以分为一下三类：
	1. 基本命令解析
	2. 自定义参数解析
	3. 解析外部命令字符串

	命令注册：
	flag.type(int,string等)，提供命令名称，默认参数值

	命令解析：
	flag.Parse()
 */
func FlagBool() {
	//基本的标记声明仅支持字符串，整数和布尔值选项
	//方式一，传入指针
	var block bool
	flag.BoolVar(&block,"b",false,"set for block")
	flag.Parse()
	fmt.Println("the block is :",block)

	//方式二: 传入变量
	var block2 =  flag.Bool("c",true,"set for block")
	flag.Parse()
	fmt.Println("the block2 of is :",*block2)
}

/**
	func Var(value Value, name string, usage string)
	自定义参数解析
	Var 方法使用指定的名字，使用信息注册一个flag. 该flag的类型和值由第一个参数表示，
	该参数应实现了Value接口。例如，用户可以创建一个flag,可以用Value接口的Set方法将逗号分隔的
	字符串转化为字符串切片
*/

type args []string
func (a *args) String() string {
	return fmt.Sprintf("%v",*a)
}
func (i *args) Set(value string) error {
	if len(*i)>0 {
		return errors.New("interval flag already set")
	}
	for _,dt := range strings.Split(value,","){
		*i = append(*i,dt)
	}
	return nil
}

func FlagVar() {
	var argFlag args
	flag.Var(&argFlag,"deltaT","comma-separated list of intervals to use between events")
	flag.Parse()
	fmt.Println("arg string",argFlag)
}

func FlagArgs() {
	var name string

	flag.StringVar(&name,"name","root","set name")

	flag.Parse()

	//===============上面的flag类型的参数必须定义==========================

	others := flag.Args()

	fmt.Println(name)
	fmt.Println(others)
}

