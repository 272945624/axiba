package router

import (
	"controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/admin/user", &controllers.UserAdminController{})
}
