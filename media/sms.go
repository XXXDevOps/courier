package media

import (
	"bytes"
	"encoding/json"
	"github.com/NioDevOps/courier/cfg"
	"github.com/NioDevOps/courier/models"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

const SMSTYPE string = "sms"

type SmsMedia struct {
	BaseMedia
	SmsUrl      string
	ContentType string
}

type PhoneNum struct {
	Mobile     int `json:"mobile"`
	Nationcode int `json:"nationcode"`
}

type SMSData struct {
	Phones  []PhoneNum `json:"phone"`
	Message string     `json:"msg"`
}

func NewSmsMedia(cfgStruct *cfg.MediaCfgStruct) *SmsMedia {
	sms := &SmsMedia{}
	sms.SmsUrl = cfgStruct.Uri
	sms.Name = cfgStruct.Name
	sms.ContentType = cfgStruct.ContentType
	return sms
}

func (sm *SmsMedia) Send(t *models.Transaction) error {
	var ps []PhoneNum
	var s SMSData

	for _, phone := range t.Receivers.Phone() {
		ps = append(ps, PhoneNum{Mobile: phone, Nationcode: 86})
	}
	s.Phones = ps
	s.Message = t.Message.Content
	b, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	body := bytes.NewBuffer([]byte(b))

	res, err := http.Post(sm.SmsUrl, sm.ContentType, body)
	if err != nil {
		log.Fatal(err)
		return err
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Info(string(result))

	return nil
}

func (sm *SmsMedia) Init() error {
	return nil
}

func (sm *SmsMedia) Redial() error {
	return nil
}

func (sm *SmsMedia) Close() error {
	return nil
}
