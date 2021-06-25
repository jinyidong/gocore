/**
* @Description: message过滤
* @Author: jinyidong
* @Date: 2021/5/20
* @Version V1.0
 */
package log

type messageFilter func(logger *logger) bool
