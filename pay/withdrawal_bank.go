package pay



//   TODO   未验证

import (
	"encoding/xml"
	"fmt"
	"github.com/cnfinder/wechat/util"
)

var WithdrawalGateway_bank = "https://api.mch.weixin.qq.com/mmpaysptrans/pay_bank"


//◆ 单商户日限额——单日10万元
//
//◆ 单次限额——单次2万元
//
//◆ 单商户给同一银行卡单日限额——单日2万元
// 企业付款到银行卡 请求参数
type WithdrawalRequest_bank struct {
	Mchid            string `json:"mchid" xml:"mchid"`                       //  必传       微信支付分配 的 商户号
	Nonce_str        string `json:"nonce_str" xml:"nonce_str"`               //  必传       随机字符串  不长于 32 位
	Enc_bank_no      string `json:"enc_bank_no" xml:"enc_bank_no"`           //  必传       收款方银行卡号 （采用标准RSA算法，公钥由微信侧提供 详见 https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_7）
	Enc_true_name    string `xml:"enc_true_name"  json:"enc_true_name"`      //  必传       收款方用户名 （采用标准RSA算法，公钥由微信侧提供 详见 https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_7）
	Bank_code        string `xml:"bank_code" json:"bank_code"`               //  必传       银行卡所在开户行编号 （银行卡所在开户行编号,详见https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_4）
	Amount           int    `xml:"amount" json:"amount"`  					 //  必传       付款金额: RMB 分 （支付总额， 不含手续费） 注  大于 0  的整数
	Desc             string `xml:"desc" json:"desc"`					 	 //  可传       企业付款到银行卡付款说明，即订单备注（UTF8编码，允许100个字符以内）
	RootCa           string   // ca 证书
}


//企业付款到银行卡 接口返回
type WithdrawalResponse_bank struct {
	ReturnCode          string `xml:"return_code"`
	ReturnMsg           string `xml:"return_msg"`     //  以上  默认返回
	ResultCode          string `xml:"result_code,omitempty"`
	ErrCode             string `xml:"err_code,omitempty"`
	ErrCodeDes          string `xml:"err_code_des,omitempty"`     // 以上字段在return_code为SUCCESS的时候有返回
	Mch_id              string `xml:"mch_id,omitempty"`
	Partner_trade_no    string `xml:"partner_trade_no，omitempty"`
	Amount              int    `xml:"amount,omitempty"`
	NonceStr            string `xml:"nonce_str,omitempty"`
	Sign                string `xml:"sign,omitempty"`
	Payment_no          string `xml:"payment_no，omitempty"`
	Cmms_amt            int    `xml:"cmms_amt，omitempty"`
}



func (pcf *Pay) WithDrawal_bank (req *WithdrawalRequest_bank)(rsp *WithdrawalResponse_bank, err error){
	nonceStr := util.RandomStr(32)
	param := make(map[string]interface{})
	param["mch_appid"] = pcf.AppID
	param["mchid"] = pcf.PayMchID
	param["nonce_str"] = nonceStr
	param["enc_bank_no"] = req.Enc_bank_no
	param["enc_true_name"] = req.Enc_true_name
	param["bank_code"] = req.Bank_code
	param["amount"] = req.Amount
	param["desc"] = req.Desc

	bizKey := "&key=" + pcf.PayKey
	str := orderParam(param, bizKey)
	sign := util.MD5Sum(str)

	request := WithdrawalRequest_bank{
		Mchid:pcf.PayMchID,
		Nonce_str:nonceStr,
		Enc_true_name: req.Enc_true_name,
		Enc_bank_no:req.Enc_bank_no,
		Bank_code:req.Bank_code,
		Amount:req.Amount,
		Desc:req.Desc,
	}
	rawRet ,err := util.PostXMLWithTLS(WithdrawalGateway_wxbalance,request,req.RootCa,req.Mchid)
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
		err = fmt.Errorf("withdrawal_bank error, errcode=%s,errmsg=%s", rsp.ErrCode, rsp.ErrCodeDes)
		return
	}
	err = fmt.Errorf("[msg : xmlUnmarshalError] [rawReturn : %s] [params : %s] [sign : %s]",
		string(rawRet), str, sign)
	return
}
