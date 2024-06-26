package main

import (
	"fmt"
	"os"
	"time"
)

type Counter struct {
	count int
}

func fun1() {
	var mapChan = make(chan map[string]Counter, 1)
	syncChan := make(chan struct{}, 2)
	go func() {
		for {
			if elem, ok := <-mapChan; ok {
				counter := elem["count"]
                                fmt.Printf("Receive counter address : %p\n", &counter)
				counter.count++
			} else {
				break
			}
		}
		fmt.Println("Stopped, [receiver]")
		syncChan <- struct{}{}
	}()
	go func() {
		countMap := map[string]Counter{
			"count": Counter{},
		}
		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v. [sender]\n", countMap)
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}

func fun2() {
	var mapChan = make(chan map[string]*Counter, 1)
	syncChanS := make(chan struct{}, 2)
	syncChanB := make(chan struct{}, 2)
	go func() {
		for {
			if elem, ok := <-mapChan; ok {
				counter := elem["count"]
				counter.count++
			} else {
				break
			}
		}
		fmt.Println("Stopped, [receiver]")
		<-syncChanS
		syncChanB <- struct{}{}
	}()
	go func() {
		countMap := map[string]*Counter{
			"count": &Counter{},
		}
		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Second)
			fmt.Printf("The count map: %v. [sender]\n", countMap)
		}
		close(mapChan)
		<-syncChanS
		syncChanB <- struct{}{}
	}()
	time.Sleep(time.Second)
	syncChanS <- struct{}{}
	syncChanS <- struct{}{}
	<-syncChanB
	<-syncChanB
}

func (this *Counter) String() string {
	return fmt.Sprintf("{count: %d}", this.count)
}

func main() {
	switch os.Args[1] {
	case "value":
		fun1()
	case "ref":
		fun2()
	default:
		fmt.Printf("args : value/ref")
	}
}
