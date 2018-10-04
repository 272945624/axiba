package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	. "model/database"
)

type UserAdminController struct {
	beego.Controller
}

func (req *UserAdminController) Get() {

	var users []*UserAdmin

	name := req.GetString("name")

	o := orm.NewOrm()
	o.Using("default")

	_, err := o.QueryTable(&UserAdmin{}).Filter("user_name", name).All(&users)

	if err != nil {
		req.Ctx.WriteString(err.Error())
		return
	}

	for _, u := range users {
		fmt.Println(u)
	}

	ms, err := json.Marshal(users)

	if err != nil {
		fmt.Printf("ERROR: %s", err.Error())
		return
	}

	fmt.Println(string(ms))

	req.Ctx.WriteString(string(ms))
}

func (req *UserAdminController) Put() {
	req.Ctx.WriteString("put hello.")
}

func (req *UserAdminController) Post() {
	req.Ctx.WriteString("post hello.")
}

func (req *UserAdminController) Delete() {
	req.Ctx.WriteString("delete hello.")
}
