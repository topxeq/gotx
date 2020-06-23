package main

import (
	"reflect"

	"github.com/stretchr/objx"
)

func init() {
	GotxSymbols["github.com/stretchr/objx"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"FromBase64":                   reflect.ValueOf(objx.FromBase64),
		"FromJSON":                     reflect.ValueOf(objx.FromJSON),
		"FromSignedBase64":             reflect.ValueOf(objx.FromSignedBase64),
		"FromURLQuery":                 reflect.ValueOf(objx.FromURLQuery),
		"HashWithKey":                  reflect.ValueOf(objx.HashWithKey),
		"MSI":                          reflect.ValueOf(objx.MSI),
		"MustFromBase64":               reflect.ValueOf(objx.MustFromBase64),
		"MustFromJSON":                 reflect.ValueOf(objx.MustFromJSON),
		"MustFromSignedBase64":         reflect.ValueOf(objx.MustFromSignedBase64),
		"MustFromURLQuery":             reflect.ValueOf(objx.MustFromURLQuery),
		"New":                          reflect.ValueOf(objx.New),
		"Nil":                          reflect.ValueOf(&objx.Nil).Elem(),
		"PathSeparator":                reflect.ValueOf(objx.PathSeparator),
		"SetURLValuesSliceKeySuffix":   reflect.ValueOf(objx.SetURLValuesSliceKeySuffix),
		"SignatureSeparator":           reflect.ValueOf(objx.SignatureSeparator),
		"URLValuesSliceKeySuffixArray": reflect.ValueOf(objx.URLValuesSliceKeySuffixArray),
		"URLValuesSliceKeySuffixEmpty": reflect.ValueOf(objx.URLValuesSliceKeySuffixEmpty),
		"URLValuesSliceKeySuffixIndex": reflect.ValueOf(objx.URLValuesSliceKeySuffixIndex),

		// type definitions
		"MSIConvertable": reflect.ValueOf((*objx.MSIConvertable)(nil)),
		"Map":            reflect.ValueOf((*objx.Map)(nil)),
		"Value":          reflect.ValueOf((*objx.Value)(nil)),

		// interface wrapper definitions
		"_MSIConvertable": reflect.ValueOf((*_github_com_stretchr_objx_MSIConvertable)(nil)),
	}
}

// _github_com_stretchr_objx_MSIConvertable is an interface wrapper for MSIConvertable type
type _github_com_stretchr_objx_MSIConvertable struct {
	WMSI func() map[string]interface{}
}

func (W _github_com_stretchr_objx_MSIConvertable) MSI() map[string]interface{} { return W.WMSI() }
