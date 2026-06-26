//go:build !debug

package main

func assert(cond bool, msg string)                  {}
func assertf(cond bool, format string, args ...any) {}
