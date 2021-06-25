/**
* @Description: (用一句话描述该文件做什么)
* @Author: jinyidong
* @Date: 2021/5/20
* @Version V1.0
 */
package log

import (
	"bytes"
	"strings"
	"text/template"
	"time"
)

type messageFormatter struct {
	Format     string
	TimeFormat string
}

// logging.messageFormatter.GetMessage, return formatted message string for output.
func (formatter *messageFormatter) GetMessage(logger *logger) string {
	if formatter.TimeFormat == "" {
		formatter.TimeFormat = time.RFC1123
	}
	logger.Record.Time = time.Now().Format(formatter.TimeFormat)
	stringBuffer := new(bytes.Buffer)
	tpl := template.Must(template.New("messageFormat").Parse(formatter.Format))
	tpl.Execute(stringBuffer, *logger.Record)
	message := stringBuffer.String()
	if strings.Index(message, "\n") != len(message)-1 {
		message += "\n"
	}
	return message
}
