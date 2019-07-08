package miniprogram

import (
	"encoding/json"
	"fmt"
	"github.com/imokyou/wechat/util"
)

var (
	MessageApi = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token="
)

type MessageData struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type MessageRequest struct {
	Touser          string                 `json:"touser"`
	TemplateId      string                 `json:"template_id"`
	Page            string                 `json:"page"`
	FormId          string                 `json:"form_id"`
	Data            map[string]MessageData `json:"data"`
	Color           string                 `json:"color"`
	EmphasisKeyword string                 `json:"emphasis_keyword"`
}

type MessageResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	// MsgId int `json:"msgid"`
}

func (wxa *MiniProgram) SendMessage(req *MessageRequest) (*MessageResponse, error) {
	ret := &MessageResponse{}

	accessToken, err := wxa.GetAccessToken()
	if err != nil {
		return ret, err
	}
	api := MessageApi + accessToken
	response, _, err := util.PostJSONWithRespContentType(api, req)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(response, &ret)
	if err != nil {
		return ret, err
	}

	if err == nil && ret.ErrCode != 0 {
		err = fmt.Errorf("send message error : errcode=%v , errmsg=%v", ret.ErrCode, ret.ErrMsg)
		return ret, err
	}

	return ret, nil
}
