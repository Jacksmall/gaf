package log

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func initFile() {
	conf := &Config{
		Dir: "/tmp",
	}
	Init(conf)
}

func testLog(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		Error("hello %s", "world")
		Errorv(context.Background(), KV("key", 2222222), KV("test2", "test"))
	})
	t.Run("Warn", func(t *testing.T) {
		Warn("hello %s", "world")
		Warnv(context.Background(), KV("key", 2222222), KV("test2", "test"))
	})
	t.Run("Info", func(t *testing.T) {
		Info("hello %s", "world!")
		Infov(context.Background(), KV("key", 2222222), KV("test2", "test"))
	})
}

func TestFile(t *testing.T) {
	initFile()
	testLog(t)
	assert.Equal(t, nil, Close())
}
