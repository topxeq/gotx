package main

import (
	"github.com/topxeq/sqltk"
	"reflect"
)

func init() {
	GotxSymbols["github.com/topxeq/sqltk"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ConnectDB":          reflect.ValueOf(sqltk.ConnectDB),
		"ConnectDBNoPing":    reflect.ValueOf(sqltk.ConnectDBNoPing),
		"ExecV":              reflect.ValueOf(sqltk.ExecV),
		"OneLineRecordToMap": reflect.ValueOf(sqltk.OneLineRecordToMap),
		"QueryDBCount":       reflect.ValueOf(sqltk.QueryDBCount),
		"QueryDBI":           reflect.ValueOf(sqltk.QueryDBI),
		"QueryDBNS":          reflect.ValueOf(sqltk.QueryDBNS),
		"QueryDBNSS":         reflect.ValueOf(sqltk.QueryDBNSS),
		"QueryDBS":           reflect.ValueOf(sqltk.QueryDBS),
		"QueryDBString":      reflect.ValueOf(sqltk.QueryDBString),
	}

	GotxSymbols["sqltk"] = GotxSymbols["github.com/topxeq/sqltk"]

}
