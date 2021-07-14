package service

import (
	"awesomeProject/third/redisdo"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
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
func CareatGif(gifType int, des string, allowTime int, valTime string, pack string, createName string) (string,interface{})  {
	//转换时间
	duration, _ := time.ParseDuration(valTime)

	//存入数据
	mapdate := make(map[string]interface{})
	mapdate["Type"] = gifType
	mapdate["Description"] = des
	mapdate["AllowedTimes"] = allowTime
	mapdate["ValidPeriod"] = time.Now().Add(duration)
	mapdate["Pack"] = pack
	mapdate["CreateName"] = createName
	mapdate["CreateTime"] = time.Now()
	mapdate["ByTime"] = 0

	//生成8位随机数
	randomCode := GetRandomString(8)
	return randomCode,mapdate
}

//创建对应礼品码领取列表
func Receive(gifcode string) error {
	mapdate := make(map[string]interface{})

	set, err := redisdo.HashSet(gifcode+":Receive", mapdate)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(set)
	return nil
}

//判断礼品码是否合法
func Judgment(id string, gifcode string) (bool, error) {
	if id == "" || gifcode == "" {
		return false,errors.New("参数为空")
	}
	//判断是否过期
	get := redisdo.HashGet(gifcode, "ValidPeriod")
	tmpTime := time.Now()
	nowTime := tmpTime.Format("2006-01-02 15:04:05")
	parseValid, err1 := time.Parse("2006-01-02 15:04:05", get)
	parseNow, err2 := time.Parse("2006-01-02 15:04:05", nowTime)
	if   err1 == nil && err2 == nil && parseValid.Before(parseNow) {
		return false,errors.New("礼品码过期")
	}

	//判断是否合法
	//判断redis中是否含有该礼品码的领取列表
	//如果没有，则说明之前没用户领取，则创建新map存入redis
	Type := redisdo.HashGet(gifcode, "Type")
	if !redisdo.HasKey(gifcode+":receive") {
		receiveMap := make(map[string]interface{})
		receiveMap[id] = time.Now()
		redisdo.HashSet(gifcode+":receive",receiveMap)
		return true,nil
	}else {
		ids := redisdo.HasHashKay(gifcode + ":receive")
		fmt.Println("ids>>>>>>>>>>>>>",ids)
		switch Type {
		//第一类兑换码：看切片ids中是否含有该用户id，如果有说明已经领取过，没有的话则返回true,并添加用户至redis标记为已领取
		case "1":
			for i := 0; i < len(ids); i++ {
				if id == ids[i]{
					return false,errors.New("用户已经领取该礼品码")
				}
			}
			receiveMap := make(map[string]interface{})
			receiveMap[id] = time.Now()
			redisdo.HashSet(gifcode+":receive",receiveMap)
			return true,nil
		//第二类兑换码：限制次数兑换，看领取次数是否还大于1，如果小于则领取次数不足
		case "2":
			AllowedTimes := redisdo.HashGet(gifcode, "AllowedTimes")
			atoi, _ := strconv.Atoi(AllowedTimes)
			if atoi < 1 {
				return false,errors.New("领取次数不足")
			}
			atoi -= 1
			itoa := strconv.Itoa(atoi)
			redisdo.ValueUpdate(gifcode,"AllowedTimes",itoa)
			receiveMap := make(map[string]interface{})
			receiveMap[id] = time.Now()
			redisdo.HashSet(gifcode+":receive",receiveMap)
			return true,nil
		//第三类兑换码：不限用户和次数，直接返回true
		case "3":
			receiveMap := make(map[string]interface{})
			receiveMap[id] = time.Now()
			redisdo.HashSet(gifcode+":receive",receiveMap)
			return true,nil
		}
	}
	return false,errors.New("错误，未进判断")
}

//给用户增加奖励
func Increase(giftcode string, id string) interface{} {
	//获取奖励内容
	get := redisdo.HashGet(giftcode, "Pack")
	bytes := []byte(get)
	tmp := make(map[string]interface{})
	err := json.Unmarshal(bytes, &tmp)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("tmp>>>>>>>>>>",tmp)
	//获取用户信息
	//获取用户金币并添加奖励金币
	gold := redisdo.HashGet(id, "Gold")
	atoi, _ := strconv.Atoi(gold)
	newgold := float64(atoi) + tmp["Gold"].(float64)
	update, err := redisdo.ValueUpdate(id, "Gold", newgold)
	fmt.Println("update>>>>>>>>>>",update)

	//获取用户钻石，并获取奖励钻石
	dia := redisdo.HashGet(id, "Diamond")
	atoiDia, _ := strconv.Atoi(dia)
	newDia := float64(atoiDia) + tmp["Diamond"].(float64)
	updateDia, _ := redisdo.ValueUpdate(id, "Diamond", newDia)
	fmt.Println("update>>>>>>>>>>",updateDia)
	return get
}