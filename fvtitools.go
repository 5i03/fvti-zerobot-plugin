package fvtitools

// import 用来放置你所需要导入的东西, 萌新推荐使用vscode, 它会帮你干很多事
import (
	"math/rand"
	"time"
	"/fvtitools/config"
	"/plugin/health"
	log "github.com/sirupsen/logrus"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	"github.com/FloatTech/zbputils/ctxext"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/extension/shell"
	"github.com/wdvxdr1123/ZeroBot/message"
)

var examplelimit = ctxext.NewLimiterManager(time.Second*10, 1)
func init() {
	u := config.GetUser(ctx.Event.UserID)
	if u.QqNumber != ctx.Event.UserID {
		return
	}
	arguments := shell.Parse(ctx.State["args"].(string))
	if len(arguments) != 1 {
		ctx.Send(message.Text("格式错误,正确格式\n/修改密码 114514"))
		return
	}
	u.Pass = arguments[0]
	if health.FvtiLogin(u) != "" {
		config.SaveUser(u)
		ctx.Send(message.Text("绑定好了，拉机器人入群吧"))
		return
	}
	ctx.Send(message.Text("密码错误，或者学校服务器抽风了"))

	engine := control.Register("fvtitools", &ctrl.Options[*zero.Ctx]{
	// 控制插件是否默认启用 true为默认不启用 false反之
	DisableOnDefault: false,
	// 插件的帮助 管理员发送 /用法  可见
	Help: "- fvtitools 插件的帮助"+
		"我要帮助 - 查看帮助"+
		"健康鉴权 - 手动登陆系统以更新令牌"+
		"健康改密 - 登记健康密码"+	
		"健康账户 - 查看用户信息"+
		"申出校码 - 申请出校入校码"+
		"清无效码 - 清理无效出校码"	,
	// 插件的数据存放路径, 分为公共和私有, 都会在/data下创建目录, 公有需要首字母大写, 私有需要首字母小写
	PublicDataFolder: "fvtitools",
	// PrivateDataFolder: "example",		// 避免问题所以注释了
	// 自定义插件开启时的回复
	OnEnable: func(ctx *zero.Ctx) {
		ctx.Send("插件已启用")
	},
	// 自定义插件关闭时的回复
	OnDisable: func(ctx *zero.Ctx) {
		ctx.Send("插件已禁用")
	},
})
}


