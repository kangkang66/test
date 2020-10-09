package main

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func RunGin()  {
	r := gin.Default()
	r.LoadHTMLGlob("./*.html")

	r.GET("/good-users", func(c *gin.Context) {
		users := []map[string]interface{}{}
		for _,uid := range RedisGetGoodUser() {
			info := RedisGetUserInfo(uid)
			ps := strings.Split(info.Basicinfo.Photos,",")
			users = append(users, map[string]interface{}{
				"uid":info.Basicinfo.UID,
				"nickname":info.Basicinfo.Nickname,
				"phone":info.Basicinfo.Phone,
				"photos":ps,
				"home":info.Basicinfo.HomelandProvince+"-"+info.Basicinfo.HomelandCity,
				"birthday":info.Basicinfo.Birthday,
				"height":info.Basicinfo.Height,
				"weight":info.Basicinfo.Weight,
			})
		}
		c.HTML(200, "users.html",gin.H{
			"users": users,
		})
	})

	r.GET("/add-block", func(c *gin.Context) {
		uid := c.Query("uid")
		if uid  == "" {
			c.String(200, "uid 为空")
			return
		}
		RedisAddBlockUser(uid)
		RedisDelGoodUser(uid)
		c.String(200, "添加黑名单成功")
	})

	r.GET("/add-collect", func(c *gin.Context) {
		uid := c.Query("uid")
		if uid  == "" {
			c.String(200, "uid 为空")
			return
		}
		resp := AddCollect(uid)
		c.JSON(200, gin.H{
			"resp": resp,
		})
	})

	r.GET("/vistors", func(c *gin.Context) {
		resp := GetVisitors()
		c.HTML(200, "visitors.html",gin.H{
			"visitors": resp.Data.Visitors,
		})
	})


	r.Run(":8080")
}

