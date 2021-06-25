/**
 * @Author: JYD
 * @Description:hash类型
 * @File:  list
 * @Version: 1.0.0
 * @Date: 2021/2/4 15:12
 */

package redisgo

import "github.com/garyburd/redigo/redis"

type hashRD struct{}

/**
 * @Description: hash set
 * @receiver l
 * @param key
 * @param filed
 * @param value
 * @param exist
 * @return *Reply
 */
func (h *hashRD) Set(key string, filed, value interface{}, exist ...bool) *Reply {
	conn := pool.Get()
	defer conn.Close()
	if len(exist) > 0 && exist[0] {
		return getReply(conn.Do("hsetex", key, filed, value))
	}
	return getReply(conn.Do("hset", key, filed, value))
}

/**
 * @Description: 获取指定字段值
 * @receiver h
 * @param key
 * @param filed
 * @return *Reply
 */
func (h *hashRD) Get(key string, filed interface{}) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("hget", key, filed))
}

/**
 * @Description: 获取所有字段及值
 * @receiver h
 * @param key
 * @return *Reply
 */
func (h *hashRD) GetAll(key string) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("hgetall", key))
}

/**
 * @Description: 设置多个字段及值 [struct]
 * @receiver h
 * @param key
 * @param obj
 * @return *Reply
 */
func (h *hashRD) HMSetFromStruct(key string, obj interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("hmset", redis.Args{}.Add(key).AddFlat(obj)...))
}

/**
 * @Description: 字段删除
 * @receiver h
 * @param key
 * @param fields
 * @return *Reply
 */
func (h *hashRD) Del(key string, fields interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("hdel", redis.Args{}.Add(key).AddFlat(fields)...))
}
