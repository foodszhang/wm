package controllers

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"math/rand"
	"time"
)

var redis_conn, err = redis.Dial("tcp", "localhost:6379", redis.DialDatabase(12))

type Token struct {
	TokenPrefix  string
	RedisPrefix  string
	MaxExpiresIn int
}

func (self *Token) make_token(salt string) string {
	//generate token and return
	t := time.Now().UnixNano()
	rand.Seed(t)
	random := rand.Int()
	origin := fmt.Sprintf("%s%s-%s-%s", self.TokenPrefix, t, random, salt)
	return origin
}

func (self *Token) Create(data interface{}, expires int) string {
	result, _ := json.Marshal(data)
	str := string(result)
	fmt.Printf("@@@@@%s\n", str)
	var token string
	for {
		origin := self.make_token(str)
		origin = self.TokenPrefix + origin

		//set key value ex time nx
		token = fmt.Sprintf("%x", sha1.Sum([]byte(origin)))
		fmt.Println(token)
		reply, err := redis_conn.Do("SET", token, str, "ex", expires, "nx")
		fmt.Println(reply, "hahaha", err)
		if reply == "OK" {
			break
		} else {
			time.Sleep(100 * time.Millisecond)
		}
	}

	return token
}

func (self *Token) Delete(token string) {
	redis_conn.Do("DEL", token)
}

func (self *Token) Retrieve(token string) interface{} {
	result, _ := redis_conn.Do("GET", token)

	data_bytes, b := result.([]byte)
	if b {
		fmt.Printf("!!!!%s\n", data_bytes)
	} else {
		fmt.Println("转化失败")
	}
	var data interface{}
	json.Unmarshal(data_bytes, &data)
	fmt.Printf("!%s\n", data)
	return data
}
