/**
* @Description: 日志信息
* @Author: jinyidong
* @Date: 2021/5/20
* @Version V1.0
 */
package log

import (
	"fmt"
	"github.com/jinyidong/gocore/util"
	"os"
	"path/filepath"
	"runtime"
)

var program string

func init() {
	serviceNameI := util.GetConfig("service_name")
	program = serviceNameI.(string)
}

type messageRecord struct {
	Level         messageLevel
	LevelString   string
	Message       string
	Pid           int
	Program       string
	Time          string
	FuncName      string
	LongFileName  string
	ShortFileName string
	Line          int
	Color         string
	ColorClear    string
}

// logging.getMessageRecord, make a record and return it's reference.
func getMessageRecord(level messageLevel, format string, a ...interface{}) *messageRecord {
	message := fmt.Sprintf(format, a...)
	pc, file, line, _ := runtime.Caller(3)
	record := &messageRecord{
		Level:         level,
		Message:       message,
		Pid:           os.Getpid(),
		Program:       program,
		Time:          "",
		FuncName:      runtime.FuncForPC(pc).Name(),
		LongFileName:  file,
		ShortFileName: filepath.Base(file),
		Line:          line,
		Color:         levelColorFlag[level],
		ColorClear:    levelColorSeqClear,
		LevelString:   LevelString[level],
	}
	return record
}
