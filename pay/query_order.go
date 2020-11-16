package pay

import (
	"encoding/xml"
	"fmt"
	"github.com/cnfinder/wechat/util"
)

var   queryOrderGateway = "https://api.mch.weixin.qq.com/pay/orderquery"


//   请求参数
type  QueryOrderRequest struct {
	Appid  string   `xml:"appid,omitempty" json:"appid,omitempty"` //   必传
	MchId  string   `xml:"mch_id,omitempty" json:"mch_id,omitempty"` //   必传
	TransactionId  string   `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`  // 二选一
	OutTradeNo    string      `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"` // 二选一
	NonceStr  string  `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`  //   必传
	Sign     string      `xml:"sign,omitempty" json:"sign,omitempty"` //   必传
	SignType  string   `xml:"sign_type,omitempty"`  // 非必传
}


//  接口返回
type   QueryOrderResponse struct {
	ReturnCode         string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg          string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	Appid              string `xml:"appid,omitempty" json:"appid,omitempty"`
	SubAppid           string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`
	MchId              string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	SubMchId           string `xml:"sub_mch_id,omitempty" json:"sub_mch_id,omitempty"`
	NonceStr           string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign               string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode         string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode            string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes         string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	DeviceInfo         string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	Openid             string `xml:"openid,omitempty" json:"openid,omitempty"`
	IsSubscribe        string `xml:"is_subscribe,omitempty" json:"is_subscribe,omitempty"`
	TradeType          string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`
	TradeState         string `xml:"trade_state,omitempty" json:"trade_state,omitempty"`
	BankType           string `xml:"bank_type,omitempty" json:"bank_type,omitempty"`
	TotalFee           string `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	SettlementTotalFee string `xml:"settlement_total_fee,omitempty" json:"settlement_total_fee,omitempty"`
	FeeType            string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`
	CashFee            string `xml:"cash_fee,omitempty" json:"cash_fee,omitempty"`
	CashFeeType        string `xml:"cash_fee_type,omitempty" json:"cash_fee_type,omitempty"`
	CouponFee          string `xml:"coupon_fee,omitempty" json:"coupon_fee,omitempty"`
	CouponCount        string `xml:"coupon_count,omitempty" json:"coupon_count,omitempty"`
	CouponType0        string `xml:"coupon_type_0,omitempty" json:"coupon_type_0,omitempty"`
	CouponType1        string `xml:"coupon_type_1,omitempty" json:"coupon_type_1,omitempty"`
	CouponType2        string `xml:"coupon_type_2,omitempty" json:"coupon_type_2,omitempty"`
	CouponId0          string `xml:"coupon_id_0,omitempty" json:"coupon_id_0,omitempty"`
	CouponId1          string `xml:"coupon_id_1,omitempty" json:"coupon_id_1,omitempty"`
	CouponId2          string `xml:"coupon_id_2,omitempty" json:"coupon_id_2,omitempty"`
	CouponFee0         string `xml:"coupon_fee_0,omitempty" json:"coupon_fee_0,omitempty"`
	CouponFee1         string `xml:"coupon_fee_1,omitempty" json:"coupon_fee_1,omitempty"`
	CouponFee2         string `xml:"coupon_fee_2,omitempty" json:"coupon_fee_2,omitempty"`
	TransactionId      string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo         string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	Attach             string `xml:"attach,omitempty" json:"attach,omitempty"`
	TimeEnd            string `xml:"time_end,omitempty" json:"time_end,omitempty"`
	TradeStateDesc     string `xml:"trade_state_desc,omitempty" json:"trade_state_desc,omitempty"`
}



func (pcf *Pay) Queryorder(p *QueryOrderRequest) (rsp QueryOrderResponse, err error){
	nonceStr := util.RandomStr(32)
	param := make(map[string]interface{})
	param["appid"] = pcf.AppID
	param["mch_id"] = pcf.PayMchID
	param["nonce_str"] = nonceStr
	param["out_trade_no"] = p.OutTradeNo
	param["sign_type"] = "MD5"

	bizKey := "&key=" + pcf.PayKey
	str := orderParam(param, bizKey)
	sign := util.MD5Sum(str)

	request := QueryOrderRequest{
		Appid:         pcf.AppID,
		MchId:         pcf.PayMchID,
		NonceStr:      nonceStr,
		Sign:          sign,
		SignType:      "MD5",
		OutTradeNo:    p.OutTradeNo,
	}

	rawRet, err := util.PostXML(queryOrderGateway, request)
	if err != nil {
		return
	}
	err = xml.Unmarshal(rawRet, &rsp)
	if err != nil {
		return
	}
	if rsp.ReturnCode == "SUCCESS" {
		if rsp.ResultCode == "SUCCESS" {
			err = nil
			return
		}
		err = fmt.Errorf("refund error, errcode=%s,errmsg=%s", rsp.ErrCode, rsp.ErrCodeDes)
		return
	}
	err = fmt.Errorf("[msg : xmlUnmarshalError] [rawReturn : %s] [params : %s] [sign : %s]",
		string(rawRet), str, sign)
	return


}

