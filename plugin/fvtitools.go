package fvtitools

import (
		"math/rand"
		"time"
	
		ctrl "github.com/FloatTech/zbpctrl"
		"github.com/FloatTech/zbputils/control"
		"github.com/FloatTech/zbputils/ctxext"
		zero "github.com/wdvxdr1123/ZeroBot"
		"github.com/wdvxdr1123/ZeroBot/message"
		)

// FVTI is a function that returns a function that can be used to send a message to the group
var examplelimit = ctxext.NewLimiterManager(time.Second*10, 1)
func init(){
	engine := control.Register("fvtitools", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Help: "- fvtitools 插件的帮助",
		PublicDataFolder: "fvtitools",
		OnEnable: func(ctx *zero.Ctx) {
			ctx.Send("插件已启用")
		},
		// 自定义插件关闭时的回复
		OnDisable: func(ctx *zero.Ctx) {
			ctx.Send("插件已禁用")
		},


	}
	// 注册一个命令	)
	engine.OnFullMatchGroup([]string{"fvtitools"}, zero.AdminPermission).SetBlock(true).SetPriority(20).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Text("fvtitools"))
	}
	