package log

import (
	"context"
	"fmt"
	"io"
	"os"
)

var (
	h Handler
	c *Config
)

var _defaultStdout = NewStdout()

var (
	_v      int
	_stdout bool
	_dir    string
)

type Config struct {
	// Family string
	Host string
	// is stdout
	Stdout bool
	// file dir
	Dir string
	// file buffer size
	FileBufferSize int64
	// max log file
	MaxLogFile int
	Filter     []string
}

// D represents a map of entry level data used for structured logging.
// D map[string]interface{}
type D struct {
	Key   string
	Value interface{}
}

func KV(key string, value interface{}) D {
	return D{
		Key:   key,
		Value: value,
	}
}

type Render interface {
	Render(io.Writer, map[string]interface{}) error
	RenderString(map[string]interface{}) string
}

func init() {
	host, _ := os.Hostname()
	c = &Config{
		Host: host,
	}

	h = newHandlers([]string{}, NewStdout())
}

func Init(conf *Config) {
	if conf == nil {
		conf = &Config{
			Stdout: _stdout,
			Dir:    _dir,
		}
	}

	if len(conf.Host) == 0 {
		host, _ := os.Hostname()
		conf.Host = host
	}
	var hs []Handler
	hs = append(hs, NewStdout())
	// when env is dev
	/*	if conf.Stdout || (isNil && (env.DeployEnv == "" || env.DeployEnv == env.DeployEnvDev)) || _noagent {
			hs = append(hs, NewStdout())
		}
		if conf.Dir != "" {
			hs = append(hs, NewFile(conf.Dir, conf.FileBufferSize, conf.RotateSize, conf.MaxLogFile))
		}
		// when env is not dev
		if !_noagent && (conf.Agent != nil || (isNil && env.DeployEnv != "" && env.DeployEnv != env.DeployEnvDev)) {
			hs = append(hs, NewAgent(conf.Agent))
		}*/
	h = newHandlers(conf.Filter, hs...)
	c = conf
}

// Info logs a message at the info log level.
func Info(format string, args ...interface{}) {
	h.Log(context.Background(), _infoLevel, KV(_log, fmt.Sprintf(format, args...)))
}

// Warn logs a message at the warning log level.
func Warn(format string, args ...interface{}) {
	h.Log(context.Background(), _warnLevel, KV(_log, fmt.Sprintf(format, args...)))
}

// Error logs a message at the error log level.
func Error(format string, args ...interface{}) {
	h.Log(context.Background(), _errorLevel, KV(_log, fmt.Sprintf(format, args...)))
}

// Infov logs a message at the info log level.
func Infov(ctx context.Context, args ...D) {
	h.Log(ctx, _infoLevel, args...)
}

// Warnv logs a message at the warning log level.
func Warnv(ctx context.Context, args ...D) {
	h.Log(ctx, _warnLevel, args...)
}

// Errorv logs a message at the error log level.
func Errorv(ctx context.Context, args ...D) {
	h.Log(ctx, _errorLevel, args...)
}

func logw(args []interface{}) []D {
	if len(args)%2 != 0 {
		Warn("log: the variadic must be plural, the last one will ignored")
	}
	ds := make([]D, 0, len(args)/2)
	for i := 0; i < len(args)-1; i = i + 2 {
		if key, ok := args[i].(string); ok {
			ds = append(ds, KV(key, args[i+1]))
		} else {
			Warn("log: key must be string, get %T, ignored", args[i])
		}
	}
	return ds
}

// Infow logs a message with some additional context. The variadic key-value pairs are treated as they are in With.
func Infow(ctx context.Context, args ...interface{}) {
	h.Log(ctx, _infoLevel, logw(args)...)
}

// Warnw logs a message with some additional context. The variadic key-value pairs are treated as they are in With.
func Warnw(ctx context.Context, args ...interface{}) {
	h.Log(ctx, _warnLevel, logw(args)...)
}

// Errorw logs a message with some additional context. The variadic key-value pairs are treated as they are in With.
func Errorw(ctx context.Context, args ...interface{}) {
	h.Log(ctx, _errorLevel, logw(args)...)
}

func SetFormat(format string) {
	h.SetFormat(format)
}

// Close close resource.
func Close() (err error) {
	err = h.Close()
	h = _defaultStdout
	return
}

func errIncr(lv Level, fn string) {
}
