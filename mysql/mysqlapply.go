package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Connect(){
	dataSourceName := `root:root@tcp(localhost:3306)/test?charset=utf8`

	db, err := sql.Open("mysql", dataSourceName)

	if  err == nil {
		fmt.Println("连接成功")
	}

	//? 是占位符，防止sql注入
	queryString := `UPDATE userinfo set name="Boss" where id=?`

	smt , err := db.Prepare(queryString)

	res , err :=smt.Exec(2)
	if err != nil {
		panic(err)
	}

	id , _ := res.RowsAffected()

	//if err != nil {
	//	panic(err)
	//}

	fmt.Println(id)
}

func Select(){
	dataSourceName := `root:root@tcp(localhost:3306)/test?charset=utf8`

	db, err := sql.Open("mysql", dataSourceName)

	if  err == nil {
		fmt.Println("连接成功")
	}

	rows , err := db.Query("select * from userinfo")

	//rows.Next() 返回的是一个bool,用于判断是否还有下一行
	for rows.Next() {
		var id int
		var name string
		var age  string
		var sex string
		//rows.Scan() 才是我们取数据的核心语句
		err = rows.Scan(&id,&name,&age,&sex)
		fmt.Println(id, name,age,sex)
	}

}

