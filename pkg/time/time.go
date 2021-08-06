// +build !windows

package time

import (
	"syscall"
	"time"
)

func Now() time.Time {
	var tv syscall.Timeval
	syscall.Gettimeofday(&tv)
	return time.Unix(0, syscall.TimevalToNsec(tv))
}

func NowUnixNano() int64 {
	var tv syscall.Timeval
	syscall.Gettimeofday(&tv)
	return syscall.TimevalToNsec(tv)
}

func NowUnix() int64 {
	var tv syscall.Timeval
	syscall.Gettimeofday(&tv)
	sec, _ := tv.Unix()
	return sec
}

func Since(t time.Time) time.Duration {
	return Now().Sub(t)
}
