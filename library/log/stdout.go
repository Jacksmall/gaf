package log

import (
	"context"
	"io"
	"os"
	"time"
)

const defaultPattern = "%L %d-%T %f %M"

type StdoutHandler struct {
	out    io.Writer
	render Render
}

func NewStdout() *StdoutHandler {
	return &StdoutHandler{
		out:    os.Stderr,
		render: newPatternRender(defaultPattern),
	}
}

// Log 处理器 - 记录日志（context + level + D）
func (s *StdoutHandler) Log(ctx context.Context, l Level, args ...D) {
	d := make(map[string]interface{}, 10+len(args))
	for _, v := range args {
		d[v.Key] = v.Value
	}
	// 记录日志其他内容
	// addExtraArgs(ctx, d)
	d[_time] = time.Now().Format(_timeFormat)
	s.render.Render(s.out, d)
	s.out.Write([]byte("\n"))
}

// SetFormat set render format on log output
func (s *StdoutHandler) SetFormat(format string) {
	s.render = newPatternRender(format)
}

// Close 关闭
func (s *StdoutHandler) Close() error {
	return nil
}
