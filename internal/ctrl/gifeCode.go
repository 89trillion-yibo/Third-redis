package ctrl

import (
	"awesomeProject/Testthird/internal/gifeerror"
	"awesomeProject/Testthird/internal/model"
	"awesomeProject/Testthird/internal/service"
	"awesomeProject/Testthird/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateCode(c *gin.Context)  {

	//接收参数
	gifType, a := c.GetPostForm("gifType")
	des , b := c.GetPostForm("des")
	allowTime, d := c.GetPostForm("allowTime")
	valTime, e := c.GetPostForm("valTime")
	createName, f := c.GetPostForm("createName")
	gold, g := c.GetPostForm("gold")
	diamond, h := c.GetPostForm("diamond")
	//判断是否有参数
	if !a && !b && !d && !e && !f && !g && !h {
		c.JSON(http.StatusBadRequest,gifeerror.Parameters)
	}else {
		//将string转成int
		giftype, _ := strconv.Atoi(gifType)
		alltime, _ := strconv.Atoi(allowTime)
		jinbi, _ := strconv.Atoi(gold)
		zuanshi, _ := strconv.Atoi(diamond)
		//初始化礼品内容
		s := model.Pack{Gold: jinbi,Diamond: zuanshi}
		marshal, _ := json.Marshal(s)
		pack := string(marshal)
		//创建礼品码
		careatGif, code := service.CareatGif(giftype, des, alltime, valTime, pack, createName)
		fmt.Println("code:",code)

		c.JSON(http.StatusOK,gifeerror.OK.WithData(careatGif))	
	}
}

func GetGifcode(c *gin.Context) {
	//接收参数
	gifcode, a := c.GetPostForm("gifcode")
	if !a {
		c.JSON(http.StatusBadRequest,gifeerror.Parameters)
	}else {
		get, err := utils.HashGetAll(gifcode)
		if err!=nil {
			c.JSON(http.StatusBadGateway,gifeerror.Error.WithData(err))
		}
		c.JSON(http.StatusOK,gifeerror.OK.WithData(get))
	}

}

func IncreGife(c *gin.Context) {
	id,a := c.GetPostForm("id")
	gifcode,b := c.GetPostForm("gifcode")
	if !a && !b {
		c.JSON(http.StatusBadRequest,gifeerror.Parameters)
	}else {
		//判断礼品码是否合法以及用户是否有资格领取
		judgment, err := service.Judgment(id, gifcode)
		if judgment && err == nil {
			increase := service.Increase(gifcode, id)
			/*get := redisdo.HashGet(gifcode, "Pack")
			bytes := []byte(get)
			tmp := make(map[string]interface{})
			json.Unmarshal(bytes, &tmp)*/
			c.JSON(http.StatusOK,gifeerror.OK.WithData(increase))
		}else {
			c.JSON(http.StatusBadGateway,gifeerror.Exchange)
		}
	}
}
