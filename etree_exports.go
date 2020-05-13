package main

import (
	"github.com/beevik/etree"
	"reflect"
)

func init() {
	GotxSymbols["github.com/beevik/etree"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"CompilePath":         reflect.ValueOf(etree.CompilePath),
		"ErrXML":              reflect.ValueOf(&etree.ErrXML).Elem(),
		"MustCompilePath":     reflect.ValueOf(etree.MustCompilePath),
		"NewCData":            reflect.ValueOf(etree.NewCData),
		"NewCharData":         reflect.ValueOf(etree.NewCharData),
		"NewComment":          reflect.ValueOf(etree.NewComment),
		"NewDirective":        reflect.ValueOf(etree.NewDirective),
		"NewDocument":         reflect.ValueOf(etree.NewDocument),
		"NewDocumentWithRoot": reflect.ValueOf(etree.NewDocumentWithRoot),
		"NewElement":          reflect.ValueOf(etree.NewElement),
		"NewProcInst":         reflect.ValueOf(etree.NewProcInst),
		"NewText":             reflect.ValueOf(etree.NewText),
		"NoIndent":            reflect.ValueOf(etree.NoIndent),

		// type definitions
		"Attr":          reflect.ValueOf((*etree.Attr)(nil)),
		"CharData":      reflect.ValueOf((*etree.CharData)(nil)),
		"Comment":       reflect.ValueOf((*etree.Comment)(nil)),
		"Directive":     reflect.ValueOf((*etree.Directive)(nil)),
		"Document":      reflect.ValueOf((*etree.Document)(nil)),
		"Element":       reflect.ValueOf((*etree.Element)(nil)),
		"ErrPath":       reflect.ValueOf((*etree.ErrPath)(nil)),
		"Path":          reflect.ValueOf((*etree.Path)(nil)),
		"ProcInst":      reflect.ValueOf((*etree.ProcInst)(nil)),
		"ReadSettings":  reflect.ValueOf((*etree.ReadSettings)(nil)),
		"Token":         reflect.ValueOf((*etree.Token)(nil)),
		"WriteSettings": reflect.ValueOf((*etree.WriteSettings)(nil)),

		// interface wrapper definitions
		"_Token": reflect.ValueOf((*_github_com_beevik_etree_Token)(nil)),
	}

}

// _github_com_beevik_etree_Token is an interface wrapper for Token type
type _github_com_beevik_etree_Token struct {
	WIndex  func() int
	WParent func() *etree.Element
}

func (W _github_com_beevik_etree_Token) Index() int             { return W.WIndex() }
func (W _github_com_beevik_etree_Token) Parent() *etree.Element { return W.WParent() }
