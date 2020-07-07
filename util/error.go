package util

import (
	"encoding/json"
	"fmt"
	"reflect"
)

/*
-1：系统错误

1：未创建直播间

1003：商品id不存在

47001：入参格式不符合规范

200002：入参错误

300001：禁止创建/更新商品 或 禁止编辑&更新房间

300002：名称长度不符合规则

300006：图片上传失败（如：mediaID过期）

300022：此房间号不存在

300023：房间状态 拦截（当前房间状态不允许此操作）

300024：商品不存在

300025：商品审核未通过

300026：房间商品数量已经满额

300027：导入商品失败

300028：房间名称违规

300029：主播昵称违规

300030：主播微信号不合法

300031：直播间封面图不合规

300032：直播间分享图违规

300033：添加商品超过直播间上限

300034：主播微信昵称长度不符合要求

300035：主播微信号不存在

300036: 主播微信号未实名认证

9410000: 直播间列表为空

9410001: 获取房间失败

9410002: 获取商品失败

9410003: 获取回放失败

 */


// CommonError 微信返回的通用错误json
type CommonError struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// DecodeWithCommonError 将返回值按照CommonError解析
func DecodeWithCommonError(response []byte, apiName string) (err error) {
	var commError CommonError
	err = json.Unmarshal(response, &commError)
	if err != nil {
		return
	}
	if commError.ErrCode != 0 {
		return fmt.Errorf("%s Error , errcode=%d , errmsg=%s", apiName, commError.ErrCode, commError.ErrMsg)
	}
	return nil
}

// DecodeWithError 将返回值按照解析
func DecodeWithError(response []byte, obj interface{}, apiName string) error {
	err := json.Unmarshal(response, obj)
	if err != nil {
		return fmt.Errorf("json Unmarshal Error, err=%v", err)
	}
	responseObj := reflect.ValueOf(obj)
	if !responseObj.IsValid() {
		return fmt.Errorf("obj is invalid")
	}
	commonError := responseObj.Elem().FieldByName("CommonError")
	if !commonError.IsValid() || commonError.Kind() != reflect.Struct {
		return fmt.Errorf("commonError is invalid or not struct")
	}
	errCode := commonError.FieldByName("ErrCode")
	errMsg := commonError.FieldByName("ErrMsg")
	if !errCode.IsValid() || !errMsg.IsValid() {
		return fmt.Errorf("errcode or errmsg is invalid")
	}
	if errCode.Int() != 0 {
		return fmt.Errorf("%s Error , errcode=%d , errmsg=%s", apiName, errCode.Int(), errMsg.String())
	}
	return nil
}
