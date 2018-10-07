package wxp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type accessTokenMsg struct {
	Token   string `json:"access_token"`
	Expires int64  `json:"expires_in"`
}

type WXPAccessToken struct {
	accessToken string
	Timeoutd    bool
	ExpireTime  time.Duration
	AccUrlFmt   string
	RWMut       sync.RWMutex
}

func (wxpat *WXPAccessToken) getAccessToken() string {

	defer wxpat.RWMut.RUnlock()

	var accToken string

	wxpat.RWMut.RLock()

	accToken = wxpat.accessToken

	return accToken
}

func (wxpat *WXPAccessToken) setAccessToken(accessToken string) {
	defer wxpat.RWMut.Unlock()

	wxpat.RWMut.Lock()

	wxpat.accessToken = accessToken
}

func (wxpat *WXPAccessToken) getAccessTokenUrl() string {
	return fmt.Sprintf(wxpat.AccUrlFmt, WXPIns.Domain, WXPIns.appID, WXPIns.appSecret)
}

func (wxpat *WXPAccessToken) SetAccURLFmt(urlFmt string) {
	wxpat.AccUrlFmt = urlFmt
}

func (wxpat *WXPAccessToken) UpdateAccessToken() error {
	var err error

	resp, err := http.Get(wxpat.getAccessTokenUrl())
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	accTM := accessTokenMsg{}

	fmt.Println(string(body))
	err = json.Unmarshal([]byte(body), &accTM)

	if err != nil {
		return err
	}

	fmt.Println("DEBUG:", accTM)

	wxpat.accessToken = accTM.Token
	wxpat.ExpireTime = time.Duration(accTM.Expires) * time.Second

	fmt.Println("GET ACC TOKEN:", wxpat.getAccessToken())
	return nil
}
