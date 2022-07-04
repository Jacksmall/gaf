package log

import (
	"context"
	pkgerr "github.com/pkg/errors"
	"time"
)

const (
	_time       = "time"
	_timeFormat = "2006-01-02T15:04:05.999999"
	_source     = "source"
	_level      = "level"
	_levelValue = "level_value"
	_log        = "log"
	_instanceID = "instance_id"
	// container environment: prod, pre, uat, fat.
	_deplyEnv = "env"
	// container area.
	_zone = "zone"
	// app name.
	_appID = "app_id"
)

type Handler interface {
	// Log 处理器 - 记录日志（context + level + D）
	Log(context.Context, Level, ...D)

	// SetFormat set render format on log output
	SetFormat(string)

	// Close 关闭
	Close() error
}

type Handlers struct {
	// filter func
	filters map[string]struct{}
	// handlers a bound of handler
	handlers []Handler
}

// newHandlers
func newHandlers(filter []string, handlers ...Handler) *Handlers {
	set := make(map[string]struct{})
	for _, s := range filter {
		set[s] = struct{}{}
	}
	return &Handlers{filters: set, handlers: handlers}
}

// Log 处理器 - 记录日志（context + level + D）
func (hs Handlers) Log(c context.Context, lv Level, d ...D) {
	var fn string
	for i := range d {
		if _, ok := hs.filters[d[i].Key]; ok {
			d[i].Value = "***"
		}
		if d[i].Key == _source {
			fn = d[i].Value.(string)
		}
	}
	if fn == "" {
		d = append(d, KV(_source, funcName(4)))
	}
	d = append(d, KV(_time, time.Now()), KV(_levelValue, int(lv)), KV(_level, string(lv)))
	errIncr(lv, fn)
	for _, h := range hs.handlers {
		h.Log(c, lv, d...)
	}
}

// SetFormat set render format on log output
func (hs Handlers) SetFormat(format string) {
	for _, h := range hs.handlers {
		h.SetFormat(format)
	}
}

// Close 关闭
func (hs Handlers) Close() (err error) {
	for _, h := range hs.handlers {
		if e := h.Close(); e != nil {
			err = pkgerr.WithStack(e)
		}
	}
	return
}
