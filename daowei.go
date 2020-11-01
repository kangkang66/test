package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type DaoWei struct {

}

var DaoWeiClient *DaoWei

func init()  {
	DaoWeiClient = new(DaoWei)
}

func (DaoWei) GetToken() {

}

func (DaoWei) SetToken(token string) {

}

func (d *DaoWei) GetFavUser() (users []map[string]interface{}) {
	data := `{
  "latitude" : "31.000162",
  "sort" : 0,
  "longitude" : "121.425165",
  "isFromProduct" : "0"
}`
	content,err := d.sendRequest(http.MethodPost,"https://i.wangyuedaojia.com/api/v3/document/fav/myFavTUserList", data)
	if err != nil {
		return
	}
	resp := new(DaoWeiFavResponse)
	err = json.Unmarshal(content,resp)
	if err != nil {
		log.Println(err)
		return
	}
	users = []map[string]interface{}{}
	for _,v := range resp.Data {
		info,err := d.GetUserInfo(v.ID)
		if err != nil {
			continue
		}
		ps := []string{
			//info.Data.User.Visible.Avator,
			info.Data.User.Visible.IdcardPhotoURL,
		}
		ps = append(ps, strings.Split(info.Data.User.Services.AvatorPending,",")...)

		users = append(users, map[string]interface{}{
			"phone":info.Data.User.Phone,
			"nickname":info.Data.User.Privacy.Name,
			"birthday":info.Data.User.Visible.Birthday,
			"money":info.Data.User.Money.TIncomeAmount,
			"photos":ps,
			"description":info.Data.User.Services.Description,
		})
	}
	return
}

func (d *DaoWei) GetUserInfo(uid string) (resp *DaoWeiUserInfoResp,err error) {
	resp = new(DaoWeiUserInfoResp)
	info,err := RedisDaoWeiGetUserInfo(uid)
	if err == nil && len(info) > 0{
		err = json.Unmarshal(info,resp)
		return
	}

	path := "https://i.wangyuedaojia.com/api/v3/document/userinfo?userId="+uid
	content,err := d.sendRequest(http.MethodGet, path, "")
	if err != nil {
		return
	}
	err = json.Unmarshal(content, resp)
	if err != nil {
		log.Println(err)
		return
	}
	RedisDaoWeiSetUserInfo(uid, content)
	return
}

func (d *DaoWei) sendRequest(method string, path string,  postData string) (content []byte,err error) {

	req,err := http.NewRequest(method, path, strings.NewReader(postData))
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Set("User-Agent","wang yue/30.62 (iPhone; iOS 13.7; Scale/3.00)")
	req.Header.Set("Content-Type","application/json")
	req.Header.Set("X-EFRESH-SECRET","96cd2955027b4d6cc89958acd94406fb")
	req.Header.Set("X-EFRESH-SESSION","N2M4MTE0MDI0MjU5NDVlODg4OGQwNjZiNjQ3NGU5N2E6NGI3MDdkYzA4ZTk4NDA2MmIwMWMxMGYzNmJjMmIxZmQ=")
	req.Header.Set("Cookie","JSESSIONID=b49eb0a1-5eaf-4fb4-8a35-de053e29ec5c")

	client := &http.Client{}
	resp,err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	content,err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
