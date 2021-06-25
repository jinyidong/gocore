/**
* @Description: (用一句话描述该文件做什么)
* @Author: jinyidong
* @Date: 2021/5/20
* @Version V1.0
 */
package log

import "io"

type messageHandler interface {
	Write(message []byte)
}

// logging.streamMessageHandler
type streamMessageHandler struct {
	Level       messageLevel
	Filter      messageFilter
	Formatter   *messageFormatter
	Destination io.Writer
}

func (handler streamMessageHandler) Write(message []byte) {
	handler.Destination.Write(message)
}

// logging.fileMessageHandler
type fileMessageHandler struct {
	Level       messageLevel
	Filter      messageFilter
	Formatter   *messageFormatter
	Destination io.Writer
}

func (handler fileMessageHandler) Write(message []byte) {
	handler.Destination.Write(message)
}
