package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func AddCollect(uid string) (resp *ResponseCollect) {
	path := fmt.Sprintf("https://mini.tuodan.tech/jstd-doger/app/collect/v2/add/%s",uid)
	content,err := sendRequest(http.MethodPost, path, "")
	if err != nil {
		return
	}
	resp = new(ResponseCollect)
	err = json.Unmarshal(content, resp)
	if err != nil {
		log.Println(err)
		return
	}
	if !resp.Success {
		log.Println("响应失败",resp)
		return
	}
	return
}

func GetCollectList() (resp *ResponseCollectList,err error) {
	path := "https://mini.tuodan.tech/jstd-doger/app/im/v1/collectList"
	data := `{"pageNo":1,"pageSize":200,"relationType":2,"unlockType":1,"firstPullTime":"2020-10-08T10:52:21.581Z"}`
	content,err := sendRequest(http.MethodPost, path, data)
	if err != nil {
		return
	}
	resp = new(ResponseCollectList)
	err = json.Unmarshal(content, resp)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
