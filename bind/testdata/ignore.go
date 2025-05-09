// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package ignore tests that exported, but otherwise
// unsupported functions, variables, fields and methods
// are ignored by the generators
package ignore

var Var interface{}

type (
	NamedString string
)

const NamedConst NamedString = "foo"

var V interface{}

func Argument(_ interface{}) {
}

func Result() interface{} {
	return nil
}

type S struct {
	F interface{}
}

type (
	F func()
)

func (_ *S) Argument(_ interface{}) {
}

func (_ *S) Result() interface{} {
	return nil
}

type I interface {
	Argument(_ interface{})
	Result() interface{}
}

var (
	Uint   uint
	Uint32 uint32
	Uint64 uint64
	C64    complex64  = 0
	C128   complex128 = 0
)

const (
	Cuint   uint   = 0
	Cuint32 uint32 = 0
	Cuint64 uint64 = 0
)
