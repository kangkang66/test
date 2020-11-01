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
			ps := strings.Split(info.Basicinfo.Avatar+","+info.Basicinfo.Photos,",")

			intro := ""
			for _,v := range info.Card {
				intro += "【"+ v.Title+"】:"+v.Contenttext+"。"
			}

			users = append(users, map[string]interface{}{
				"uid":info.Basicinfo.UID,
				"nickname":info.Basicinfo.Nickname,
				"phone":info.Basicinfo.Phone,
				"photos":ps,
				"home":info.Basicinfo.HomelandProvince+"-"+info.Basicinfo.HomelandCity,
				"birthday":info.Basicinfo.Birthday,
				"height":info.Basicinfo.Height,
				"weight":info.Basicinfo.Weight,
				"intro":intro,
				"position":info.Basicinfo.Position,
				"company":info.Basicinfo.Company,
				"school":info.Basicinfo.Highestschool,
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

	r.GET("/token", func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			c.JSON(200, gin.H{
				"error": "token empty",
			})
			return
		}
		err := RedisSetToken(token)
		c.JSON(200, gin.H{
			"error": err,
		})
	})

	r.GET("/dao-wei/fav", func(c *gin.Context) {
		users := DaoWeiClient.GetFavUser()
		c.HTML(200, "dao_wei_fav_user.html",gin.H{
			"users": users,
		})
	})


	r.Run(":8081")
}

