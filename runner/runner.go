/*
 * 运行任务包-适用于cron
 */
package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	// 系统中断信号通道
	interrupt chan os.Signal
	// 是否完成, 错误进入的通道
	complete chan error
	// 任务是否超时, 是否超时通道
	timeout <-chan time.Time
	// tasks 要运行的任务列表
	// 包含多个函数的切片，函数参数为任务id int
	tasks []func(int)
}

// 中断错误
var ErrInterrupt = errors.New("os signal interrupt")

// 超时错误
var ErrTimeout = errors.New("timeout")

// New Runner
func New(d time.Duration) *Runner {
	return &Runner{
		// buffer 量为1的系统信号通道, 至少能接收一个来自语言运行时的os.Signal
		// 确保语言运行时发生这个事件时不会被堵塞
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// 添加要执行的tasks ...表示[]func(int)
func (r *Runner) Add(fn ...func(int)) {
	r.tasks = append(r.tasks, fn...)
}

// 开始
func (r *Runner) Start() error {
	// 我们希望接收到所有中断信号
	signal.Notify(r.interrupt, os.Interrupt)
	// 用不同的goroutine 运行任务
	go func() {
		r.complete <- r.run()
	}()
	// 判断是否完成或者超时
	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

// 运行每一个已注册的任务
func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		// 执行
		task(id)
	}
	return nil
}

// 验证是否接收到了中断信号
func (r *Runner) gotInterrupt() bool {
	select {
	// 接收到了中断信号
	case <-r.interrupt:
		// 中断
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
