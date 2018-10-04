package main

import (
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	. "model/database"
	"os"
	_ "router"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:@/axiba?charset=utf8", 30, 30)
	orm.Debug = true
	orm.DebugLog = orm.NewLog(os.Stdout)
}

func main() {
	beego.Run()
	user := new(UserAdmin)

	user.UserName = "pppp"
	user.Id = 2

	o := orm.NewOrm()
	o.Using("default")
	//	fmt.Println(user.Update(o, user, "user_name"))
	o.QueryTable(user.TableName()).Filter("user_name", "guoheng2").Filter("id", 4).Delete()
}
