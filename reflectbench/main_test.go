package main

import (
	"log"
	"reflect"
	"testing"
)

func TestPlaceholder(t *testing.T) {
	// do nothing.
}

// 2012 rMPB, Go 1.1
// BenchmarkFuncCall	2000000000	         0.30 ns/op
// 2012 rMPB, Go 1.2
// BenchmarkFuncCall	2000000000	         0.29 ns/op
// 2012 rMPB, Go 1.3.1
// BenchmarkFuncCall	2000000000	         0.30 ns/op
func BenchmarkFuncCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doNothing()
	}
}

// 2012 rMPB, Go 1.1
// BenchmarkReflectFuncCall	10000000	       251 ns/op
// as of go1.1, this is a thousand times slower than the BenchmarkFuncCall
// 2012 rMPB, Go 1.2
// BenchmarkReflectFuncCall	10000000	       202 ns/op
// 2012 rMPB, Go 1.3.1
// BenchmarkReflectFuncCall	10000000	       186 ns/op
func BenchmarkReflectFuncCall(b *testing.B) {
	reflectedValue := reflect.ValueOf(doNothing)
	args := []reflect.Value{}
	for i := 0; i < b.N; i++ {
		reflectedValue.Call(args)
	}
}

// 2012 rMPB, Go 1.1
//BenchmarkMethodCall	2000000000	         0.30 ns/op
// 2012 rMPB, Go 1.2
// BenchmarkMethodCall	2000000000	         0.29 ns/op
// 2012 rMPB, Go 1.3.1
// BenchmarkMethodCall	2000000000	         0.30 ns/op
func BenchmarkMethodCall(b *testing.B) {
	var ss SomeStruct
	for i := 0; i < b.N; i++ {
		ss.DoNothing()
	}
}

// 2012 rMPB, Go 1.1
// BenchmarkMethodFuncCall	1000000000	         2.07 ns/op
// as of go1.1, this is ~ ten times slower than BenchmarkMethodCall
// 2012 rMPB, Go 1.2
// BenchmarkMethodFuncCall	1000000000	         2.02 ns/op
// 2012 rMPB, Go 1.3.1
// BenchmarkMethodFuncCall	1000000000	         2.12 ns/op
func BenchmarkMethodFuncCall(b *testing.B) {
	var ss SomeStruct
	f := ss.DoNothing
	for i := 0; i < b.N; i++ {
		f()
	}
}

// 2012 rMPB, Go 1.1
// BenchmarkReflectMethodCall	10000000	       265 ns/op
// as of go1.1, this is a thousand times slower than BenchmarkMethodCall
// 2012 rMPB, Go 1.2
// BenchmarkReflectMethodCall	10000000	       229 ns/op
// 2012 rMPB, Go 1.3.1
// BenchmarkReflectMethodCall	10000000	       261 ns/op
func BenchmarkReflectMethodCall(b *testing.B) {
	b.StopTimer()
	var ss *SomeStruct
	reflectedStruc := reflect.ValueOf(ss)
	args := []reflect.Value{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		reflectedStruc.Method(0).Call(args)
	}
}

// 2012 rMPB, Go 1.1
// BenchmarkReflectMethodCachedCall	10000000	       253 ns/op
// as of go1.1, this is a thousand times slower than BenchmarkMethodCall
// 2012 rMPB, Go 1.2
// BenchmarkReflectMethodCachedCall	10000000	       210 ns/op
// 2012 rMPB, Go 1.3.1
// BenchmarkReflectMethodCachedCall	10000000	       249 ns/op
func BenchmarkReflectMethodCachedCall(b *testing.B) {
	b.StopTimer()
	var ss *SomeStruct
	reflectedStruc := reflect.ValueOf(ss)
	method := reflectedStruc.Method(0)
	args := []reflect.Value{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		method.Call(args)
	}
}

// 2012 rMPB, Go 1.1
// BenchmarkReflectMethodInterfacedCall	10000000	       191 ns/op
// as of go1.1, this is just under a thousand times slower than BenchmarkMethodCall
// 2012 rMPB, Go 1.2
// BenchmarkReflectMethodInterfacedCall	20000000	       142 ns/op
// 2012 rMPB, Go 1.3.1
// BenchmarkReflectMethodInterfacedCall	10000000	       170 ns/op
func BenchmarkReflectMethodInterfacedCall(b *testing.B) {
	b.StopTimer()
	var ss *SomeStruct
	reflectedStruc := reflect.ValueOf(ss)
	method := reflectedStruc.Method(0)
	funcPtr := method.Interface().(func())
	log.Println(funcPtr)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		funcPtr()
	}
}

// 2012 rMPB, Go 1.1
// BenchmarkReflectMethodTypeInterfacedCall	1000000000	         2.08 ns/op
// as of go1.1, this is ~ ten times slower than BenchmarkMethodCall
// 2012 rMPB, Go 1.2
// BenchmarkReflectMethodTypeInterfacedCall	1000000000	         2.32 ns/op
// 2012 rMPB, Go 1.3.1
// BenchmarkReflectMethodTypeInterfacedCall	1000000000	         2.11 ns/op
func BenchmarkReflectMethodTypeInterfacedCall(b *testing.B) {
	b.StopTimer()
	var ss *SomeStruct
	reflectedStructType := reflect.TypeOf(ss)
	method := reflectedStructType.Method(0)
	funcPtr := method.Func.Interface().(func(*SomeStruct))
	log.Println(funcPtr)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		funcPtr(nil)
	}
}

// 2012 rMPB, Go 1.1
// BenchmarkCreateMethodCall	2000000000	         0.30 ns/op
// 2012 rMPB, Go 1.2
// BenchmarkCreateMethodCall	2000000000	         0.29 ns/op
// 2012 rMPB, Go 1.3.1
// BenchmarkCreateMethodCall	2000000000	         0.30 ns/op
func BenchmarkCreateMethodCall(b *testing.B) {
	ssPtr := new(SomeStruct)
	for i := 0; i < b.N; i++ {
		ssPtr.DoNothing()
	}
}

// 2012 rMPB, Go 1.1
// BenchmarkReflectCreateMethodCall	50000000	        47.3 ns/op
// as of go1.1, this is a hundred times slower than BenchmarkCreateMethodCall
// 2012 rMPB, Go 1.2
// BenchmarkReflectCreateMethodCall	50000000	        48.8 ns/op
// 2012 rMPB, Go 1.3.1
// BenchmarkReflectCreateMethodCall	50000000	        58.7 ns/op
func BenchmarkReflectCreateMethodCall(b *testing.B) {
	var ss SomeStruct
	ssType := reflect.TypeOf(ss)
	for i := 0; i < b.N; i++ {
		ssPtr := reflect.New(ssType).Interface().(*SomeStruct)
		ssPtr.DoNothing()
	}
}

// 2012 rMPB, Go 1.2
// BenchmarkCreateStructPointer	2000000000	         0.57 ns/op
// 2012 rMPB, Go 1.3.1
// BenchmarkCreateStructPointer	2000000000	         0.61 ns/op
func BenchmarkCreateStructPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ssPtr := new(SomeStruct)
		ssPtr.DoNothing()
	}
}

// 2012 rMPB, Go 1.2
// BenchmarkCreateStructPointerFunc	500000000	         6.55 ns/op
// 2012 rMPB, Go 1.3.1
// BenchmarkCreateStructPointerFunc	500000000	         6.87 ns/op

func BenchmarkCreateStructPointerFunc(b *testing.B) {
	f := func() *SomeStruct {
		return new(SomeStruct)
	}
	for i := 0; i < b.N; i++ {
		ssPtr := f()
		ssPtr.DoNothing()
	}
}

// 2012 rMPB, Go 1.2
// BenchmarkCreateStructPointerWithRef	50000000	        50.7 ns/op
// 2012 rMPB, Go 1.3.1
// BenchmarkCreateStructPointerWithRef	50000000	        57.6 ns/op
func BenchmarkCreateStructPointerWithRef(b *testing.B) {
	var ss SomeStruct
	reflectedStructType := reflect.TypeOf(ss)
	for i := 0; i < b.N; i++ {
		val := reflect.New(reflectedStructType)
		ssPtr := val.Interface().(*SomeStruct)
		ssPtr.DoNothing()
	}
}
