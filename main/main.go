package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func foo() {
	defer wg.Done()
	fmt.Println("foo")
	time.Sleep(time.Second)
	fmt.Println("foo end")

}

func bar() {
	defer wg.Done()
	fmt.Println("bar")
	time.Sleep(time.Second * 2)
	fmt.Println("bar end")
}
func main() {

	start := time.Now()
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
	fmt.Println("程序结束，运行时间为", time.Now().Sub(start))

}
