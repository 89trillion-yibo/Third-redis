package router

import (
	"awesomeProject/Testthird/internal/ctrl"
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) *gin.Engine {
	r.POST("/creategif", ctrl.CreateCode)
	r.POST("/getGifcode",ctrl.GetGifcode)
	r.POST("/reward",ctrl.IncreGife)
	return r
}
