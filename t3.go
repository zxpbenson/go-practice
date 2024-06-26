package main

import (
	"fmt"
	//"github.com/gofiber/fiber"
	"sync"
	"time"
)

func main() {
	defer myrecover()
	//var ctx = fiber.Ctx
	//var ctxp = &ctx
	//fmt.Printf("%s\n", ctxp.Method())
	var c chan int
	var wg sync.WaitGroup
	wg.Add(1)
	fmt.Printf("start go :\n")
	go func() {
		fmt.Printf("start receive from chan c :\n")
		d, ok := <-c // 从未初始化的channel中取值会导致当前routine永远阻塞下去

		fmt.Printf("receive from chan ok ? %v, data = %v\n", ok, d)
		wg.Done()
	}()
	go func() {
		defer myrecover()
		var cnt int32 = 0
		for {
			cnt += 1
			fmt.Printf("sleep for time : %v\n", cnt)
			time.Sleep(time.Second)
			if cnt > 10 {
				fmt.Printf("close nil channel and then panic ...\n")
				close(c) // 关闭未初始化的channel会导致 panic: close of nil channel
			}
		}
	}()
	wg.Wait() //如果所有的用户routine都阻塞了，会导致 fatal error: all goroutines are asleep - deadlock!
	fmt.Printf("go done .\n")
}

func myrecover() {
	if err := recover(); err != nil {
		fmt.Printf("Recover err : %v\n", err)
	} else {
		fmt.Printf("no err .")
	}
}
