package mongodbapply

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

/**
	获取mongodb 客户端
 */
func getClient() *mongo.Client {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	ctx2, _ := context.WithTimeout(context.Background(), 2*time.Second)

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		panic(err.Error())
	}

	//if you with to know if a MongoDB server has been found and connected to , use the Ping method
	err = mongoClient.Ping(ctx2, readpref.Primary())

	return mongoClient
}

/**
	获取限定时间的context对象
*/
func getContext(num int) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(num)*time.Second)
	return ctx
}

/**
	统计
 */
func FindCount() {
	client := getClient()

	collection := client.Database("Douban").Collection("movies")

	//bson.D represents a BSON Document , This type should be used in situations where orders matters
	// the condition equals : bson.D{{"page":bson.D{"$gt":"6"}}}  or bson.M{"page":bson.M{"$gt":"6"}}
	//
	//cur , err := collection.CountDocuments(ctx,bson.D{{"page",bson.M{"$gt":"6"}}})
	cur, err := collection.CountDocuments(getContext(10), bson.M{"page": bson.M{"$gt": "6"}})

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(cur)
}

/**
	查询单条
 */
func FindOne() {
	client := getClient()

	//接收查询出来的值
	//	var result interface{} , 虽然也可以接收不推荐，因为遍历出来的结果是集合，不是map， 需要遍历多次
	var result map[string]interface{}

	collection := client.Database("Douban").Collection("movies")

	if err := collection.FindOne(getContext(10), bson.M{"page": "6"}).Decode(&result); err != nil {
		panic(err.Error())
	}

	fmt.Println(result)
}

/**
	批量查询
 */
func FindMulti() {
	client := getClient()

	collection := client.Database("Douban").Collection("movies")

	//设置条件
	var option [] *options.FindOptions

	//分页
	option = append(option, options.Find().SetSkip(6))
	option = append(option, options.Find().SetLimit(10))
	//排序
	option = append(option, options.Find().SetSort(bson.M{"page": -1}))
	//filed字段显示
	option = append(option, options.Find().SetProjection(bson.M{"name": true, "page": true, "_id": false}))

	//bson.M{"page":bson.M{"$gt":"6"}}   ===  page > 6
	//cur, err := collection.Find(getContext(10),bson.M{"page":bson.M{"$gt":"6"}},options.Find().SetLimit(2),options.Find().SetSort(bson.M{"page":-1}))

	cur, err := collection.Find(getContext(10), bson.M{"page": bson.M{"$gt": "6"}}, option...)
	if err != nil {
		panic(err.Error())
	}

	defer cur.Close(getContext(10))

	var data []map[string]interface{}
	//通过指针向下逐个获取
	for cur.Next(getContext(10)) {
		var result map[string]interface{}
		//cur 是指针类型，直接使用Decodel来获取结果
		if err = cur.Decode(&result); err != nil {
			panic(err.Error())
		}
		data = append(data, result)
	}

	for k, v := range data {
		fmt.Println(k, "==>", v)
	}
}

/**
	单条插入
 */
func AddOne() {
	client := getClient()
	collection := client.Database("Person").Collection("Teacher")
	//成功则返回 _id 值的地址
	id, err := collection.InsertOne(getContext(10), bson.M{"_id": 11, "name": "Bob", "sex": "F", "age": 30, "hobbit": "like basketball"})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("insert success :", id)
}

/**
	批量插入
 */
func AddMulti() {
	client := getClient()
	collection := client.Database("Person").Collection("Teacher")
	//多条bson.M 的记录 可以使用 interface{} 这个类型来表示
	data := []interface{}{
		bson.M{"_id": 22, "name": "Smith", "age": 40, "sex": "M"},
		bson.M{"_id": 33, "name": "Hua", "age": 18},
	}
	//成功则返回所有 _id 值的地址
	result, err := collection.InsertMany(getContext(10), data)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(result)
}

/**
	更新一条
*/
func UpdateOneDemo() {
	client := getClient()
	collection := client.Database("Person").Collection("Teacher")
	//_id = 1  not  _id = "1" , 因为在初始化的时候_id 插入的是整型
	err :=collection.FindOneAndUpdate(getContext(10),bson.M{"_id":1},bson.M{"$set":bson.M{"age":"99"}})
	//更新成功则err.Err() 就是nil
	if err.Err() != nil {
		fmt.Println(err.Err())
	} else {
		fmt.Println("update one success")
	}
}

/**
	更新多条记录，
		------ 新增一个字段，或者修改字段都可以
 */
func UpdateMulti() {
	client := getClient()
	collection := client.Database("Person").Collection("Teacher")
	//返回{匹配记录数，已经修改的记录数，UpsertedCount数
	data , err:=collection.UpdateMany(getContext(10),bson.M{"_id":bson.M{"$gt":11}},bson.M{"$set":bson.M{"class":"Chinese"}})

	if err != nil {
		panic(err.Error())
	}
	fmt.Println(data)
}

/**
	Upsert : 存在则更新，不存在则插入
		----通过options.Update().SetUpsert(true) 可以开启
 */
func UpsertedDemo() {
	client := getClient()
	collection := client.Database("Person").Collection("Teacher")
	//&{0 0 1 44}  == {MatchedCount,ModifiedCount,UpsertedCount, _id of insert}
	//通过options.Update().SetUpsert(true) 可以开启
	data , err:=collection.UpdateMany(getContext(10),bson.M{"_id":44},bson.M{"$set":bson.M{"class":"Chinese"}},options.Update().SetUpsert(true))

	if err != nil {
		panic(err.Error())
	}
	fmt.Println(data)
}

/**
	单条删除
 */
func DeleteOneDemo() {
	client := getClient()
	collection := client.Database("Person").Collection("Teacher")
	result , err := collection.DeleteOne(getContext(10),bson.M{"_id":3})

	if err != nil {
		panic(err.Error())
	}
	fmt.Println(result)
}

/**
	批量删除
 */
func DeleteManyDemo() {
	client := getClient()
	collection := client.Database("Person").Collection("Teacher")
	//删除 _id >= 1 的所有记录 , 返回删除的条数
	result,err := collection.DeleteMany(getContext(10),bson.M{"_id":bson.M{"$gte":1}})

	if err!= nil {
		panic(err.Error())
	}

	fmt.Println(result)
}

/**
	删除 Student 中的douban 这个集合
 */
func DeleteCollection() {
	client := getClient()

	if err := client.Database("student").Collection("douban").Drop(getContext(10)); err != nil {
		panic(err.Error())
	} else {
		fmt.Println("delete success")
	}
}