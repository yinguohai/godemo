package jsonapply

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string `jsonapply:"username"`
	NickName string `jsonapply:"nickname"`
	Sex      string
	//继承的另类写法
	PrivateInfo struct {
		Age   int
		Phone string
	}
}

//struct to jsonapply
func StructToJson() string {
	user1 := User{
		UserName: "user1",
		NickName: "huahua",
		Sex:      "nv",
		//PrivateInfo 是个匿名的struct,实例化时也需要使用是struct类型的值来赋予
		PrivateInfo: struct {
			Age   int
			Phone string
		}{Age: 20, Phone: "13566228415"},
	}

	data, err := json.Marshal(user1)

	if err != nil {
		panic("jsonapply format error!!!")
	}

	return string(data)
}

//map to jsonapply  or  slice to jsonapply
func MapToJson() string {
	//map类型的slice
	//map声明方式：   map[keyType]valueType
	var mmp map[string]interface{}

	mmp = make(map[string]interface{})

	mmp["name"] = "xiaohua"
	mmp["age"] = 20
	mmp["sex"] = "nv"

	result, _ := json.Marshal(mmp)

	return string(result)
}

//Int to Json
func IntToJson() []uint8 {
	result, _ := json.Marshal(100)
	return result
}

/**
	Json To Struct
	Unmarshal ([]byte(string),&result) error :
		1.解析出来的json格式是和result的类型相关的，我们想解析成Struct ,就把它声明为struct
		2.返回的是error , 所以接收解析结果的result参数一定是传的引用过去接受的
 */
func JsonToStruct() {
	var userInfo User
	result := StructToJson()

	fmt.Println(result)

	json.Unmarshal([]byte(result), &userInfo)

}

/**
	Json to map
	同JsonToStruct一样，用于接收解析结果的result需要map类型 或者 interface{}类型
 */

func JsonToMap() {
	//想要解析出来的结果为map类型，就将接收参数userInfo 声明为map类型
 	var userInfo map[string]interface{}
 	data := StructToJson()
 	json.Unmarshal([]byte(data),&userInfo)
 	fmt.Println(userInfo)
}
