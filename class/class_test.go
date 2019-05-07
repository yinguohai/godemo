package class

import "testing"
//继承，封装测试
func TestDemo(t *testing.T) {
	Demo()
}

//工厂模式，通过interface来实现多态，继而构造工厂函数
func TestFactory(t *testing.T) {
	cat := Factory("cat")
	cat.Sleep()

	dog := Factory("dog")
	dog.Sleep()
}