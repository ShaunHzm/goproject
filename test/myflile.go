package test

import (
	"fmt"
	"reflect"
)

func mystring() {
	var s string = "张三"
	fmt.Println(s, reflect.TypeOf(s))

	b := []byte(s)

	fmt.Println(b, reflect.TypeOf(b))

}

func FileTest() {
	mystring()
}
