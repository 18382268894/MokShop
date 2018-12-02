/**
*FileName: model
*Create on 2018/12/2 上午4:46
*Create by mok
 */

package model

type Order struct {
	DeliverPreID uint64 `db:"deliver_pre_id",json:"deliver_pre_id"`
	UserInfo
	OrderDetail
}

type OrderDetail struct {
	OrderID    uint64    `db:"order_id",json:"order_id"`
	GoodID     uint64    `db:"good_id",json:"good_id"`
	GoodName   string    `db:"good_name",json:"good_name"`
	Price      float64   `db:"good_price",json:"price"`
	Quantity   uint8     `db:"quantity",json:"quantity"`
	Cost       float64   `db:"order_cost",json:"cost"`
	CreateTime int64     `db:"order_create_time",json:"create_time"`
	ModifyTime int64     `db:"order_modify_time",json:"modify_time"`
	Coupons    []*Coupon `db:"coupons",json:"coupons"`
}

//快递信息
type Deliver struct {
	PreID       uint64  `db:"pre_id",json:"pre_id"`
	DeliverID   uint64  `db:"deliver_id",json:"deliver_id"`
	Type        uint8   `db:"deliver_type",json:"type"`
	Cost        float64 `db:"deliver_cost",json:"cost"`
	Status      uint32  `db:"status",json:"status"`
	DeliverName string  `db:"deliver_name",json:"deliver_name"`
	CreateTime  int64   `db:"deliver_create_time",json:"create_time"`
	Listener    string  `db:"listener",json:"listener"` //deliver的具体内容
	UserInfo
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
