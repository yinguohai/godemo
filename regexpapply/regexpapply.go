package regexpapply

import (
	"fmt"
	"regexp"
)

/**
	func Match(pattern string, b []byte) (matched bool, err error)
	功能：
		检测b中是否存在匹配pattern的子序列。存在则返回true , 不存在则返回false
 */
func Match() {
	str := "abc"

	//ok ---bool , 成功： true , 失败: false
	//err --- 匹配过程中出现错误，就抛错误信息，没有出错就是nil
	//        匹配不到不代表错误
	ok , err := regexp.Match("[a-z{3}]",[]byte(str))
	fmt.Println(ok,"***********",err)
}


/**
	func MatchString(pattern string, s string) (matched bool, err error)
	功能：
		MatchString 类似于Match , 但是匹配的对象是字符串
 */
func MatchString() {
	str := "hello world"

	ok , err := regexp.MatchString("[a-z]{3}",str)

	fmt.Println(ok ,"#######",err)
}


/**
	func Compile(expr string) (*Regexp , error)
	功能：
		Compile 解析并返回一个正则表达式。 如果成功返回改Regexp就可用于匹配文本
		即：实例一个 regexp对象，并且让其设置好pattern


	注意：
		在匹配文本时该正则表达式选择的是非贪婪匹配模式进行匹配的。
		一般都是会和Match， MatchString 等其它函数进行配合使用
 */
func Compile() {
	r , _ := regexp.Compile("[a-z]{4}")

	str := "hello world"

	ok := r.MatchString(str)

	fmt.Println(ok)
}

/**
	func MustCompile(str string) *Regexp
	功能:
		MustCompile 和 Compile 类似，只是它只有一个返回值，而且匹配不成功会直接panic ,这个就是为了安全初始化
 */
 func MustCompile(){
 	r := regexp.MustCompile("[a-z]{4}")

 	str := "hello world"

 	ok := r.MatchString(str)

 	fmt.Println(ok)
 }


/**
	func (re *Regexp) Find(b []byte) []byte
	功能:
		返回保管正则表达式re在b中的最左侧的一个匹配结果的[]byte 切片。 如果没有匹配到 ，返回nil
 */
func Find() {
	str := "hello world!"

	r , _ := regexp.Compile("[a-z]{4}")


	ok := r.Find([]byte(str))

	fmt.Println(ok)  //[103 101 108 108]

	fmt.Println(string(ok)) // hell
}

/**
	func (re *Regexp) FindString(s string) string
	功能：
		跟Find类似，只是他的查找被查找对象和返回值都是string类型，而且如果没有匹配到返回的也是""
 */
func FindString(){
	str := "hello world"

	str2 := "123456"

	r , _ := regexp.Compile("[a-z]{4}")

	ok := r.FindString(str)

	ok2 := r.FindString(str2)

	fmt.Println(ok)  // hell

	fmt.Println(ok2)  // ""
}

/**
	func (re *Regexp) FindStringIndex(s string) (loc []int)
	功能：
		Find 返回保管正则表达式re在b中"最左侧"的一个匹配结果的起始位置下标的切片
 */
 func FindStringIndex(){
 	str := "hello world"

 	r , _ := regexp.Compile("[a-z]{4}")

 	ok := r.FindStringIndex(str)

 	fmt.Println(ok)
 }

/**
	func (re *Regexp) FindSubmatch(b []byte) [][]byte
	功能：
		Find 返回一个保管正则表达式re在b中最左侧的一个匹配结果以及"分组"匹配的结果的[][]byte切片，没有匹配到返回nil
	注意：
		Submatch 指的是正则表达式中具有分组的时候会起作用
 */

 func FindStringSubmatch(){
 	str := "peach punch"
 	r, _ := regexp.Compile("p([a-z]+)ch")

 	fmt.Println(r.FindStringSubmatch(str))

 }

/**
	func (re *Regexp) FindAllString(s string, n int) []string
	功能：
		Find 返回保管正则表达式re在s中所有不重叠的匹配结果[]string切片。如果没有匹配到，会返回nil
		n ： 表示的结果个数

 */

 func FindAllString() {
 	str := "peach punch world pinch hello "
 	r , _ := regexp.Compile("p([a-z]+)ch")

 	fmt.Println(r.FindAllString(str,5))
 }

/**
	func (re *Regexp) Split(s string, n int) []string
	功能：
		切割字符串，切割符是re中匹配到的结果，返回一个[]string切片。
 */
 func Split(){
 	str := "abaacaadaaaeaaaafaaaaagaaaaaae"

 	r := regexp.MustCompile("a+")
	//最多拆分成5部分
 	result := r.Split(str,5)

 	fmt.Println(result,len(result))
 }

 /**
 	func (re *Regexp) Expand(dst []byte, template []byte, src []byte, match []int) []byte
 	功能：
 		Expand 将匹配模板所匹配部分叠加至dst 尾部并返回
  */
  func Expand() {
  	reg := regexp.MustCompile(`(\w+),(\w+)`)

  	//源文本
  	src := []byte("Golang,World")

  	//目标文本
  	desc := []byte("Say:")

  	//模板
  	tmplate := []byte("hello $1 , Hello $2")

  	//解析源文本
  	match := reg.FindSubmatchIndex(src)

  	result := reg.Expand(desc,tmplate,src,match)

  	fmt.Println("match ....", match)

  	fmt.Println("expand...",string(result))

  }

  /**
  	func (re *Regexp) ExpandString(dst []byte, template string, src string, match []int) []byte
  	功能：
  		将template 处理后的结果追加到 dst的尾部
  	注意：
  		1.  template 中要有 $1 , $2 , ${name1} , {$name2} 这种分组引用符
  		2.  match 是由FindSubmatchIndex 方法返回的结果，存储的是各分组的位置信息
  		3.	src中取出相应的子串，替换掉template 中的 $1 , $2 等引用符号
  		4.  如果template中有“分组引用符”，则以match为标准
   */
  func ExpandString(){
  	reg := regexp.MustCompile(`(\w+),(\w+)`)
  	src := "Golang,World"
  	desc := []byte("Say:")

  	tmplate := "hello $1 , Hello $2"

  	match := reg.FindStringSubmatchIndex(src)

  	result := reg.ExpandString(desc,tmplate,src,match)
	// [0 12 0 6 7 12]
	// 0 12 是匹配的整个结果， 0~6 是第一个(\w+)匹配的结果， 7~12 是第二个(\w+)匹配的结果
  	fmt.Println(match)

  	// Say:hello Golang , Hello World
  	fmt.Println(string(result))

  }

/**
	func (re *Regexp) ReplaceAllString(src, repl string) string
	功能：
		在src中搜索匹配项， 并替换为repl指定的内容，返回替换后的结果
 */

 func ReplaceAllString(){
 	src := "Hello World,123 Go!!!"
 	reg := regexp.MustCompile(`Hello`)

 	template := "ooo"

 	result := reg.ReplaceAllString(src,template)

 	fmt.Println(result)
 }