package live

import (
	"encoding/json"
	"fmt"
	"github.com/cnfinder/wechat/context"
	"github.com/cnfinder/wechat/util"
)


const (
	goodsAddURL     = "https://api.weixin.qq.com/wxaapi/broadcast/goods/add"
	resetauditURL = "https://api.weixin.qq.com/wxaapi/broadcast/goods/resetaudit"
	goodsAuditURL = "https://api.weixin.qq.com/wxaapi/broadcast/goods/audit"
	goodsDeleteURL = "https://api.weixin.qq.com/wxaapi/broadcast/goods/delete"
	goodsUpdateURL = "https://api.weixin.qq.com/wxaapi/broadcast/goods/update"
	getgoodswarehouseURL = "https://api.weixin.qq.com/wxa/business/getgoodswarehouse"
	getapprovedURL = "https://api.weixin.qq.com/wxaapi/broadcast/goods/getapproved"
)


//直播间商品管理
type LiveGoodsMgr struct {
	*context.Context
}



//init
func NewLiveGoodsMgr(context *context.Context) *LiveGoodsMgr {
	liveGoodsMgr := new(LiveGoodsMgr)
	liveGoodsMgr.Context = context
	return liveGoodsMgr
}


type GoodsInfoRequest struct {
	GoodsInfo GoodsInfo `json:"goodsInfo"`
}

type GoodsInfoRes struct {
	util.CommonError

	GoodsId int `json:"goodsId"`
	AuditId int `json:"auditId"`

}


type GoodsInfo struct {
	CoverImgUrl string `json:"coverImgUrl"`

	Name string `json:"name"`
	PriceType int `json:"priceType"`
	Price float64 `json:"price"`
	Price2 float64 `json:"price2"`
	
	Url string `json:"url"`
	GoodsId int `json:"goodsId"`
}


// 商品添加并提审
//接口说明
//调用此接口上传并提审需要直播的商品信息，审核通过后商品录入【小程序直播】商品库

// 注意：开发者必须保存【商品ID】与【审核单ID】，如果丢失，则无法调用其他相关接口

// 调用频率
// 调用额度：500次/一天
func (this *LiveGoodsMgr) GoodsAdd(req GoodsInfoRequest) (res GoodsInfoRes,err error){
	var accessToken string
	accessToken, err = this.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", goodsAddURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri,req)
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












// 撤回审核
//接口说明
//调用此接口，可撤回直播商品的提审申请，消耗的提审次数不返还


// 调用频率
// 调用额度：500次/一天
func (this *LiveGoodsMgr) GoodsResetaudit(audit int,goodsId int) (res util.CommonError,err error){
	var accessToken string
	accessToken, err = this.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", resetauditURL, accessToken)
	var response []byte

	mapdata:=map[string]int{
		"auditId":audit,
		"goodsId":goodsId,
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







type GoodsAuditRes struct {
	util.CommonError
	AuditId int `json:"auditId"`
}



// 重新提交审核
//接口说明
//调用此接口，可撤回直播商品的提审申请，消耗的提审次数不返还


// 调用频率
// 调用额度：500次/一天
func (this *LiveGoodsMgr) GoodsAudit(goodsId int) (res util.CommonError,err error){
	var accessToken string
	accessToken, err = this.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", goodsAuditURL, accessToken)
	var response []byte

	mapdata:=map[string]int{
		"goodsId":goodsId,
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





// 调用频率
// 调用额度：500次/一天
func (this *LiveGoodsMgr) GoodsDelete(goodsId int) (res util.CommonError,err error){
	var accessToken string
	accessToken, err = this.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", goodsDeleteURL, accessToken)
	var response []byte

	mapdata:=map[string]int{
		"goodsId":goodsId,
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













// 更新商品
// 调用此接口可以更新商品信息，审核通过的商品仅允许更新价格类型与价格，审核中的商品不允许更新，未审核的商品允许更新所有字段， 只传入需要更新的字段。

// 调用额度：1000次/一天
func (this *LiveGoodsMgr) GoodsUpdate(goodsInfo GoodsInfo) (res util.CommonError,err error){
	var accessToken string
	accessToken, err = this.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", goodsUpdateURL, accessToken)
	var response []byte


	response, err = util.PostJSON(uri,goodsInfo)
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




type GoodswarehouseRes struct {

	util.CommonError

	Goods []GoodsWare `json:"goods"`
	Total int `json:"total"`

}



type GoodsWare struct {
	CoverImgUrl string `json:"cover_img_url"`

	Name string `json:"name"`
	PriceType int `json:"priceType"`
	Price float64 `json:"price"`
	Price2 float64 `json:"price2"`

	Url string `json:"url"`
	GoodsId int `json:"goods_id"`
	AuditStatus int `json:"audit_status"`
	Third_party_tag int `json:"third_party_tag"`
}




// 获取商品状态
//接口说明
// 调用此接口可获取商品的信息与审核状态

// 调用额度：1000次/一天
func (this *LiveGoodsMgr) GetGoodswarehouse(goods_ids []int) (res GoodswarehouseRes,err error){
	var accessToken string
	accessToken, err = this.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", getgoodswarehouseURL, accessToken)
	var response []byte

	mapdata:=map[string][]int{
		"goods_ids":goods_ids,
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









type GoodsListRes struct {

	util.CommonError

	Goods []GoodsWare2 `json:"goods"`
	Total int `json:"total"`

}



type GoodsWare2 struct {
	CoverImgUrl string `json:"coverImgUrl"`

	Name string `json:"name"`
	PriceType int `json:"priceType"`
	Price float64 `json:"price"`
	Price2 float64 `json:"price2"`

	Url string `json:"url"`
	GoodsId int `json:"goods_id"`
	 int `json:"audit_status"`
	Third_party_tag int `json:"thirdPartyTag"`
}




// 获取商品列表
//接口说明
// 调用此接口可获取商品列表

// 调用额度：10000次/一天
func (this *LiveGoodsMgr) Getapproved(offset int,limit int,status int) (res GoodsListRes,err error){
	var accessToken string
	accessToken, err = this.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", getapprovedURL, accessToken)
	var response []byte

	mapdata:=map[string]interface{}{
		"offset":offset,
		"limit":limit,
		"status":status,
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
