package model

import "time"

//领取列表结构
type ClaimList struct {
	Id   string     //用户id
	Time time.Time  //领取时间
}
