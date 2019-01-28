package media

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/NioDevOps/courier/cfg"
	"github.com/NioDevOps/courier/models"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

const WORKWEIXINTYPE string = "work-weixin"

type WorkWeixinMedia struct {
	BaseMedia
	CorpId      string
	CorpSecret  string
	AgentId     int
	AccessToken string
}

type WorkWeixinText struct {
	Content string `json:"content"`
}

type WorkWeixinMessage struct {
	ToUser  string         `json:"touser"`
	MsgType string         `json:"msgtype"`
	AgentId int            `json:"agentid"`
	Text    WorkWeixinText `json:"text"`
}

type WorkWeixinResponse struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"erpires_in"`
}

func NewWorkWeixinMedia(cfgStruct *cfg.MediaCfgStruct) *WorkWeixinMedia {
	// wx := &WorkWeixinMedia{Name: cfgStruct.Name, CorpId: cfgStruct.CorpId, CorpSecret: cfgStruct.CorpSecret, AgentId: cfgStruct.AgentId}
	wx := &WorkWeixinMedia{CorpId: cfgStruct.CorpId, CorpSecret: cfgStruct.CorpSecret, AgentId: cfgStruct.AgentId}
	wx.Name = cfgStruct.Name
	return wx
}

// 40014不合法的access_token 41001缺少access_token参数 42001access_token已过期

func (wxm *WorkWeixinMedia) Send(t *models.Transaction) error {
	touser := strings.Join(t.Receivers.Addresses(), "|")
	text := WorkWeixinText{Content: t.Message.Content}
	message := WorkWeixinMessage{ToUser: touser, MsgType: "text", AgentId: wxm.AgentId, Text: text}

	b, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
	}
	body := bytes.NewBuffer([]byte(b))

	if wxm.AccessToken == "" {
		if err = wxm.UpdateToken(); err != nil {
			return err
		}
	}

	res, err := http.Post(fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", wxm.AccessToken), "application/json", body)
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
	var work_weixin_response WorkWeixinResponse
	if err := json.Unmarshal(result, &work_weixin_response); err == nil {
		if work_weixin_response.ErrCode == 0 {
			return nil
		} else if work_weixin_response.ErrCode == 40014 || work_weixin_response.ErrCode == 41001 || work_weixin_response.ErrCode == 42001 {
			err = wxm.UpdateToken()
			if err != nil {
				return err
			}
			_, err := http.Post(fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", wxm.AccessToken), "application/json", body)
			return err
		} else {
			return errors.New(work_weixin_response.ErrMsg)
		}
	} else {
		return err
	}

	return nil
}

func (wxm *WorkWeixinMedia) UpdateToken() error {
	uri := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", wxm.CorpId, wxm.CorpSecret)
	var work_weixin_response WorkWeixinResponse
	res, err := http.Get(uri)
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
	if err := json.Unmarshal(result, &work_weixin_response); err == nil {
		if work_weixin_response.ErrCode == 0 {
			wxm.AccessToken = work_weixin_response.AccessToken
		} else {
			return errors.New(work_weixin_response.ErrMsg)
		}
	} else {
		return err
	}

	return nil
}

func (wxm *WorkWeixinMedia) Init() error {
	return nil
}

func (wxm *WorkWeixinMedia) Redial() error {
	return nil
}

func (wxm *WorkWeixinMedia) Close() error {
	return nil
}
