package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	//"fmt"
	"model/wxp"
	"os"
	_ "router"
)

var appID string = "wx2ad19cfcaba9984f"
var appSecret string = "37e646b3a7ad514181bde49528873fcd"
var domain string = "api.weixin.qq.com"
var accessTokenURLFmt string = "https://%s/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:@/axiba?charset=utf8", 30, 30)
	orm.Debug = true
	orm.DebugLog = orm.NewLog(os.Stdout)
}

func main() {

	wxp.Init(appID, appSecret, domain, accessTokenURLFmt)
	wxp.WXPIns.AccessTokenIns.SetAccURLFmt(accessTokenURLFmt)

	go wxp.WXPIns.WatchDog()

	beego.Run(":3000")
}
