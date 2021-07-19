package main

import (
	"awesomeProject/Testthird/internal/model"
	"awesomeProject/Testthird/internal/service"
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
	var s = model.Pack{Gold: 1000,Diamond: 100}
	marshal, _ := json.Marshal(s)
	pack := string(marshal)
	gifcode, gifemap := service.CareatGif(giftype, des, alltime, valTime, pack, createName)
	t.Log(gifcode)
	t.Log(gifemap)
}