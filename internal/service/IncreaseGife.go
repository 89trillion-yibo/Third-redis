package service

import (
	"awesomeProject/Testthird/utils"
	"encoding/json"
	"fmt"
	"strconv"
)

//给用户增加奖励
func Increase(giftcode string, id string) interface{} {
	//获取奖励内容
	get := utils.HashGet(giftcode, "Pack")
	bytes := []byte(get)
	tmp := make(map[string]interface{})
	err := json.Unmarshal(bytes, &tmp)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("tmp>>>>>>>>>>",tmp)
	//获取用户信息
	//获取用户金币并添加奖励金币
	gold := utils.HashGet(id, "Gold")
	atoi, _ := strconv.Atoi(gold)
	newgold := float64(atoi) + tmp["Gold"].(float64)
	update, err := utils.ValueUpdate(id, "Gold", newgold)
	fmt.Println("update>>>>>>>>>>",update)

	//获取用户钻石，并获取奖励钻石
	dia := utils.HashGet(id, "Diamond")
	atoiDia, _ := strconv.Atoi(dia)
	newDia := float64(atoiDia) + tmp["Diamond"].(float64)
	updateDia, _ := utils.ValueUpdate(id, "Diamond", newDia)
	fmt.Println("update>>>>>>>>>>",updateDia)
	return get
}
