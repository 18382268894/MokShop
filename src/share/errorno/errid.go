/**
*FileName: errorno
*Create on 2018/11/26 下午5:31
*Create by mok
*/

package errorno

import "github.com/satori/go.uuid"

//为每一个错误码建立一个全局唯一的ID

//err id
func NewErrID()string{
	id,_ := uuid.NewV4()
	return id.String()
}
