package sysparams

import (
	"reflect"
	"runtime"
)

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
