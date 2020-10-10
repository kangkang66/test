package main

import (
	"log"
	"time"
)

func main()  {
	go getGoodUsers()
	RunGin()
}


func getGoodUsers()  {
	for {
		goodUsers := []*User{}
		ru := RecommendListGoodUser()
		goodUsers = append(goodUsers, ru...)
		du := DynamicListGoodUser()
		goodUsers = append(goodUsers, du...)

		//放进本地goodUsers
		for _,u := range goodUsers {
			if RedisCheckBlockUser(u.Basicinfo.UID) {
				continue
			}
			RedisAddGoodUser(u.Basicinfo.UID)
		}
		log.Println("合格用户:", len(goodUsers))

		//同步外网ip
		SyncIp()
		//请求签到
		Sign()

		time.Sleep(30 * time.Minute)
	}
}