/**
*FileName: email
*Create on 2018/12/4 上午1:43
*Create by mok
*/

package email

import (
	"html/template"
	"github.com/go-gomail/gomail"
	"bytes"
)

type Email interface {
	Headers()map[string]string
	Template() *template.Template
	DialerInfo()*EmailDialerInfo
	Data() interface{}
}

func SendEmail(email Email)error{
	m := gomail.NewMessage()
	headers := email.Headers()
	setHeaders(m,headers)
	buf := &bytes.Buffer{}
	t := email.Template()
	t.Execute(buf,email.Data())
	m.SetBody("text/html",buf.String())
	dialerInfo := email.DialerInfo()
	d := gomail.NewDialer(dialerInfo.Host,dialerInfo.Port,dialerInfo.Username,dialerInfo.Password)
	err:=d.DialAndSend(m)
	return err
}

func setHeaders(m *gomail.Message,headers map[string]string){
	for field,val := range headers{
		m.SetHeader(field,val)
	}
}



type EmailConfig struct {
	emailTemplate   *template.Template
	emailHeaders     map[string]string
}

type EmailDialerInfo struct {
	Host		string
	Port		int
	Username 	string
	Password	string
	SSL        bool
}


type EmailContainer struct {
	H NewsHandler
	data interface{}
	*EmailConfig
	*EmailDialerInfo
}

func (c *EmailContainer)Headers()map[string]string{
	c.emailHeaders = map[string]string{
		"To": c.H.GetNews()["To"].(string),
		"From": "MokShop电子商城",
	}
	return c.emailHeaders
}

func(c *EmailContainer)Template()*template.Template{
	return c.emailTemplate
}

func(c *EmailContainer)DialerInfo()*EmailDialerInfo{
	return c.EmailDialerInfo
}

func(c *EmailContainer)Data() interface{}{
	c.data = c.H.GetNews()["data"]
	return c.data
}
