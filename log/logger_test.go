/**
* @Description: (用一句话描述该文件做什么)
* @Author: jinyidong
* @Date: 2021/5/20
* @Version V1.0
 */
package log

import "testing"

func TestLogger_Error(t *testing.T) {
	Debug("hello world, %s", "logging")
}
