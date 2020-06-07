package main

import (
	"reflect"

	"github.com/sciter-sdk/go-sciter/window"
)

func init() {
	GotxSymbols["github.com/sciter-sdk/go-sciter/window"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"New": reflect.ValueOf(window.New),

		// type definitions
		"Window": reflect.ValueOf((*window.Window)(nil)),
	}
}
