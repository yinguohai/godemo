package mysql

import (
	"fmt"
	"github.com/go-xorm/xorm"
)

type Userinfo struct {
	Id int	`xorm:"int notnull pk autoincr"`
	Name string	`xorm:"varchar(32)"`
	Age string	`xorm:"varchar(32)"`
	Sex string	`xorm:"varchar(32)"`
}

/**
	判断
 */
func (u *Userinfo)getClient() *xorm.Engine{
	dataSourceName := `root:root@tcp(localhost:3306)/test?charset=utf8`

	engine , err := xorm.NewEngine("mysql",dataSourceName)

	if err != nil {
		panic(err)
	}
	return engine
}

/**
	IsTableExist 判断表是否存在
 */
func XormIsTableExist() {
	user1 := Userinfo{}

	engine := user1.getClient()
	ok , _ := engine.IsTableExist("userinfo2")

	fmt.Println(ok)
}

/**
	Insert 插入一条或者多条记录
 */
func XormInsert(){
	user1 := Userinfo{
		Name:"xiaohua",
		Age:"35",
		Sex:"Girl",
	}

	user2 := Userinfo{
		Name:"James",
		Age:"32",
		Sex:"Boy",
	}
	engine := user1.getClient()
	ok ,_ := engine.Insert(&user1,&user2)

	fmt.Println(ok)
}


/**
	查找 Find
	1. 初始化一个Userinfo 数组类型的变量
	2. 获取engine
	3. Select 或者 Cols指定需要显示的字段
	4. Where 或者 In等关键字来限制查找的范围
	5. Find 查找出所有结果，并把user的指针传入到Find中，用于接收结果集

 */
func XormSelect() {
	//user 声明一个slice 用来接收表中的数据
	user := make([]Userinfo,0)
	client := Userinfo{}
	engine := client.getClient()
	//engine.Where(`name=?`,"xiaoming").Find(&user)
	//engine.Desc("id").Find(&user)
	//利用select 设置查询指定的字段，并且使用 In来设置查询范围
	//engine.Select(`id,name`).In(`id`,1,2,3).Find(&user)

	//Cols 和 select 都可以用来指定需要查询的字段
	engine.Cols("id","name").In("id",1,2,3,4).Find(&user)

	for k,v := range user {
		fmt.Println(k,v)
	}
}

/**
	更新 Update
	1. 声明并初始化一个 Userinfo 结构体
	2. 获取engine
	3. 通过Cols指定需要更新的字段，并通过Update进行批量更新【传入user进去】
 */
func XormUpdate(){
	user := Userinfo{
		Name:"huahua",
	}

	client := Userinfo{}

	engine := client.getClient()

	//把所有的记录中的name 都更新为 "huahua"
	ok , err := engine.Cols(`name`).Update(&user)

	if err != nil {
		panic(err)
	}

	fmt.Println("更新成功==>", ok)
}
