/**
*FileName: db
*Create on 2018/12/1 上午5:07
*Create by mok
*/

package db

import (
	"context"
	"time"
	"database/sql"
	proto"MokShop/src/share/proto/cart"
)

//增加商品到购物车
func AddCargo(UID ,itemID ,goodID uint64,goodName string,price float64,quantity uint32,createTime,modifyTime int64)error{
	ctx,_ := context.WithTimeout(context.TODO(),3*time.Second)
	tx,err := DB.BeginTx(ctx,&sql.TxOptions{Isolation:sql.LevelRepeatableRead,ReadOnly:false})
	if err != nil{
		return err
	}
	sqlstr := "insert into `cart`(uid,item_id) values(?,?)"
	if _,err = tx.Exec(sqlstr,UID,itemID);err != nil{
		tx.Rollback()
		return err
	}
	sqlstr = "insert into `items`(item_id,good_id,good_name,price,quantity,create_time,modify_time) values(?,?,?,?,?,?,?)"
	if _,err = tx.Exec(sqlstr,itemID,goodID,goodName,price,quantity,createTime,modifyTime);err != nil{
		tx.Rollback()
		return err
	}
	if err = tx.Commit();err != nil{
		tx.Rollback()
		return err
	}
	return nil
}


//删除购物车中的某一个商品
func DelCargo(itemID uint64)error{
	sqlstr := "delete from cart where item_id=?"
	_,err := DB.Exec(sqlstr,itemID)
	return err
}

//删除整个购车中所有的商品
func DelAllCargo(UID uint64)error{
	sqlstr := "delete from cart where uid=?"
	_,err := DB.Exec(sqlstr,UID)
	return err
}


//修改购物车中的商品数量
func Update(itemID uint64,quantity uint32)error{
	sqlstr := "update items set quantity=? where itemID=?"
	_,err := DB.Exec(sqlstr,quantity,itemID)
	return err
}

//展示购物车中所有商品
func Display(UID uint64)([]*proto.Item, error){
	var items = make([]*proto.Item,0)
	sqlstr := "select * from `itmes` where items.item_id in (select item_id from `cart` where uid=?)"
	if err := DB.Select(&items,sqlstr,UID);err != nil{
		return nil,err
	}
	return items,nil
}
