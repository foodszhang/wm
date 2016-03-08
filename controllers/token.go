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
	var token string
	var origin string
	for {
		origin = self.make_token(str)

		//set key value ex time nx
		token = fmt.Sprintf("%x", sha1.Sum([]byte(origin)))
		reply, _ := redis_conn.Do("SET", self.RedisPrefix+token, str, "ex", expires, "nx")
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

// 出去之后再进行类型转化
func (self *Token) Retrieve(token string) []byte {
	result, _ := redis_conn.Do("GET", self.RedisPrefix+token)
	data_bytes, b := result.([]byte)
	if !b {
		fmt.Println("转化失败")
	}
	return data_bytes
}

var AuthToken = Token{"auth", "wm", 60 * 60 * 24 * 30}
