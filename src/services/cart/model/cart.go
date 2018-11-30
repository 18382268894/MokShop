/**
*FileName: model
*Create on 2018/12/1 上午4:34
*Create by mok
 */

package model

type Cart struct {
	UID uint64 `db:"uid",json:"uid"`
	Item
}


type Item struct {
	ItemID uint64 `db:"item_id",json:"item_id"`
	GoodID uint64 `db:"good_id",json:"good_id"`
	Price float64 `db:"price",json:"price"`
	GoodName string `db:"good_name",json:"good_name"`
	Quantity uint32 `db:"quantity",json:"quantity"`
	CreateTime int64 `db:"create_time",json:"create_time,omitempty"`
	ModifyTime int64 `db:"modify_time",json:"modify_time,omitempty"`
}