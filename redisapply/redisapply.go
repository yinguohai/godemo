package redisapply

import (
	"fmt"
	"github.com/go-redis/redis"
)

func Connect()  *redis.Client{
	redisdb := redis.NewClient(&redis.Options{
		Addr:"127.0.0.1:6379",
		Password:"123456",
		DB:0,
	})
	//判断redis  server  是否还可以连通
	pong , err := redisdb.Ping().Result()

	if err != nil {
		panic(err)
	}

	if pong == "PONG" {
		return redisdb
	} else {
		panic("连接错误")
	}
}

func Get(){
	clinet := Connect()

	result := clinet.Get("name").Val()

	result1 := clinet.Get("age").Val()

	fmt.Println(result,result1)
}

func Lpush() {
	client := Connect()
	student := []string{
		"1",
		"2",
		"3",
	}
	client.LPush("student",student)
}

func Rpop() {
	clinet := Connect()
	bb := clinet.RPop("student")
	fmt.Println(bb)

}

func Pipeline(){
	clinet := Connect()

	//声明一个管道
	pipe := clinet.Pipeline()

	//组合多个命令扔入到管道中，管道中默认53个命令侯会自动触发执行
	//LPop student
	pipe.LPop("student")
	//get name
	pipe.Get("name")


	//统一执行上面的多个命令，手工指定
	//[lpop student: 3 get name: ygh] <nil>
	result , err := pipe.Exec()

	fmt.Println(result , err)
}