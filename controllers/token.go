package controllers

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"time"
)

var redis_conn, err = redis.Dial("tcp", ":6379")

type Token struct {
	max_expires_in time.Duration
	token_prefix   string
	redis_prefix   string
}

func (self *Token) Create(data interface{}, expires time.Duration) string {
	result, _ := json.Marshal(data)
	//generate token and return
	return string(result)
}

func (self *Token) Delete(token string) {
	redis_conn.Do("DEL", token)
}

func (self *Token) Retrieve(token string) interface{} {
	result, _ := redis_conn.Do("GET", token)
	return result

}
