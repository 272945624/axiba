package wxp

import (
	"fmt"
	"time"
)

var WXPIns WXP

type WXP struct {
	appID          string
	appSecret      string
	AccessTokenIns WXPAccessToken
	Domain         string
}

func Init(id, secret, domain, accessTokenURLFmt string) {
	WXPIns.appID = id
	WXPIns.appSecret = secret
	WXPIns.Domain = domain
	WXPIns.AccessTokenIns.AccUrlFmt = accessTokenURLFmt
}

func (wxp *WXP) getAppID() string {
	return wxp.appID
}

func (wxp *WXP) getAppSecret() string {
	return wxp.appSecret
}

func (wxp *WXP) GetAccessToken() string {
	return wxp.AccessTokenIns.getAccessToken()
}

func (wxp *WXP) WatchDog() {
	//get accesstoken url

	//'''https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wx2ad19cfcaba9984f&secret=37e646b3a7ad514181bde49528873fcd'''
	wxp.AccessTokenIns.UpdateAccessToken()

	tik := time.NewTicker(wxp.AccessTokenIns.ExpireTime)

	for {
		select {
		case now := <-tik.C:
			//update access token

			fmt.Println("DEBUG: ", now, " Update access token.")
			err := wxp.AccessTokenIns.UpdateAccessToken()
			if err != nil {
				wxp.AccessTokenIns.ExpireTime = time.Duration(60 * time.Second)
			}
			tik.Stop()
			tik = time.NewTicker(wxp.AccessTokenIns.ExpireTime)

		}
	}
}
