package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func GetVisitors() (resp *ResponseVistors) {
	path := "https://mini.tuodan.tech/jstd-doger/app/user/visitors"
	data := fmt.Sprintf(`{"before":%d,"size":10}`,time.Now().Unix()*1000)
	content,err := sendRequest(http.MethodPost, path, data)
	if err != nil {
		return
	}
	resp = new(ResponseVistors)
	err = json.Unmarshal(content, resp)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(resp.Data.Visitors)
	return
}
