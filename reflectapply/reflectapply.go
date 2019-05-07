package reflectapply

import (
	"fmt"
	"reflect"
)

//通过反射获取到传入变量的 type , kind , value
func ReflectTest(b interface{}){

	//1. 先获取到 reflectapply.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println(rTyp)

	//2. 获取到 reflectapply.Value
	rVal := reflect.ValueOf(b)
	//无法相加
	//result := rVal + 3
	//fmt.Println(rVal)
	//查看rVal的真正类型 ,可以看到  ==> reflectapply.Value
	fmt.Printf("rVal = %v , rVal of Type : %T \n",rVal,rVal)

	//我们不要 reflectapply.Value 类型的rVal我们就要我们传进来的值类型

	//ok,使用Go中的reflect.Value的一种反射方法 .Int()来解决
	result := rVal.Int() + 1

	fmt.Printf("result = %v\n",result)

	//将rVal转成Interface() 类型

	iV := rVal.Interface()

	//将 interface{} 通过断言，转成我们需要的类型

	num2 := iV.(int)

	fmt.Println("num2=",num2)
}


type student struct {
	Name string
	Age int
}

//演示对结构体的反射
func Reflect02 (b interface{}) {
	rVal := reflect.ValueOf(b).Interface()
	fmt.Printf("rVal = %v , rVal of Type = %T\n",rVal,rVal)


	//将 reflectapply.student 类型转换成 interface 类型，便于我们使用断言
	//iv := rVal.Interface()
	//
	//stu , ok := iv.(student)
	//
	//
	//if ok {
	//	fmt.Println("stu.Name = ",stu.Name,"stu.Age =",stu.Age)
	//} else {
	//	logapplay.Fatal("stu 类型判断错误~~~~~~ ",ok)
	//	fmt.Printf("stu type is : %T",iv)
	//}

}

//通过反射，值进行修改
func Reflect03 (b interface{}) {
	rVal := reflect.ValueOf(b)

	//获取rVal指针指向的值

	iVal := rVal.Elem()

	//设置 iVal指向的值

	iVal.SetInt(20)
}

//通过反射设置结构体字段的值
func Reflect04 (b interface{}) {
	rVal := reflect.ValueOf(b).Elem()

	//先找到对应的属性，然后设置其值
	rVal.FieldByName("Name").SetString("Smith")

	rVal.FieldByName("Age").SetInt(30)
}