package main

import (
	"log"
	"net/http"
)

func main()  {

	user := []*User{}

	ru := RecommendListGoodUser()
	user = append(user, ru...)

	du := DynamicListGoodUser()
	user = append(user, du...)

	oldCollectListResp,err := GetCollectList()
	if err != nil {
		return
	}
	oldNum := len(oldCollectListResp.Data.List)

	//放进收藏夹
	for _,u := range user {
		AddCollect(u.Basicinfo.UID)
	}

	newCollectListResp,err := GetCollectList()
	if err != nil {
		return
	}
	newNum := len(newCollectListResp.Data.List)

	//有新增合格用户
	if newNum > oldNum {
		log.Println("有合格用户被添加进去")
		http.Get("https://sc.ftqq.com/SCU117176Tded7c113c4412e2d9f14944176b8d91f5f7ee9badf795.send?text=有合格用户被添加进去")
	}else{
		log.Println("没找到新合格用户")
	}
}