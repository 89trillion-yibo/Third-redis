package main

import (
	"awesomeProject/Testthird/app/http"
	"awesomeProject/Testthird/internal/model"

	"fmt"
)

func main() {

	//初始化redis
	err := model.RedisCli("127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
	}
	//启动服务
	http.InitRun()

}
