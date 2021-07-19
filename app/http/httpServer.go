package http

import (
	"awesomeProject/Testthird/internal/router"
	"github.com/gin-gonic/gin"
)

func InitRun() error {
	r := gin.Default()
	//启动路由
	router.Routers(r)
	err := r.Run()
	if err!=nil{
		return err
	}
	return nil
}
