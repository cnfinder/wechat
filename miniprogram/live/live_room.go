package live

import (
	"encoding/json"
	"fmt"
	"github.com/cnfinder/wechat/context"
	"github.com/cnfinder/wechat/util"
)

const (
	createRoomURL     = "https://api.weixin.qq.com/wxaapi/broadcast/room/create"
	getLiveinfoURL = "https://api.weixin.qq.com/wxa/business/getliveinfo"
	getliveinfoReplayURL = "https://api.weixin.qq.com/wxa/business/getliveinfo"
	addgoodsURL = "https://api.weixin.qq.com/wxaapi/broadcast/room/addgoods"
)


//直播间管理
type LiveRoom struct {
	*context.Context
}


//init
func NewLiveRoom(context *context.Context) *LiveRoom {
	liveRoom := new(LiveRoom)
	liveRoom.Context = context
	return liveRoom
}


type Room struct {

	Name string `json:"name"`
	CoverImg string `json:"coverImg"`
	StartTime int64 `json:"startTime"`
	EndTime int64 `json:"endTime"`
	AnchorName string `json:"anchorName"`
	AnchorWechat string `json:"anchorWechat"`
	ShareImg string `json:"shareImg"`
	Type int `json:"type"`
	ScreenType int `json:"screenType"`
	CloseLike int `json:"closeLike"`
	CloseGoods int `json:"closeGoods"`
	CloseComment int `json:"closeComment"`
	CloseReplay int `json:"closeReplay"`
	CloseShare  int `json:"closeShare"`
	CloseKf int `json:"closeKf"`
}

type RoomRes struct {
	util.CommonError

	RoomId int `json:"roomId"`
}


// 创建直播间
// 接口说明：
// 调用此接口创建直播间，创建成功后将在直播间列表展示
// 需要先搜索小程序直播，实名认证后微信号才可以创建直播间

// 调用频率
// 调用额度：10000次/一天

// 请求方式
// POST
func (this *LiveRoom) CreateLiveRoom(room Room) (res RoomRes,err error){
	var accessToken string
	accessToken, err = this.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", createRoomURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri,room)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("CreateRoom error : errcode=%v , errmsg=%v", res.ErrCode, res.ErrMsg)
		return
	}
	return
}










type RoomInfoRes struct {

	util.CommonError

	RoomInfo []RoomInfo `json:"room_info"`
}

type RoomInfo struct {

	Name string `json:"name"`
	Roomid int `json:"roomid"`
	CoverImg string `json:"cover_img"`
	StartTime int64 `json:"start_time"`
	EndTime int64 `json:"end_time"`
	AnchorName string `json:"anchor_name"`
	ShareImg string `json:"shareImg"`
	Goods []Goods `json:"goods"`
	Total int `json:"total"`
	Live_status int `json:"live_status"`
}

type Goods struct {
	CoverImg string `json:"cover_img"`
	Url string `json:"url"`
	Price int `json:"price"`
	Name string `json:"name"`
}





// 获取直播间列表
//接口说明
//调用此接口获取直播间列表及直播间信息

//调用频率
//调用额度：100000次/一天（与获取回放接口共用次数）

//请求方式
//POST
func (this *LiveRoom) GetLiveinfoList(start,limit int) (res RoomInfoRes,err error){
	var accessToken string
	accessToken, err = this.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", getLiveinfoURL, accessToken)
	var response []byte

	mapdata:=map[string]int{
		"start":start,
		"limit":limit,
	}


	response, err = util.PostJSON(uri,mapdata)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("CreateRoom error : errcode=%v , errmsg=%v", res.ErrCode, res.ErrMsg)
		return
	}
	return
}










type RoomReplayRes struct {

	util.CommonError

	LiveReplay []LiveReplay `json:"live_replay"`
	Total int `json:"total"`
}

type LiveReplay struct {
	ExpireTime string `json:"expire_time"`
	CreateTime string `json:"create_time"`
	MediaUrl string `json:"media_url"`

}



// 获取直播间回放
// 接口说明
// 调用接口获取已结束直播间的回放源视频（一般在直播结束后10分钟内生成，源视频无评论等内容）

// 调用频率
// 调用额度：100000次/一天

// 请求方法
// POST

func (this *LiveRoom) GetLiveinfoReplayList(start,limit int,room_id int) (res RoomReplayRes,err error){
	var accessToken string
	accessToken, err = this.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", getliveinfoReplayURL, accessToken)
	var response []byte

	mapdata:=map[string]interface{}{
		"action":"get_replay",
		"room_id": room_id,
		"start":start,
		"limit":limit,

	}


	response, err = util.PostJSON(uri,mapdata)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("CreateRoom error : errcode=%v , errmsg=%v", res.ErrCode, res.ErrMsg)
		return
	}
	return
}







type AddGoodsRes struct {
	util.CommonError
}





// 直播间导入商品
// 接口说明
// 调用接口往指定直播间导入已入库的商品

// 调用频率
// 调用额度：10000次/一天
// 请求方法
// POST
//数组列表，可传入多个，里面填写 商品 ID

func (this *LiveRoom) AddGoods(ids []int,room_id int) (res AddGoodsRes,err error){
	var accessToken string
	accessToken, err = this.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", addgoodsURL, accessToken)
	var response []byte

	mapdata:=map[string]interface{}{
		"ids":ids,
		"roomId": room_id,

	}


	response, err = util.PostJSON(uri,mapdata)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("CreateRoom error : errcode=%v , errmsg=%v", res.ErrCode, res.ErrMsg)
		return
	}
	return
}



