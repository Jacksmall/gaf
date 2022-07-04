package log

import (
	"runtime"
	"strconv"
	"sync"
)

var fm sync.Map

func funcName(skip int) (name string) {
	if pc, _, line, ok := runtime.Caller(skip); ok {
		if v, ok := fm.Load(pc); ok {
			name = v.(string)
		} else {
			name = runtime.FuncForPC(pc).Name() + ":" + strconv.FormatInt(int64(line), 10)
			fm.Store(pc, name)
		}
	}
	return
}
