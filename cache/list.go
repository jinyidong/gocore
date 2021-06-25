/**
 * @Author: JYD
 * @Description:list类型，左侧为头，右侧为尾；list数据结构是双向链表结构；可以用作消息队列；
 * @File:  list
 * @Version: 1.0.0
 * @Date: 2021/2/4 15:12
 */

package redisgo

type listRD struct{}

/**
 * @Description: 向列表头插入元素
 * @receiver l
 * @param key
 * @param token
 * @param exTime
 * @return *Reply
 */
func (l *listRD) LPush(key string, value interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("lpush", key, value))
}

/**
 * @Description: 向列表尾插入元素
 * @receiver l
 * @param key
 * @param token
 * @param exTime
 * @return *Reply
 */
func (l *listRD) RPush(key string, value interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("rpush", key, value))
}

/**
 * @Description: 返回列表头元素
 * @receiver l
 * @param key
 * @param token
 * @param exTime
 * @return *Reply
 */
func (l *listRD) LPop(key string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("lpop", key))
}

/**
 * @Description: 返回列表尾元素
 * @receiver l
 * @param key
 * @param token
 * @param exTime
 * @return *Reply
 */
func (l *listRD) RPop(key string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("cpop", key))
}
