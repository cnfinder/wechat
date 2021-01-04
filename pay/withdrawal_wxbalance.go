package pay


//   TODO   未验证
import (
	"encoding/xml"
	"fmt"
	"github.com/cnfinder/wechat/util"
)

var WithdrawalGateway_wxbalance = "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers"


// 企业付款到零钱 请求参数
type WithdrawalRequest_wxbalance struct {
	Mch_appid        string `json:"mch_appid"  xml:"mch_appid"`              //  必传       申请商户号的 app_id  或  商户号 绑定的 app_id
	Mchid            string `json:"mchid" xml:"mchid"`                       //  必传       微信支付分配 的 商户号
	Device_info      string `json:"device_info" xml:"device_info"`           //  可传       微信支付分配 的 终端设备号
	Nonce_str        string `json:"nonce_str" xml:"nonce_str"`               //  必传       随机字符串  不长于 32 位
	Sign             string `json:"sign" xml:"sign"`                         //  可传       签名  详见 签名算法  （https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=4_3）
	Partner_trade_no string `json:"partner_trade_no" xml:"partner_trade_no"` //  必传       商户订单号 需保持 唯一性 （只能是 字母 或者 数字 ， 不能 包含 其他 字符）
	Openid           string `xml:"openid" json:"openid"`                     //  必传       商户 app_id 下 ，某用户的 openid
	Check_name       string `xml:"check_name" json:"check_name"`             //  必传       NO_CHECK: 不校验真实性     FORCE_CHECK:强校验真实姓名
	Re_user_name     string `xml:"re_user_name" json:"re_user_name"`         //  可传       收款用户真实姓名
	Amount           int    `xml:"amount" json:"amount"`                     //  必传       企业付款金额 , 单位为 分
	Desc             string `xml:"desc" json:"desc"`                         //  必传              企业付款备注
	Spbill_create_ip string `xml:"spbill_create_ip" json:"spbill_create_ip"` //  可传    该ip同在商户平台设置的ip白名单中的IP没有关联，该ip 可传用户端或者 服务端的ip
	RootCa           string         // ca证书
	CaKey string // 证书key
}


//企业付款到零钱 接口返回
type WithdrawalResponse_wxbalance struct {
	ReturnCode          string `xml:"return_code"`
	ReturnMsg           string `xml:"return_msg"`     //  以上  默认返回
	Mch_appid           string `xml:"mch_appid,omitempty"`
	Mchid               string `xml:"mchid,omitempty"`
	Device_info         string `xml:"device_info,omitempty"`
	NonceStr            string `xml:"nonce_str,omitempty"`
	ResultCode          string `xml:"result_code,omitempty"`
	ErrCode             string `xml:"err_code,omitempty"`
	ErrCodeDes          string `xml:"err_code_des,omitempty"`     // 以上字段在return_code为SUCCESS的时候有返回
	Partner_trade_no    string `xml:"partner_trade_no，omitempty"`
	Payment_no          string `xml:"payment_no，omitempty"`
	Payment_time          string `xml:"payment_time，omitempty"`
}





func (pcf *Pay) Withdrawal_wxbalance(req  *WithdrawalRequest_wxbalance)(rsp *WithdrawalResponse_wxbalance, err error){
	nonceStr := util.RandomStr(32)
	param := make(map[string]interface{})
	param["mch_appid"] = pcf.AppID
	param["mchid"] = pcf.PayMchID
	param["nonce_str"] = nonceStr
	param["partner_trade_no"] = req.Partner_trade_no
	param["openid"] = req.Openid
	param["check_name"] = req.Check_name
	param["amount"] = req.Amount
	param["desc"] = req.Desc
	param["sign"] = req.Sign

	bizKey := "&key=" + pcf.PayKey
	str := orderParam(param, bizKey)
	sign := util.MD5Sum(str)

	request := WithdrawalRequest_wxbalance{
		Mch_appid:pcf.AppID,
		Mchid:pcf.PayMchID,
		Nonce_str:nonceStr,
		Sign:sign,
		Partner_trade_no:req.Partner_trade_no,
		Openid:req.Openid,
		Check_name:req.Check_name,
		Amount:req.Amount,
		Desc:req.Desc,
	}
	rawRet ,err := util.PostXMLWithTLS(WithdrawalGateway_wxbalance,request,req.RootCa,req.CaKey)
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
		err = fmt.Errorf("withdrawal_wxbalance error, errcode=%s,errmsg=%s", rsp.ErrCode, rsp.ErrCodeDes)
		return
	}
	err = fmt.Errorf("[msg : xmlUnmarshalError] [rawReturn : %s] [params : %s] [sign : %s]",
		string(rawRet), str, sign)
	return
}
