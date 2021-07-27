package ctrl

import (
	"awesomeProject/Testthird/internal/gifeerror"
	"awesomeProject/Testthird/internal/handler"
	"awesomeProject/Testthird/internal/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
//创建礼品码
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
	if !a || !b || !d || !e || !f || !g || !h {
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
		careatGif, mapdate := handler.CareatGif(giftype, des, alltime, valTime, pack, createName)
		fmt.Println("code:",mapdate)

		c.JSON(http.StatusOK,gifeerror.OK.WithData(careatGif))	
	}
}

//查询礼品码
func GetGifcode(c *gin.Context) {
	//接收参数
	gifcode, a := c.GetPostForm("gifcode")
	if !a || gifcode == "" {
		c.JSON(http.StatusBadRequest,gifeerror.Parameters)
	}else {
		//获取礼品码信息
		data, err := model.HashGetAll(gifcode)
		//获取领取列表信息
		receive, err := model.HashGetAll(gifcode + ":receive")
		gifeAndReceive := make(map[string]interface{})
		gifeAndReceive["gifeInfo"] = data
		gifeAndReceive["receiveInfo"] = receive
		if err!=nil {
			c.JSON(http.StatusBadGateway,gifeerror.Error.WithData(err))
		}
		c.JSON(http.StatusOK,gifeerror.OK.WithData(gifeAndReceive))
	}

}

//用户增加奖励
func IncreGife(c *gin.Context) {
	id,a := c.GetPostForm("id")
	gifcode,b := c.GetPostForm("gifcode")
	if !a || !b || id == "" || gifcode == "" {
		c.JSON(http.StatusBadRequest,gifeerror.Parameters)
	}else {
		//判断礼品码是否合法以及用户是否有资格领取
		judgment, code := handler.Judgment(id, gifcode)
		if judgment{
			increase := handler.Increase(gifcode, id)
			/*get := utils.HashGet(gifcode, "Pack")
			bytes := []byte(get)
			tmp := make(map[string]interface{})
			json.Unmarshal(bytes, &tmp)
			fmt.Println(tmp)*/
			c.JSON(http.StatusOK,gifeerror.OK.WithData(increase))
		}else {
			switch code {
			case "102":
				c.JSON(http.StatusBadRequest,gifeerror.Exchange)
			case "103":
				c.JSON(http.StatusBadRequest,gifeerror.Expired)
			case "104":
				c.JSON(http.StatusBadRequest,gifeerror.Exhausted)
			case "105":
				c.JSON(http.StatusBadRequest,gifeerror.AlreadyReceived)
			}
		}
	}
}
