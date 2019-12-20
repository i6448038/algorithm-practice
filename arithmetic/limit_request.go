package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

const (
	KEY_PREFIX = "IP_%s_REQ"
	REDIS_PATH = "127.0.0.1:6379" //使用docker的官方redis镜像
	LIMIT_NUMS = 1000
	LIMIT_TIME = 60
	FORBIDDEN_TIME = 10
)



func limit(ip string) (bool, error){
	c, err := redis.Dial("tcp", REDIS_PATH)
	if err != nil {
		fmt.Println(fmt.Sprintf("connect to redis error, %s", err.Error()))
		return false, err
	}
	defer c.Close()

	key := fmt.Sprintf(KEY_PREFIX, ip)
	now := time.Now().Unix()


	length, err := redis.Int64(c.Do("llen", key))
	if err != nil {
		return false, err
	}

	if length < LIMIT_NUMS {
		_, err := c.Do("rpush", key, now)
		if err != nil {
			return false, err
		}
		return false, nil
	}else {
		begin, err := redis.Int64(c.Do("lindex", key, 0))
		if err != nil {
			return false, err
		}
		if now - begin > LIMIT_TIME {
			_, err := c.Do("del", key)
			if err != nil {
				return false, err
			}
			return false, nil
		}else {
			_, err := c.Do("expire", key, FORBIDDEN_TIME)
			if err != nil {
				return false, err
			}
			return true, nil
		}

	}



}