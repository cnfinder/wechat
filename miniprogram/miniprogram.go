package miniprogram

import (
	"github.com/cnfinder/wechat/context"
	"github.com/cnfinder/wechat/miniprogram/live"
)

// MiniProgram struct extends context
type MiniProgram struct {
	*context.Context
}

// NewMiniProgram 实例化小程序接口
func NewMiniProgram(context *context.Context) *MiniProgram {
	miniProgram := new(MiniProgram)
	miniProgram.Context = context
	return miniProgram
}


// 获取直播间实列
func (this *MiniProgram) GetLiveRoom() *live.LiveRoom{
	return live.NewLiveRoom(this.Context)
}


// 获取直播商品管理实列
func (this *MiniProgram) GetLiveGoodsMgr() *live.LiveGoodsMgr{
	return live.NewLiveGoodsMgr(this.Context)
}


// 获取主播管理实例
func (this *MiniProgram) GetLiveAnchor()*live.LiveAnchor{
	return  live.NewLiveAnchor(this.Context)
}