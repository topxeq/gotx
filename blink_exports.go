// +build windows

package main

import (
	"reflect"

	"github.com/topxeq/blink"
)

func init() {
	GotxSymbols["github.com/topxeq/blink"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"InitBlink":            reflect.ValueOf(blink.InitBlink),
		"LoadIconFromBytes":    reflect.ValueOf(blink.LoadIconFromBytes),
		"LoadIconFromFile":     reflect.ValueOf(blink.LoadIconFromFile),
		"NewWebView":           reflect.ValueOf(blink.NewWebView),
		"RegisterFileSystem":   reflect.ValueOf(blink.RegisterFileSystem),
		"SetDebugMode":         reflect.ValueOf(blink.SetDebugMode),
		"TempPath":             reflect.ValueOf(&blink.TempPath).Elem(),
		"UnregisterFileSystem": reflect.ValueOf(blink.UnregisterFileSystem),

		// type definitions
		"WebView": reflect.ValueOf((*blink.WebView)(nil)),
	}
}
