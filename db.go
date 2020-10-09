package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"time"
)

var RedisClient *redis.Client

func init()  {
	RedisClient = newRedisClient()
}

func newRedisClient() (client *redis.Client) {
	client = redis.NewClient(
		&redis.Options{
			Addr:         ":6379",
			Password:     "",
			DB:                 0,
			DialTimeout:        10 * time.Second,
			ReadTimeout:        30 * time.Second,
			WriteTimeout:       30 * time.Second,
			PoolSize:           10,
			PoolTimeout:        30 * time.Second,
			IdleTimeout:        500 * time.Millisecond,
			IdleCheckFrequency: 500 * time.Millisecond,
		})
	err := client.Ping().Err()
	if err != nil {
		panic(err)
	}
	return
}

func RedisAddUser(user *User)  {
	err := RedisClient.HSetNX("users", user.Basicinfo.UID, user.ToJson()).Err()
	if err != nil {
		log.Println(err)
	}
}

func RedisGetUserInfo(uid string) (info *User) {
	info = new(User)
	c,err := RedisClient.HGet("users",uid).Bytes()
	if err != nil {
		log.Println(err)
		return
	}
	info = new(User)
	json.Unmarshal(c, info)
	return
}

func RedisAddGoodUser(uid string) (err error) {
	member := redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: uid,
	}
	err = RedisClient.ZAdd("good_users", member).Err()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(err)
	return
}

func RedisDelGoodUser(uid string)  {
	err := RedisClient.ZRem("good_users", uid).Err()
	if err != nil {
		log.Println(err)
	}
	return
}

func RedisGetGoodUser() (data []string) {
	data = []string{}
	var err error
	data,err = RedisClient.ZRevRange("good_users", 0, -1).Result()
	if err != nil {
		log.Println(err)
		return
	}
	return
}


func RedisCheckBlockUser(uid string) (exist bool) {
	var err error
	exist,err = RedisClient.SIsMember("block_users", uid).Result()
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func RedisAddBlockUser(uid string)  {
	err := RedisClient.SAdd("block_users", uid).Err()
	if err != nil {
		log.Println(err)
	}
}

func RedisSetIp(newIp string) (oldIp string, err error) {
	oldIp,err = RedisClient.GetSet("ip",newIp).Result()
	if err == redis.Nil {
		err = nil
	}
	if err != nil {
		log.Println("RedisSetIp",err)
		return "", err
	}
	return
}


