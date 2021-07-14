package redisdo

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
)

var ctx = context.Background()

var RedisClient *redis.Client

//连接redis数据库
func RedisCli(httpport string) (*redis.Client,error)  {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     httpport,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Println("connect redis failed")
		return nil, err
	}
	fmt.Println("RedisClient:",RedisClient)
	return RedisClient,nil
}

//往redis存入
func HashSet(key string, mapData interface{}) (bool, error) {
	//参数验证
	if key == "" || mapData == nil {
		return false,errors.New("参数为空")
	}

	if RedisClient == nil {
		fmt.Println("客户端为空")
		return false,errors.New("客户端链接断开")
	}else {
		if err := RedisClient.HSet(ctx, key, mapData).Err(); err != nil {
			fmt.Println(err)
			return false, errors.New("添加失败")
		} else {
			return true, nil
		}
	}
}

//从redis中取所有
func HashGetAll(key string) (interface{},error){
	//参数非空判断
	if key == ""  {
		return "",errors.New("参数为空")
	}
	value, err := RedisClient.HGetAll(ctx, key).Result()
	if err == redis.Nil{
		return "",errors.New("key不存在")
	}else if err!= nil {
		return "",errors.New("获取失败")
	}
	return value,nil
}

//从redis的map中取一个属性
func HashGet(key string,field string) (string){
	//参数非空判断
	if key == "" || field == "" {
		return ""
	}
	value := RedisClient.HGet(ctx, key,field).Val()
	return value
}

//判断是否含有该key
func HasKey(keyname string) (bool) {
	//参数非空判断
	if keyname == "" {
		return false
	}
	exists,err := RedisClient.Exists(ctx, keyname).Result()
	if err != nil {
		fmt.Println(err)
	}
	if exists == 1 {
		return true
	}else {
		return false
	}
}

//获取map中的所有value值
func HasHashKay(keyname string) []string {
	//参数非空判断
	if keyname == "" {
		return nil
	}
	exists,err := RedisClient.HKeys(ctx, keyname).Result()
	if err != nil {
		fmt.Println(err)
	}
	return exists
}

//redis中map修改值
func ValueUpdate(key string, field string, newvalue interface{}) (bool,error) {
	result, err := RedisClient.HMSet(ctx, key, field, newvalue).Result()
	if err!=nil{
		fmt.Println(err)
	}
	return result,nil
}