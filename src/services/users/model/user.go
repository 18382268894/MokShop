/**
*FileName: model
*Create on 2018/11/26 下午7:23
*Create by mok
*/

package model


type User struct {
	UID                  uint64   `json:"UID,omitempty",,db:"uid"`
	Name                 string   `json:"name,omitempty",,db:"name"`
	Password             string   `json:"password,omitempty",,db:"password"`
	Phone                int64    `json:"phone,omitempty",,db:"phone"`
	Email                string   `json:"email,omitempty",,db:"email"`
	Status               int32    `json:"status,omitempty",,db:"status"`
	Ip                   string   `json:"ip,omitempty",,db:"ip"`
	LastTime             int64    `json:"lastTime,omitempty",,db:"last_time"`
	Score                int32    `json:"score,omitempty",db:"score"`
}
