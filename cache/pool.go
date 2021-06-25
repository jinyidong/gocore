/**
 * @Author: JYD
 * @Description:
 * @File:  pool
 * @Version: 1.0.0
 * @Date: 2021/2/4 13:34
 */

package redisgo

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	log "github.com/sirupsen/logrus"
	"time"
)

var pool *redis.Pool

func initPool(host, password string) {
	pool = &redis.Pool{
		MaxIdle:     100, //最大空闲数
		MaxActive:   500,
		IdleTimeout: 30 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			return setDialog(host, password)
		},
	}
}

func setDialog(host, password string) (redis.Conn, error) {
	conn, err := redis.Dial("tcp", host)
	if err != nil {
		log.Error(fmt.Sprintf("init redis failed! %v", host))
	}
	if len(password) != 0 {
		if _, err := conn.Do("AUTH", password); err != nil {
			conn.Close()
			log.Error(err)
		}
	}
	if _, err := conn.Do("SELECT", 0); err != nil {
		conn.Close()
		log.Error(err)
	}
	r, err := redis.String(conn.Do("PING"))
	if err != nil || r != "PONG" {
		panic("连接失败")
	}
	return conn, nil
}
