package md5

import (
	"crypto/md5"
	"fmt"
)

func Md5(){
	//声明一个md5对象
	Md5obj := md5.New()
	//往md5对象中添加byte字节数组类型的目标字符
	Md5obj.Write([]byte("zhangsan"))
	//计算md5的值
	result := Md5obj.Sum([]byte(""))

	fmt.Printf("%v\n",result)
	//打印md5的值
	fmt.Printf("%xx",result)
}
