package gift

import "time"

//礼品码结构
type GiftCode struct {
	Type          int       //礼品码类型
	Description   string    //礼品描述
	AllowedTimes  int      //可领取次数
	ValidPeriod   time.Time   //有效期
    Pack          string     //礼品内容列表
    ByTime        int      //已领取次数
    CreateName    string   //创建人员
    CreateTime    time.Time   //创建时间
}

//礼品内容结构
type Pack struct {
	Gold int      //金币
	Diamond int   //钻石
}

//领取列表结构
type ClaimList struct {
	Id   string     //用户id
	Time time.Time  //领取时间
}

//用户结构
type User struct {
	Id string
	Gold int
	Diamond int
}


