// +build windows

package time

import (
	"syscall"
	"time"
)

func Now() time.Time {
	var tv syscall.Timeval
	syscall.Gettimeofday(&tv)
	return time.Unix(0, tv.Nanoseconds()) //for windows
}

func NowUnixNano() int64 {
	var tv syscall.Timeval
	syscall.Gettimeofday(&tv)
	return tv.Nano() //for windows
}

func NowUnix() int64 {
	var tv syscall.Timeval
	syscall.Gettimeofday(&tv)
	unixSecond, _ := tv.Unix()
	return unixSecond
}

func Since(t time.Time) time.Duration {
	return Now().Sub(t)
}
