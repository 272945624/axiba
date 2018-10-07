package controllers

import (
	"crypto/sha1"
	//	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"io/ioutil"
	"sort"
)

type WeiXinPlatform struct {
	beego.Controller
}

func (req *WeiXinPlatform) Get() {
	var list []string
	var signature, timestamp, nonce, echostr, token string

	h := sha1.New()

	list = make([]string, 3, 3)

	signature = req.GetString("signature")
	if signature == "" {
		goto ferr
	}
	timestamp = req.GetString("timestamp")
	if timestamp == "" {
		goto ferr
	}
	nonce = req.GetString("nonce")
	if nonce == "" {
		goto ferr
	}
	echostr = req.GetString("echostr")

	if echostr == "" {
		goto ferr
	}

	token = "563855"

	list[0] = token
	list[1] = timestamp
	list[2] = nonce

	sort.Slice(list,
		func(i, j int) bool {
			return list[i] < list[j]
		})

	for _, l := range list {
		io.WriteString(h, l)
	}

	req.Ctx.WriteString(echostr)
	return

ferr:
	req.Ctx.WriteString("error")

}

func (req *WeiXinPlatform) Post() {
	ret := "<xml><ToUserName><![CDATA[ox2gRwzcX6WuAFMq1HHsyzAThHRI]]></ToUserName> <FromUserName><![CDATA[gh_aeb14d0007d0]]></FromUserName> <CreateTime>1538753493</CreateTime> <MsgType><![CDATA[text]]></MsgType> <Content><![CDATA[quququ]]></Content> <MsgId>6608901590228999999</MsgId> </xml>"

	fmt.Println(req.Ctx.Input.Method())
	fmt.Println(req.Ctx.Input.Protocol())
	fmt.Println(req.Ctx.Request.ContentLength)
	body, err := ioutil.ReadAll(req.Ctx.Request.Body)
	if err != nil {
		goto ferr
	}

	fmt.Println("body", string(body))
	req.Ctx.WriteString(ret)

	return
ferr:
	req.Ctx.WriteString("success")
}
