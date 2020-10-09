package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func GetUserInfo(uid string) (user *User,err error) {
	path := fmt.Sprintf("https://mini.tuodan.tech/jstd-doger/app/user/v5/info?uid=%s",uid)
	content,err := sendRequest(http.MethodGet, path, "")
	if err != nil {
		log.Println(err)
		return
	}
	resp := new(ResponseUserInfo)
	err = json.Unmarshal(content, resp)
	if err != nil {
		log.Println(err)
		return
	}
	if !resp.Success {
		log.Println("响应失败",resp)
		err = errors.New("获取用户信息响应失败")
		return
	}
	user = resp.Data.UserInfo
	return
}
