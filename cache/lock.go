/**
 * @Author: JYD
 * @Description:分布式锁
 * @File:  lock
 * @Version: 1.0.0
 * @Date: 2021/2/4 14:15
 */

package redisgo

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

type lockRD struct{}

/**
 * @Description: 获取锁
 * @receiver l
 * @param key
 * @param token
 * @param expire
 * @return *Reply
 */
func (l *lockRD) lock(key, token string, expire ...int64) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("set", key, token, "ex", expire[0], "nx"))
}

const luaScript = `
if redis.call('get', KEYS[1])==ARGV[1] then
	return redis.call('del', KEYS[1])
else
	return 0
end
`

/**
 * @Description: 释放锁，只能够释放自己的锁   lua保证原子操作
 * @receiver l
 * @param key
 * @return err
 */
func (l *lockRD) Unlock(key, token string) *Reply {
	conn := pool.Get()
	defer conn.Close()
	lua := redis.NewScript(1, luaScript)
	return getReply(lua.Do(conn, key, token))
}

/**
 * @Description: 自动续期
 * @receiver l
 * @param key
 * @param token
 * @param exTime
 * @return *Reply
 */
func (l *lockRD) AddTimeout(key, token string, exTime int64) *Reply {
	conn := pool.Get()
	defer conn.Close()
	reply := getReply(conn.Do("ttl", key))
	ttl, err := reply.Int64()
	if err != nil {
		log.Fatal("redis get failed:", err)
	}
	if ttl > 0 {
		return getReply(conn.Do("set", key, token, "ex", int(ttl+exTime)))
	}
	return reply
}
