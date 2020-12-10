package miniprogram

import (
	"encoding/json"
	"fmt"

	"github.com/cnfinder/wechat/context"
	"github.com/cnfinder/wechat/util"
)

/* 微信小程序发送订阅消息
 * 文档地址：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html
 */

const (
	subMessageURL = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=%s"
)

//Subscribe 订阅消息
type Subscribe struct {
	*context.Context
}

// Message 订阅消息请求参数
type Message struct {
	ToUser           string               `json:"touser"`            //必选，接收者（用户）的 openid
	TemplateID       string               `json:"template_id"`       //必选，所需下发的订阅模板id
	Page             string               `json:"page"`              //可选，点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
	Data             map[string]*DataItem `json:"data"`              //必选, 模板内容
	MiniprogramState string               `json:"miniprogram_state"` //可选，跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
	Lang             string               `json:"lang"`              //入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
}

//DataItem 模版内某个 .DATA 的值
type DataItem struct {
	Value string `json:"value"`
}

//subscribeMessage 订阅消息返回结果
type ResSubMessage struct {
	util.CommonError
}

// subMessage 开始发送消息
func (this *Subscribe) SubMessage(msg *Message) (result ResSubMessage, err error) {

	// d := DataItem{}
	// d.Value = "dddd"

	// s := Message{}
	// s.Data = d

	accessToken, err := this.GetAccessToken()
	if err != nil {
		return
	}
	urlStr := fmt.Sprintf(accessToken)

	var response []byte
	response, err = util.PostJSON(urlStr, msg)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}

	if result.ErrCode != 0 {
		err = fmt.Errorf("SubscribeMessage error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}
