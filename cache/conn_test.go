/**
* @Description: 单元测试
* @Author: jinyidong
* @Date: 2021/6/25
* @Version V1.0
 */
package redisgo

import "testing"

func TestNewConnection(t *testing.T) {
	err := NewConnection()
	if nil != err {
		t.Error(err)
		return
	}
	RedisConn.Hash.Set("sys_salary", "1", "10~20")
}
