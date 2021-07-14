package main

import (
	"awesomeProject/third/gift"
	"awesomeProject/third/redisdo"
	"awesomeProject/third/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	_, err := redisdo.RedisCli("127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
	}
    //模拟创建用户
	/*a := make(map[string]interface{})
	a["Id"] = "2"
	a["Gold"] = 3000
	a["Diamond"] = 1500
	set, err := redisdo.HashSet("2", a)
	fmt.Println(set)
	fmt.Println(a)*/


	engine := gin.Default()
	//创建礼品码并存入redis中
	engine.POST("/creategif", func(c *gin.Context) {
		//接收参数
		codetype := c.PostForm("gifType")
		des := c.PostForm("des")
		allowTime := c.PostForm("allowTime")
		valTime := c.PostForm("valTime")
		createName := c.PostForm("createName")
		gold := c.PostForm("gold")
		diamond := c.PostForm("diamond")

		//将string转成int
		giftype, _ := strconv.Atoi(codetype)
		alltime, _ := strconv.Atoi(allowTime)
		jinbi, _ := strconv.Atoi(gold)
		zuanshi, _ := strconv.Atoi(diamond)
		//初始化礼品内容
		s := gift.Pack{Gold: jinbi,Diamond: zuanshi}
		marshal, _ := json.Marshal(s)
		pack := string(marshal)
		//创建礼品码
		careatGif, code := service.CareatGif(giftype, des, alltime, valTime, pack, createName)
		fmt.Println("code:",code)
		//存入redis
		_, err = redisdo.HashSet(careatGif, code)
		if err!=nil {
			fmt.Println(err)
		}
		c.JSON(200,careatGif)
	})
	//根据礼品码查询redis并返回结果
	engine.POST("/getGifcode", func(c *gin.Context) {
		//接收参数
		gifcode := c.PostForm("gifcode")
		get, err := redisdo.HashGetAll(gifcode)
		if err!=nil {
			fmt.Println(err)
		}
		c.JSON(200,get)
	})

	//用户输入礼品码返回结果
	engine.POST("/reward", func(c *gin.Context) {
		id := c.PostForm("id")
		gifcode := c.PostForm("gifcode")
		judgment, err2 := service.Judgment(id, gifcode)
        if judgment && err2 == nil {
			increase := service.Increase(gifcode, id)
			c.JSON(200,increase)
		}else {
			c.JSON(200,"兑换错误")
		}
	})
	engine.Run()
}
