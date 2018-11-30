/**
*FileName: db
*Create on 2018/12/1 上午5:09
*Create by mok
*/

package db

import (
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB


func Init()error{
	return nil
}

func Close(){
	DB.Close()
}
