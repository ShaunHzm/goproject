package test

import (
	"fmt"
	"reflect"
	"runtime"
)

// 闭包
func getCounter(f func()) func() {
	calledNum := 0
	return func() {
		f()
		calledNum += 1
		fn := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		fmt.Printf("函数%s第%d次被调用\n", fn, calledNum)
	}
}

func foo() {
	fmt.Println("foo func 执行")
}

func bar() {
	fmt.Println("bar func 执行")
}

func TestClouser() {
	foo := getCounter(foo)
	foo()
	foo()
	foo()

	bar := getCounter(bar)
	bar()
	bar()
}
