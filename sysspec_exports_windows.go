// +build windows

package main

import (
	"reflect"

	"syscall"
)

var (
	user32           *syscall.LazyDLL
	getSystemMetrics *syscall.LazyProc
)

const (
	SM_CXSCREEN = 0
	SM_CYSCREEN = 1
)

func GetSystemMetrics(nIndex int) int {
	if user32 == nil {
		user32 = syscall.NewLazyDLL("User32.dll")
	}

	if getSystemMetrics == nil {
		getSystemMetrics = user32.NewProc("GetSystemMetrics")

	}

	index := uintptr(nIndex)
	ret, _, _ := getSystemMetrics.Call(index)
	return int(ret)
}

func GetScreenResolution() (int, int) {
	return GetSystemMetrics(SM_CXSCREEN), GetSystemMetrics(SM_CYSCREEN)
}

func init() {
	GotxSymbols["sysspec"] = map[string]reflect.Value{
		"GetScreenResolution": reflect.ValueOf(GetScreenResolution),
		"GetSystemMetrics":    reflect.ValueOf(GetSystemMetrics),
	}
}
