package test

import "fmt"

func TestDefer() {
	foo := func() {
		fmt.Println("I am function foo1")
	}

	defer foo()
	foo = func() {
		fmt.Println("I am function foo2")
	}
	foo()

	x := 10
	defer func(a int) {
		fmt.Println(a)
	}(x)
	x++
	fmt.Println(x)

	x = 20
	defer func() {
		fmt.Println(x)
	}()
	x++
	fmt.Println(x)
}
