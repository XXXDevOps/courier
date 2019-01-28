package media

import (
	"crypto/tls"
	"github.com/NioDevOps/courier/cfg"
	"github.com/NioDevOps/courier/models"
	"gopkg.in/gomail.v2"
)

const SMTPTYPE string = "smtp"

type StmpMedia struct {
	BaseMedia
	Obj *gomail.Dialer
}

func NewSmtpMedia(cfgStruct *cfg.MediaCfgStruct) *StmpMedia {
	sm := &StmpMedia{}
	if cfgStruct.NeedAuth {
		sm.Obj = gomail.NewDialer(cfgStruct.Host, cfgStruct.Port, cfgStruct.Username, cfgStruct.Password)
		sm.Obj.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		// sm.Obj.TLSConfig = &tls.Config{InsecureSkipVerify: !cfgStruct.NeedAuth}
	} else {
		sm.Obj = &gomail.Dialer{Host: cfgStruct.Host, Port: cfgStruct.Port}
	}
	sm.Name = cfgStruct.Name
	return sm
}

func (sm *StmpMedia) Send(t *models.Transaction) error {
	m := gomail.NewMessage()
	m.SetHeader("From", t.Sender.Address)
	m.SetHeader("To", t.Receivers.Addresses()...)
	m.SetHeader("Cc", t.Cc.Addresses()...)
	m.SetHeader("Subject", t.Message.Subject)
	m.SetBody("text/plain", t.Message.Content)
	return sm.Obj.DialAndSend(m)
}

func (sm *StmpMedia) Init() error {
	return nil
}

func (sm *StmpMedia) Redial() error {
	return nil
}

func (sm *StmpMedia) Close() error {
	return nil
}
