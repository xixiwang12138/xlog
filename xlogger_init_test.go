package xlog

import "testing"

func TestTestInitXLogger(t *testing.T) {
	xl := NewXLogger("hello")
	xl.Error(xl, "你好")
}
