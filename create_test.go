package main

import (
	"awesomeProject/third/gift"
	"awesomeProject/third/service"
	"encoding/json"
	"testing"
)

func TestCareatGif(t *testing.T)  {
	var (
		giftype = 1
		des = "金币"
		alltime = 3
		valTime = "10m"
		createName = "ccc"
	)
	var s = gift.Pack{Gold: 1000,Diamond: 100}
	marshal, _ := json.Marshal(s)
	pack := string(marshal)
	gifcode, gifemap := service.CareatGif(giftype, des, alltime, valTime, pack, createName)
	t.Log(gifcode)
	t.Log(gifemap)
}