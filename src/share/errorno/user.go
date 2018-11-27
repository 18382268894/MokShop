/**
*FileName: share
*Create on 2018/11/26 下午5:27
*Create by mok
*/

package errorno
import "github.com/micro/go-micro/errors"

const(
	USER_REGIST_SUCCESS int32 =2000
	USER_REGIST_FAIL int32 = 20001
	USER_LOGIN_SUCCESS int32 = 20002
	USER_LOGIN_FAIL int32 = 20003
)

var(
	EMPTY_USERNAME = errors.New(NewErrID(),"用户名不能为空",USER_REGIST_FAIL)
	EMPTY_PASSWORD = errors.New(NewErrID(),"密码不能为空",USER_REGIST_FAIL)
	EMPTY_EMAIL = errors.New(NewErrID(),"邮箱不能为空",USER_REGIST_FAIL)
	WRONG_EMAIL = errors.New(NewErrID(),"邮箱格式不正确",USER_REGIST_FAIL)
	CREATE_FAIL = errors.New(NewErrID(),"服务故障，注册失败",USER_REGIST_FAIL)
	CREATE_SUCCESS = errors.New(NewErrID(),"注册用户成功",USER_REGIST_SUCCESS)
	LOGIN_FAIL = errors.New(NewErrID(),"登录失败",USER_LOGIN_FAIL)
	LOGIN_SUCCESS = errors.New(NewErrID(),"登录成功",USER_LOGIN_SUCCESS)
)