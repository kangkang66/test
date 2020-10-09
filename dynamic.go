package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//获取用户动态
func DynamicListGoodUser() (goodUser []*User) {
	goodUser = []*User{}
	path := "https://mini.tuodan.tech/jstd-doger/app/dynamic/v2/getList"

	for i:=1;i<10;i++ {
		data := fmt.Sprintf(`{"pageNo":%d,"pageSize":10,"gender":"","type":"STAR"}`,i)
		content,err := sendRequest(http.MethodPost, path, data)
		if err != nil {
			continue
		}

		list := new(ResponseDynamicList)
		err = json.Unmarshal(content, list)
		if err != nil {
			log.Println(err)
			continue
		}

		if !list.Success {
			log.Println("响应失败",list)
			continue
		}

		for _,v := range list.Data {
			if v.City == "上海" && v.Gender == 2 {
				user,err := GetUserInfo(v.UID)
				if err != nil {
					continue
				}
				if isGoodUser(user) {
					goodUser = append(goodUser, user)
				}
			}
		}
	}

	return
}


