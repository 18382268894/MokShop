/**
*FileName: handler
*Create on 2018/12/4 上午12:55
*Create by mok
 */

package email

import (
	"encoding/json"
	"github.com/nsqio/go-nsq"
)

type NewsHandler interface {
	GetNews() map[string]interface{}
}


//邮件信息格式
type News map[string]interface{}

//1.注册邮件
type RegisterHandler struct {
	News
}

func (h *RegisterHandler) GetNews() map[string]interface{} {
	return h.News
}

func (h *RegisterHandler) HandleMessage(message *nsq.Message) error {
	var news = make(News)
	if err := json.Unmarshal(message.Body, news); err != nil {
		return err
	}
	h.News = news
	return nil
}



//2.密码修改邮件
type ChangePasswordHandler struct {
	News
}

func (h *ChangePasswordHandler) GetNews() map[string]interface{} {
	return h.News
}

func (h *ChangePasswordHandler) HandleMessage(message *nsq.Message) error {
	var news = make(News)
	if err := json.Unmarshal(message.Body, news); err != nil {
		return err
	}
	h.News = news
	return nil
}



//3.订单提交邮件
type OrderCommitHandler struct {
	News
}

func (h *OrderCommitHandler) GetNews() map[string]interface{} {
	return h.News
}

func (h *OrderCommitHandler) HandleMessage(message *nsq.Message) error {
	var news = make(News)
	if err := json.Unmarshal(message.Body, news); err != nil {
		return err
	}
	h.News = news
	return nil
}

//4.支付完成邮件
type PayFinishHandler struct {
	News
}

func (h *PayFinishHandler) GetNews() map[string]interface{} {
	return h.News
}

func (h *PayFinishHandler) HandleMessage(message *nsq.Message) error {
	var news = make(News)
	if err := json.Unmarshal(message.Body, news); err != nil {
		return err
	}
	h.News = news
	return nil
}

//5.订单取消邮件
type OrderDelHandler struct {
	News
}

func (h *OrderDelHandler) GetNews() map[string]interface{} {
	return h.News
}

func (h *OrderDelHandler) HandleMessage(message *nsq.Message) error {
	var news = make(News)
	if err := json.Unmarshal(message.Body, news); err != nil {
		return err
	}
	h.News = news
	return nil
}
