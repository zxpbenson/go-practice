package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Printf("runtime version : %s\n", runtime.Version())
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(seq int) {
			fmt.Printf("before once.Do %v\n", seq)
			once.Do(onceBody)
			fmt.Printf("after once.Do %v\n", seq)
			done <- true
			fmt.Printf("after write chan %v\n", seq)
		}(i)
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("main routine before read chan %v\n", i)
		<-done
		fmt.Printf("main routine after read chan %v\n", i)
	}

}
