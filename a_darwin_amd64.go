//go:build darwin
// +build darwin

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	fmt.Println("main routine start")
	go func() {
		defer wg.Done()
		fmt.Println("sub routine start")
		time.Sleep(time.Second * 2)
		fmt.Println("sub routine terminated")
	}()
	fmt.Println("main routine ternimated")
	wg.Wait()
}
