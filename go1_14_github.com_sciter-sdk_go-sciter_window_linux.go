package main

import (
	"github.com/sciter-sdk/go-sciter/window"
	"reflect"
)

func init() {
	GotxSymbols["github.com/sciter-sdk/go-sciter/window"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"New": reflect.ValueOf(window.New),

		// type definitions
		"Window": reflect.ValueOf((*window.Window)(nil)),
	}
}
