package service

import (
	"awesomeProject/Testthird/utils"
	"errors"
	"fmt"
	"strconv"
	"time"
)

//判断礼品码是否合法
func Judgment(id string, gifcode string) (bool, error) {
	if id == "" || gifcode == "" {
		return false,errors.New("参数为空")
	}
	//判断是否过期
	get := utils.HashGet(gifcode, "ValidPeriod")
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
	Type := utils.HashGet(gifcode, "Type")
	if !utils.HasKey(gifcode+":receive") {
		receiveMap := make(map[string]interface{})
		receiveMap[id] = time.Now()
		utils.HashSet(gifcode+":receive",receiveMap)
		return true,nil
	}else {
		ids := utils.HasHashKay(gifcode + ":receive")
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
			utils.HashSet(gifcode+":receive",receiveMap)
			return true,nil
		//第二类兑换码：限制次数兑换，看领取次数是否还大于1，如果小于则领取次数不足
		case "2":
			AllowedTimes := utils.HashGet(gifcode, "AllowedTimes")
			atoi, _ := strconv.Atoi(AllowedTimes)
			if atoi < 1 {
				return false,errors.New("领取次数不足")
			}
			atoi -= 1
			itoa := strconv.Itoa(atoi)
			utils.ValueUpdate(gifcode,"AllowedTimes",itoa)
			receiveMap := make(map[string]interface{})
			receiveMap[id] = time.Now()
			utils.HashSet(gifcode+":receive",receiveMap)
			return true,nil
		//第三类兑换码：不限用户和次数，直接返回true
		case "3":
			receiveMap := make(map[string]interface{})
			receiveMap[id] = time.Now()
			utils.HashSet(gifcode+":receive",receiveMap)
			return true,nil
		}
	}
	return false,errors.New("错误，未进判断")
}
