package applog

import (
	"fmt"
	"os"
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
	file               string
	file_exec_metadata string
	line               int
	pc                 string
}

func NewLogger(e error, level string, file_exec_metadata string, line int, pc string, log_folder string) (*Logger, error) {

	if log_folder == "" {
		return nil, errapp.NewUndefinedLogFolderError()

	}

	for _, element := range levels {
		if level == element {
			time_log := time.Now().String()
			file := "log-" + time_log[:10] + "[" + level + "]" + ".log"
			log_file := log_folder + "/" + file

			ok, err := checkFileLogExist(log_file)

			if err != nil {
				return nil, err

			} else if !ok {
				os.Create(log_file)
			}

			return &Logger{e: e, timestamp: time_log[:19], level: level, file: log_file, line: line, pc: pc, file_exec_metadata: file_exec_metadata}, nil
		}
	}
	return nil, errapp.NewUndefinedLevelLogError()

}

func (u *Logger) Log() error {
	f, err := os.OpenFile(u.file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	logMessage := fmt.Sprintf("%s\t[%s]\t%s\t%s\t{in}\t%s\t{line} %d\n", u.timestamp, u.level, u.e.Error(), u.file_exec_metadata, u.pc, u.line)

	_, err = f.WriteString(logMessage)
	if err != nil {
		return err
	}

	return nil
}

func (u *Logger) LogDebug() error {
	f, err := os.OpenFile(u.file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Printf("%s\t[%s]\t%s\t%s\t{in}\t%s\t{line} %d\n", u.timestamp, u.level, u.e.Error(), u.file_exec_metadata, u.pc, u.line)

	return nil
}

func FuncName(f interface{}) string {
	rf := reflect.ValueOf(f).Pointer()
	function_name := runtime.FuncForPC(rf).Name()
	return function_name
}

func checkFileLogExist(f string) (bool, error) {
	_, err := os.Stat(f)

	if err == nil {

		return true, nil
	} else if os.IsNotExist(err) {

		return false, nil
	} else {

		return false, err
	}
}

func GetExecutionMetadata() (string, int, string) {
	pc, file, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)

	return file, line, fn.Name()
}
