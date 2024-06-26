package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//c := make(chan interface{})
	ticker := time.NewTicker(time.Second)
	var wg sync.WaitGroup
	wg.Add(1)
	go closeChanWhenRead(ticker.C, &wg) //不传指针过去会把结构体复制然诺后传过去
	fmt.Printf("main : send one data\n")
	//c <- 1
	//fmt.Printf("main : close chan\n")
	//close(c)
	fmt.Printf("main : sleep\n")
	time.Sleep(5 * time.Second)
	fmt.Printf("main : close ticker.C\n")
	ticker.Stop()
	fmt.Printf("main : sleep again\n")
	time.Sleep(5 * time.Second)
	fmt.Printf("main : wait\n")
	wg.Wait()
	fmt.Printf("main : done\n")
}

func closeChanWhenRead(cc <-chan time.Time, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		fmt.Printf("sub : wg.Done()\n")
		//time.Sleep(5 * time.Second)
		//fmt.Printf("sub : sleep done\n")
		if err := recover(); err != nil {
			fmt.Printf("sub : recover panic : %v\n", err)
		}
	}()
LABEL:
	for {
		select {
		case _, ok := <-cc:
			if ok {
				fmt.Printf("sub : read data\n")
			} else {
				fmt.Printf("sub : channel closed\n")
				break LABEL
			}
		}
	}
	fmt.Printf("sub : closeChanWhenRead done\n")
}
