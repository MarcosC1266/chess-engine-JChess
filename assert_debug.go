//go:build debug

package main

import (
	"fmt"
	"runtime"
)

func assert(cond bool, msg string) {
	if !cond {
		_, file, line, _ := runtime.Caller(1)
		panic(fmt.Sprintf("ASSERTION FAILED: %s\n\t%s:%d", msg, file, line))
	}
}

func assertf(cond bool, format string, args ...any) {
	if !cond {
		_, file, line, _ := runtime.Caller(1)
		panic(fmt.Sprintf("ASSERTION FAILED: "+format+"\n\t%s:%d", append(args, file, line)...))
	}
}
