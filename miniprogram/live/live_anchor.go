package live

import (
	"encoding/json"
	"fmt"
	"github.com/cnfinder/wechat/context"
	"github.com/cnfinder/wechat/util"
)

const (
	addRoleUrl = "https://api.weixin.qq.com/wxaapi/broadcast/role/addrole?access_token="
    delRoleUrl = "https://api.weixin.qq.com/wxaapi/broadcast/role/deleterole?access_token="
    queryListUrl = "https://api.weixin.qq.com/wxaapi/broadcast/role/getrolelist?access_token="
)

var ROLE_ERR_MSG = map[int]string{
	400001: "微信号不合规",
	400002: "微信号需要实名认证",
	400003: "添加角色达到上限",
	400004: "重复添加角色",
	400005:" 主播角色删除失败，该主播存在未开播的直播间",
}

//主播管理
type LiveAnchor struct {
	*context.Context
}


func NewLiveAnchor(context  *context.Context)*LiveAnchor{
	liveAnchor := new(LiveAnchor)
	liveAnchor.Context = context
	return liveAnchor
}



/*
err msg

-1：系统错误

400001: 微信号不合规

400002: 微信号需要实名认证，仅设置主播角色时可能出现

400003: 添加角色达到上限（管理员10个，运营者500个，主播500个）

400004: 重复添加角色

400005: 主播角色删除失败，该主播存在未开播的直播间
*/



type  RoleReq struct {
	Username string `json:"username" form:"username"`
	Role   int  `json:"role" form:"role"`
}


/*
{
    role: 1, // 取值 [-1-所有成员， 0-超级管理员，1-管理员，2-主播，3-运营者]
    offset: 0, // 起始偏移量
    limit: 10, // 查询个数，最大30，默认10
    keyword: 'test_1' // 搜索的微信号，不传返回全部
}
*/


type RoleListReq struct {
	Role int `json:"role" form:"role"`
	Offset int `json:"offset" form:"offset"`
	Limit    int `json:"limit" form:"limit"`
	Keyword   string `json:"keyword" form:"keyword"`
}

type  RoleRes struct {
	Errcode  int  `json:"errcode" form:"errcode"`
}


/*
        "headingimg": "http://wx.qlogo.cn/mmhead/Q3auHgzwzM5jBhFwrHoeoaxTlhP9YzlVica7wu6lZLnGreKAj7CVicA/0", // 头像
        "nickname": "test1", // 昵称
        "openid": "o7esq5MvImF2SEm7OHYohausj2o",
        "roleList": [2, 3], // 具有的身份，[0-超级管理员，1-管理员，2-主播，3-运营者]
        "updateTimestamp": "1600340080", // 更新时间
        "username": "o0****0o", //脱敏微信号
 */


type  RoleList struct {
	Headingimg    string `json:"headingimg" form:"headingimg"`
	Nickname       string `json:"nickname" form:"nickname"`
	Openid        string `json:"openid" from:"openid"`
	UpdateTimestamp  []int `json:"updateTimestamp" form:"updateTimestamp"`
	Username      string `json:"username" form:"username"`
}


/*
{
    "errcode": 0,
    "total" : 1, // 总个数
    "list": [{
        "headingimg": "http://wx.qlogo.cn/mmhead/Q3auHgzwzM5jBhFwrHoeoaxTlhP9YzlVica7wu6lZLnGreKAj7CVicA/0", // 头像
        "nickname": "test1", // 昵称
        "openid": "o7esq5MvImF2SEm7OHYohausj2o",
        "roleList": [2, 3], // 具有的身份，[0-超级管理员，1-管理员，2-主播，3-运营者]
        "updateTimestamp": "1600340080", // 更新时间
        "username": "o0****0o", //脱敏微信号
    }]
}
*/

type RoleListRes struct {
     Errcode   int  `json:"errcode" form:"errcode"`
     Total  int `json:"total" form:"total"`
     List    []RoleList `json:"list" form:"list"`
}



/*
接口说明：
调用此接口设置小程序直播成员的管理员、运营者和主播角色

调用频率
调用额度：10000次/一天

请求方式
POST
*/


func (this *LiveAnchor) AddRole(req *RoleReq)(int,error){
	var accessToken string
	accessToken, err := this.GetAccessToken()
	if err != nil {
		return   -1,err
	}

	uri := fmt.Sprintf("%s?access_token=%s", addRoleUrl, accessToken)
	response, err := util.PostJSON(uri,req)
	if err != nil {
		return -1,err
	}

	res := RoleRes{}
	err = json.Unmarshal(response, &res)
	if err != nil {
		return   -1,err
	}
	return  res.Errcode,nil
}


/*
接口说明：
调用此接口移除小程序直播成员的管理员、运营者和主播角色

调用频率
调用额度：10000次/一天

请求方式
POST


*/

func (this *LiveAnchor) DelRole(req *RoleReq)(int,error){
	var accessToken string
	accessToken, err := this.GetAccessToken()
	if err != nil {
		return  -1,err
	}

	uri := fmt.Sprintf("%s?access_token=%s", delRoleUrl, accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return -1,err
	}

	res := RoleRes{}
	err = json.Unmarshal(response, &res)
	if err != nil {
		return -1,err
	}
	return res.Errcode,nil
}


/*
接口说明：
调用此接口查询小程序直播成员列表

调用频率
调用额度：10000次/一天

请求方式
GET

*/


func (this *LiveAnchor) QueryRoleList(req  *RoleListReq)(*RoleListRes,error){
	var accessToken string
	accessToken, err := this.GetAccessToken()
	if err != nil {
		return  nil , err
	}

	uri := fmt.Sprintf("%s?access_token=%s&role=%d&offset=%d&limit=%d&keyword=%s", queryListUrl, accessToken,req.Role,req.Offset,req.Limit,req.Keyword)

	response, err := util.HTTPGet(uri)
	if err != nil {
		return  nil , err
	}
	res := RoleListRes{}
	err = json.Unmarshal(response, &res)
	return  &res , nil
}