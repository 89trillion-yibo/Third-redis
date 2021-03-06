package model

//礼品码结构
type GiftCode struct {
	Type          int       //礼品码类型
	Description   string    //礼品描述
	AllowedTimes  int      //可领取次数
	ValidPeriod   int64   //有效期
	Pack          string     //礼品内容列表
	CreateName    string   //创建人员
	CreateTime    int64   //创建时间
}

//礼品内容结构
type Pack struct {
	Gold int      //金币
	Diamond int   //钻石
}