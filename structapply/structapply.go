package structapply

/**
	go 中的对象是比较“另类”，是以struct的方式来实现的。

 */

 //声明struct的时候等于声明对象的属性
type Person struct {
	Name string
	Age  int
}

//封装类方法
func (p *Person)say(){
	print("I am name is :",p.Name," and I am  :",p.Age," years old")
}

//继承
type Student struct {
	Person
	school string
}

type Tearch struct {
	Person
	subject string
}

func Demo() {
	stu := Student{Person{"笑话",20},"wangtang"}
	//使用继承过来的方法
	stu.say()
}


