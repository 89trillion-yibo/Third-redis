package handler

import (
	"awesomeProject/Testthird/internal/model"
	"fmt"
	"math/rand"
	"time"
)

// 随机生成指定位数的大写字母和数字的组合
func  GetRandomString(l int) string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//创建礼品码
func CareatGif(gifType int, des string, allowTime int, valTime string, pack string, createName string) (string,map[string]interface{})  {
	//转换时间
	duration, _ := time.ParseDuration(valTime)

	//存入数据
	mapdate := make(map[string]interface{})
	mapdate["Type"] = gifType
	mapdate["Description"] = des
	mapdate["AllowedTimes"] = allowTime
	mapdate["ValidPeriod"] = time.Now().Add(duration).Unix()
	mapdate["Pack"] = pack
	mapdate["CreateName"] = createName
	mapdate["CreateTime"] = time.Now().Unix()
	//mapdate["ByTime"] = 0

	//生成8位随机数
	randomCode := GetRandomString(8)

	//存入redis
	_, err := model.HashSet(randomCode, mapdate)
	if err!=nil {
		fmt.Println(err)
	}

	//创建已领取次数
	setBool, err := model.StringSet(randomCode+":Bytime", 0)
	if err!=nil || !setBool{
		fmt.Println(err)
	}

	return randomCode,mapdate
}
