package xlog

import (
	"fmt"
	"testing"
)

func TestFormat(t *testing.T) {
	fmt.Printf("%-10s %-30s  ===>  %d  %dms\n", "POST", "/core/data", 200, 20)
}
