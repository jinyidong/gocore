/**
* @Description: 日志处理模块
* @Author: jinyidong
* @Date: 2021/5/20
* @Version V1.0
 */
package log

import (
	"fmt"
	"os"
)

type messageLevel int

const (
	_            = iota + 30 // black
	ColorRed                 // red
	ColorGreen               // green
	ColorYellow              // yellow
	ColorBlue                // blue
	ColorMagenta             // magenta
	_                        // cyan
	ColorWhite               // white
)

const levelColorSeqClear = "\033[0m"

// levelColorFlag, messageLevel color flag.
var levelColorFlag = []string{
	DEBUG:    levelColorSeq(ColorBlue, 0),
	INFO:     levelColorSeq(ColorGreen, 0),
	NOTICE:   levelColorSeq(ColorWhite, 0),
	WARNING:  levelColorSeq(ColorYellow, 0),
	ERROR:    levelColorSeq(ColorRed, 1),
	CRITICAL: levelColorSeq(ColorMagenta, 1),
}

// LevelString, messageLevel string.
var LevelString = map[messageLevel]string{
	DEBUG:    "DEBUG",
	INFO:     "INFO",
	NOTICE:   "NOTICE",
	WARNING:  "WARNING",
	ERROR:    "ERROR",
	CRITICAL: "CRITICAL",
}

const (
	DEBUG    = messageLevel(10 * iota) // DEBUG = 10
	INFO     = messageLevel(10 * iota) // INFO = 20
	NOTICE   = messageLevel(10 * iota) // INFO = 30
	WARNING  = messageLevel(10 * iota) // WARNING = 40
	ERROR    = messageLevel(10 * iota) // ERROR = 50
	CRITICAL = messageLevel(10 * iota) // CRITICAL = 60
)

func levelColorSeq(l messageLevel, way int) string {
	return fmt.Sprintf("\033[%d;%dm", way, messageLevel(l))
}

// logger, define logger entity.
type logger struct {
	Level         messageLevel          // continue only message level gte Level
	Filter        messageFilter         // logger message filter, you can define it as your will.
	Record        *messageRecord        // message entity, you must not instance it.
	StreamHandler *streamMessageHandler // streamMessageHandler
	FileHandler   *fileMessageHandler   // fileMessageHandler
}

var defaultLog *logger

func init() {
	defaultLog = &logger{
		Level: DEBUG,
		StreamHandler: &streamMessageHandler{
			Level: DEBUG,
			Formatter: &messageFormatter{
				Format:     `{{.Time}} [{{.Program}}] [{{.LevelString}}] [{{.FuncName}} {{.Line}}] {{.Message}}`,
				TimeFormat: "2006-01-02 15:04:05.999",
			},
			Destination: os.Stdout,
		},
	}
}

// logger.log, sed message to different handler.
func (l *logger) log(level messageLevel, format string, a ...interface{}) {

	if level >= l.Level {

		l.Record = getMessageRecord(level, format, a...)

		if l.Filter == nil || (l.Filter != nil && l.Filter(l)) {

			if l.StreamHandler != nil && level >= l.StreamHandler.Level {
				if l.StreamHandler.Filter == nil || (l.StreamHandler.Filter != nil && l.StreamHandler.Filter(l)) {
					l.StreamHandler.Write([]byte(l.StreamHandler.Formatter.GetMessage(l)))
				}
			}

			if l.FileHandler != nil && level >= l.FileHandler.Level {
				if l.FileHandler.Filter == nil || (l.FileHandler.Filter != nil && l.FileHandler.Filter(l)) {
					l.FileHandler.Write([]byte(l.FileHandler.Formatter.GetMessage(l)))
				}
			}

		}
	}
}

// logger.Debug, record DEBUG message.
func Debug(format string, a ...interface{}) {
	defaultLog.log(DEBUG, format, a...)
}

// logger.Info, record INFO message.
func Info(format string, a ...interface{}) {
	defaultLog.log(INFO, format, a...)
}

// logger.Notice, record INFO message.
func Notice(format string, a ...interface{}) {
	defaultLog.log(NOTICE, format, a...)
}

// logger.Warning, record WARNING message.
func Warning(format string, a ...interface{}) {
	defaultLog.log(WARNING, format, a...)
}

// logger.Error, record ERROR message.
func Error(format string, a ...interface{}) {
	defaultLog.log(ERROR, format, a...)
}

// logger.Critical, record CRITICAL message.
func Critical(format string, a ...interface{}) {
	defaultLog.log(CRITICAL, format, a...)
}
