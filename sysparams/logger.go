package applog

import (
	"fmt"
	"reflect"
	"runtime"
	"time"

	errapp "github.com/milnner/b_modules/errors"
)

const (
	INFO    = "INFO"
	WARNING = "WARNING"
	ERROR   = "ERROR"
	DEBUG   = "DEBUG"
	FATAL   = "FATAL"
)

var (
	levels = []string{INFO, WARNING, ERROR, DEBUG, FATAL}
)

type Logger struct {
	timestamp          string
	e                  error
	level              string
	file_exec_metadata string
	line               int
	pc                 string
}

func NewLogger(e error, level string, file_exec_metadata string, line int, pc string) (*Logger, error) {

	for _, element := range levels {
		if level == element {
			time_log := time.Now().String()
			return &Logger{e: e, timestamp: time_log[:19], level: level, line: line, pc: pc, file_exec_metadata: file_exec_metadata}, nil
		}
	}
	return nil, errapp.NewUndefinedLevelLogError()

}

func (u *Logger) Log() error {
	logMessage := fmt.Sprintf("%s\t[%s]\t%s\t%s\t{in}\t%s\t{line} %d\n", u.timestamp, u.level, u.e.Error(), u.file_exec_metadata, u.pc, u.line)
	fmt.Print(logMessage)
	return nil
}

func (u *Logger) LogDebug() error {
	fmt.Printf("%s\t[%s]\t%s\t%s\t{in}\t%s\t{line} %d\n", u.timestamp, u.level, u.e.Error(), u.file_exec_metadata, u.pc, u.line)
	return nil
}

func FuncName(f interface{}) string {
	rf := reflect.ValueOf(f).Pointer()
	function_name := runtime.FuncForPC(rf).Name()
	return function_name
}

func GetExecutionMetadata() (string, int, string) {
	pc, file, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)

	return file, line, fn.Name()
}
