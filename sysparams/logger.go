package applog

import (
	"reflect"
	"runtime"
)

type Logger struct {
	timestamp          string
	e                  error
	level              string
	file_exec_metadata string
	line               int
	pc                 string
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
