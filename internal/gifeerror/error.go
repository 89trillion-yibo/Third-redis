package gifeerror

var (
	OK = response(200,"ok")
	Error = response(500,"error")

	Parameters = response(101,"获取参数错误")
	Exchange = response(102,"兑换错误")

)

//异常结构
type GifeErr struct {
	Data    interface{}      //返回数据
	Code    int              //错误码
	Message string           //错误信息
}

//不返回数据
func response(code int , message string) *GifeErr{
	return &GifeErr{
		Code: code,
		Message: message,
		Data: nil,
	}
}

//返回数据
func (gif *GifeErr) WithData(data interface{}) GifeErr {
	return GifeErr{
		Code: gif.Code,
		Message:  gif.Message,
		Data: data,
	}
}