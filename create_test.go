package main

import (
	"awesomeProject/Testthird/internal/handler"
	"awesomeProject/Testthird/internal/model"
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
	gifcode, gifemap := handler.CareatGif(giftype, des, alltime, valTime, pack, createName)
	t.Log(gifcode)
	t.Log(gifemap)
}