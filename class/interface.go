package class

import "fmt"

//声明一个interface 作为基类
type Animal interface {
	Sleep()
	GetAge() int
}

//设置子类 cat 实现Animal中的数据
type Cat struct {
	Name string
	Age int
}

func (cat *Cat)Sleep() {
	fmt.Println("cat need to sleep")
}

func (cat *Cat) GetAge() int {
	return cat.Age
}

//设置子类dog 实现Animal

type Dog struct {
	Name string
	Age int
}

func (dog *Dog) Sleep() {
	fmt.Println("dog need to sleep")
}

func (dog *Dog) GetAge() int {
	return dog.Age
}

//工厂模式
func Factory(name string) Animal{
	switch name {
	case "cat":
		return &Cat{"xiaohua",20}
	case "dog":
		return &Dog{"xiaoqiang",30}
	default:
		panic("no this animal")
	}

}