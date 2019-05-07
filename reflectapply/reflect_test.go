package reflectapply

import (
	"fmt"
	"testing"
)

func TestReflectTest(t *testing.T) {
	i := 10
	ReflectTest(i)
}

func TestReflect02(t *testing.T) {
	stu := student{
		Name:"tome",
		Age:20,
	}
	Reflect02(stu)
}

//测试通过反射修改值类型
func TestReflect03(t *testing.T) {
	var  num int= 10

	fmt.Println("Before Reflect03  , num = ",num)

	Reflect03(&num)

	fmt.Println("After Reflect03  , num = ",num)
}

//测试通过反射修改struct类型
func TestReflect04(t *testing.T) {
	stu := student{"bob",20}
	fmt.Println(stu)
	Reflect04(&stu)
	fmt.Println(stu)
}