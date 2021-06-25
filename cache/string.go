/**
 * @Author: JYD
 * @Description:kv类型
 * @File:  string
 * @Version: 1.0.0
 * @Date: 2021/2/4 13:49
 */

package redisgo

type stringRD struct{}

/**
 * @Description: string类型 set操作
 * @receiver s
 * @param key
 * @param value
 * @param expire 过期时间
 * @return *Reply
 */
func (s *stringRD) Set(key string, value interface{}, expire ...int64) *Reply {
	conn := pool.Get()
	defer conn.Close()
	if len(expire) == 0 {
		return getReply(conn.Do("set", key, value))
	}
	return getReply(conn.Do("set", key, value, "ex", expire[0]))
}

/**
 * @Description: string类型 get操作
 * @receiver s
 * @param key
 * @return *Reply
 */
func (s *stringRD) Get(key string) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("get", key))
}

/**
 * @Description: 自增
 * @receiver s
 * @param key
 * @return *Reply
 */
func (s *stringRD) Incr(key string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("incr", key))
}

/**
 * @Description: 增加指定值
 * @receiver s
 * @param key
 * @param increment
 * @return *Reply
 */
func (s *stringRD) IncrBy(key string, increment int64) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("incrby", key, increment))
}
