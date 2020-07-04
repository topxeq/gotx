package main

import (
	"reflect"

	"github.com/sciter-sdk/go-sciter"
)

func WrapFunc(f *(func(...*sciter.Value) *sciter.Value)) *(func(...*sciter.Value) *sciter.Value) {
	return f
}

func init() {
	GotxSymbols["sciterspec"] = map[string]reflect.Value{
		"WrapFunc": reflect.ValueOf(WrapFunc),
	}
}
