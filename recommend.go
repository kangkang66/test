package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

//获取合适的推荐用户
func RecommendListGoodUser() (goodUser []*User) {
	goodUser = []*User{}
	path := "https://mini.tuodan.tech/jstd-doger/app/user/getRecommendList"

	for i:=1;i<6;i++ {
		data := fmt.Sprintf(`{"type":"","city":"上海","gender":2,"pageNo":%d,"starttime":%d}`, i, time.Now().Unix())
		list := new(ResponseRecommendList)
		content,err := sendRequest(http.MethodPost, path, data)
		if err != nil {
			continue
		}
		err = json.Unmarshal(content, list)
		if err != nil {
			log.Println(err)
			continue
		}
		if !list.Success {
			log.Println("响应失败",list)
			continue
		}
		for _,user := range list.Data {
			RedisAddUser(user)
			if isGoodUser(user) {
				goodUser = append(goodUser, user)
			}
		}
	}

	return
}

func parseBirthday(birthday string) (btime time.Time,err error) {
	btime,err = time.ParseInLocation("2006-01-02", birthday, time.Local)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func isGoodUser(user *User) (is bool) {
	btime,err := parseBirthday(user.Basicinfo.Birthday)
	if err != nil {
		return
	}

	goodHomelandProvince := "安徽"
	goodBirthday	:= 1993
	goodHeight	:= 160
	goodHighestdegree := 1

	//家乡是安徽，年龄<=1993，身高>=160，学历<=本科
	if user.Basicinfo.HomelandProvince == goodHomelandProvince &&
		btime.Year() >= goodBirthday &&
		user.Basicinfo.Height >= goodHeight &&
		user.Basicinfo.Highestdegree <= goodHighestdegree {
		is = true
	}
	return
}


func sendRequest(method string, path string,  postData string) (content []byte,err error) {

	req,err := http.NewRequest(method, path, strings.NewReader(postData))
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Set("user-agent","Mozilla/5.0 (iPhone; CPU iPhone OS 14_0_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/7.0.17(0x17001124) NetType/WIFI Language/zh_CN")
	req.Header.Set("content-type","application/json")
	req.Header.Set("platform","weapp")
	req.Header.Set("referer","https://servicewechat.com/wxdbf2b50ca37c30c0/172/page-frame.html")
	req.Header.Set("token",RedisGetToken())


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

