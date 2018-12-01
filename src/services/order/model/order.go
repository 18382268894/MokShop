/**
*FileName: model
*Create on 2018/12/2 上午4:46
*Create by mok
 */

package model

type Order struct {
	UserInfo
	Coupon
	Deliver
	OrderID    uint64  `db:"order_id",json:"order_id"`
	GoodID     uint64  `db:"good_id",json:"good_id"`
	GoodName   string  `db:"good_name",json:"good_name"`
	Price      float64 `db:"good_price",json:"price"`
	Quantity   uint8	`db:"quantity",json:"quantity"`
	Cost       float64 `db:"order_cost",json:"cost"`
	CreateTime int64   `db:"order_create_time",json:"create_time"`
	ModifyTime int64   `db:"order_modify_time",json:"modify_time"`
}

//快递信息
type Deliver struct {
	Type        uint8   `db:"deliver_type",json:"type"`
	Cost        float64 `db:"deliver_cost",json:"cost"`
	DeliverName string  `db:"deliver_name",json:"deliver_name"`
	DeliverID   string  `db:"deliver_id",json:"deliver_id"`
	CreateTime  int64   `db:"deliver_create_time",json:"create_time"`
}

//用户信息
type UserInfo struct {
	UID      uint64 `db:"uid",json:"uid"`
	Username string `db:"username",json:"username"`
	Email    string `db:"email",json:"email"`
	Phone    uint64 `db:"phone",json:"phone"`
	Addr     string `db:"addr",json:"addr"`
}

//优惠券
type Coupon struct {
	CouponID       uint64  `db:"coupon_id",json:"coupon_id"`
	Much           float64 `db:"much",json:"much"`
	Type           uint8   `db:"coupon_type",json:"type"`
	ConditionType  uint8   `db:"condition_type",json:"condition_type"`
	ConditionLimit float64 `db:"condition_limit",json:"condition_limit"`
}
