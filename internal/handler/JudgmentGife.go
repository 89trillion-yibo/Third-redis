package handler

import (
	"awesomeProject/Testthird/internal/model"
	"fmt"
	"strconv"
	"time"
)

//判断礼品码是否合法
func Judgment(id string, gifcode string) (bool, string) {
	//判断礼品码是否过期
	get := model.HashGet(gifcode, "ValidPeriod")
	ValidTime, _ := strconv.ParseInt(get,10,64)
	tmpTime := time.Now().Unix()
	if   ValidTime < tmpTime {
		return false,"103"
	}

	//判断是否合法
	//判断redis中是否含有该礼品码的领取列表
	//如果没有，则说明之前没用户领取，则创建新map存入redis
	Type := model.HashGet(gifcode, "Type")
	if !model.HasKey(gifcode+":receive") {
		receiveMap := make(map[string]interface{})
		receiveMap[id] = time.Now().Unix()
		model.HashSet(gifcode+":receive",receiveMap)
		incr := model.Incr(gifcode + ":Bytime")
		fmt.Println("自增之后已领取次数",incr)
		return true,""
	}else {
		ids := model.HasHashKay(gifcode + ":receive")
		fmt.Println("ids>>>>>>>>>>>>>",ids)
		switch Type {
		//第一类兑换码：仅限领取一次
		case "1":
			stringGet, _ := model.StringGet(gifcode + ":Bytime")
			if stringGet == "1" {
				return false,"104"
			}
			receiveMap := make(map[string]interface{})
			receiveMap[id] = time.Now().Unix()
			model.HashSet(gifcode+":receive",receiveMap)
			//自增1
			model.Incr(gifcode + ":Bytime")
			return true,""
		//第二类兑换码：限制次数兑换，根据已领取次数判断是否可以领取
		case "2":
			AllowedTimes := model.HashGet(gifcode, "AllowedTimes")
			Bytime, _ := model.StringGet(gifcode + ":Bytime")
			//判断剩余领取次数
			if Bytime >= AllowedTimes {
				return false,"104"
			}
			//判断用户是否已领取过
			for i := 0; i < len(ids); i++ {
				if id == ids[i] {
					return false,"105"
				}
			}
			//添加领取用户
			receiveMap := make(map[string]interface{})
			receiveMap[id] = time.Now().Unix()
			model.HashSet(gifcode+":receive",receiveMap)
			model.Incr(gifcode + ":Bytime")
			return true,""
		//第三类兑换码：不限用户和次数，检查用户是否领取过
		case "3":
			//检查是否重复领取
			for i := 0; i < len(ids); i++ {
				if id == ids[i] {
					return false, "105"
				}
			}
			receiveMap := make(map[string]interface{})
			receiveMap[id] = time.Now().Unix()
			model.HashSet(gifcode+":receive",receiveMap)
			model.Incr(gifcode + ":Bytime")
			return true,""
		}
	}
	return false,"102"
}
