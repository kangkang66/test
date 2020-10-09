package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func SyncIp()  {
	path := "http://ip-api.com/json/"
	resp,err := http.Get(path)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	content,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	ipResp := new(ResponseIp)
	err = json.Unmarshal(content, ipResp)
	if err != nil {
		log.Println(err)
		return
	}
	if ipResp.Status != "success" {
		log.Println(resp)
		return
	}

	oldIp,err := RedisSetIp(ipResp.Query)
	if err != nil {
		return
	}
	if oldIp != ipResp.Query {
		path = "https://sc.ftqq.com/SCU117176Tded7c113c4412e2d9f14944176b8d91f5f7ee9badf795.send?text=新的ip:"+ipResp.Query
		http.Get(path)
	}
}
