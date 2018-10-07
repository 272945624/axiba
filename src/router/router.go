package router

import (
	"controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/admin/user", &controllers.UserAdminController{})
	beego.Router("/wx", &controllers.WeiXinPlatform{})
}
