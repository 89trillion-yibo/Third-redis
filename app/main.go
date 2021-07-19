package main

import (
	"awesomeProject/Testthird/app/http"

	"awesomeProject/Testthird/utils"

	"fmt"

)

func main() {

	//初始化redis
	err := utils.RedisCli("127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
	}
	//启动服务
	http.InitRun()

}
