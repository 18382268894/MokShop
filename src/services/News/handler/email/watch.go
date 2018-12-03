/**
*FileName: email
*Create on 2018/12/4 上午4:25
*Create by mok
 */

package email

import (
	"MokShop/src/services/News/utils"
)

func WatchAll() {
	watchRegister("User_Register_Email","register_ch",1)
	watchChangePassword("User_ChangePassword_Email","ChangePassword_ch",1)
	watchOrderCommit("Order_OrderCommit_Email","OrderCommit_ch",5)
	watchOrderDel("Order_OrderDel_Email","OrderDel_ch",2)
	watchPayFinish("Order_PayFinish_Email","PayFinish_ch",3)
}

func watchRegister(topic, ch string, num int) {
	for i := 0; i < num; i++ {
		go func() {
			h := new(RegisterHandler)
			utils.NewConsumer(topic, ch, h)
			c := &EmailContainer{
				H: h,
			}
			SendEmail(c)
		}()
	}

}

func watchChangePassword(topic, ch string, num int){
	for i := 0; i < num; i++ {
		go func() {
			h := new(ChangePasswordHandler)
			utils.NewConsumer(topic, ch, h)
			c := &EmailContainer{
				H: h,
			}
			SendEmail(c)
		}()
	}
}

func watchOrderCommit(topic, ch string, num int){
	for i := 0; i < num; i++ {
		go func() {
			h := new(OrderCommitHandler)
			utils.NewConsumer(topic, ch, h)
			c := &EmailContainer{
				H: h,
			}
			SendEmail(c)
		}()
	}
}

func watchPayFinish(topic, ch string, num int){
	for i := 0; i < num; i++ {
		go func() {
			h := new(PayFinishHandler)
			utils.NewConsumer(topic, ch, h)
			c := &EmailContainer{
				H: h,
			}
			SendEmail(c)
		}()
	}
}


func watchOrderDel(topic, ch string, num int){
	for i := 0; i < num; i++ {
		go func() {
			h := new(OrderDelHandler)
			utils.NewConsumer(topic, ch, h)
			c := &EmailContainer{
				H: h,
			}
			SendEmail(c)
		}()
	}
}

