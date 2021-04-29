package gcd

import "fmt"

type Log interface {
	Println(args ...interface{})
}

type LogDiscarder struct{}

func (l LogDiscarder) Println(args ...interface{}) {}

type DebugLogger struct{}

func (l DebugLogger) Println(args ...interface{}) {
	fmt.Println(args...)
}
