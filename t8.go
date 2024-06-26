package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go send(c1)
	go send(c2)

	cnt := 1

	for {
		select {
		case data, ok := <-c1:
			fmt.Printf("receive data from c1 : %v,%v\n", data, ok)
		case data, ok := <-c2:
			fmt.Printf("receive data from c2 : %v,%v\n", data, ok)
			//default:
			//	fmt.Println("default case")
		}
		cnt += 1
		fmt.Printf("select count : %v\n", cnt)
	}
}

func send(c chan int) {
	data := 0
	for {
		time.Sleep(time.Second * 5)
		data += 1
		c <- data
	}
}
