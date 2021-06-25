/**
 * @Author: JYD
 * @Description:redis连接池
 * @File:  conn
 * @Version: 1.0.0
 * @Date: 2021/2/4 13:30
 */

package redisgo

import (
	"fmt"
	"github.com/jinyidong/gocore/log"
	"github.com/jinyidong/gocore/util"
	"github.com/mitchellh/mapstructure"
)

type redisConn struct {
	String stringRD
	Hash   hashRD
	Lock   lockRD
}

var RedisConn = new(redisConn)

type redisCfg struct {
	Name     string
	Host     string
	Password string
}

func NewConnection() error {
	redisI := util.GetConfig("redis")
	elementsMap := redisI.([]interface{})
	var cfgs []redisCfg
	for _, vMap := range elementsMap {
		eachElementsMap := vMap.(map[interface{}]interface{})
		var tmpCfg redisCfg
		err := mapstructure.Decode(eachElementsMap, &tmpCfg)
		if err != nil {
			log.Warning(fmt.Sprintf("map转struct失败,%v", err))
			continue
		}
		cfgs = append(cfgs, tmpCfg)
	}

	if len(cfgs) == 0 {
		return fmt.Errorf("获取到redis配置！")
	}

	initPool(cfgs[0].Host, cfgs[0].Password)
	return nil
}
