package main

import (
	"encoding/json"
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
	c,err := RedisClient.HGet("users",uid).Bytes()
	if err != nil {
		log.Println(err)
		return
	}
	info = new(User)
	json.Unmarshal(c, info)
	return
}

func RedisAddGoodUser(user *User) (err error) {
	member := redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: user.Basicinfo.UID,
	}
	err = RedisClient.ZAddNX("good_users", member).Err()
	if err != nil {
		log.Println(err)
	}
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
	data,err := RedisClient.ZRevRange("good_users", 0, -1).Result()
	if err != nil {
		log.Println(err)
		return
	}
	return
}


func RedisCheckBlockUser(user *User) (exist bool) {
	var err error
	exist,err = RedisClient.SIsMember("block_users", user.Basicinfo.UID).Result()
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


