package main

import (
	"fmt"
	"log"
	"runtime"
)

const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)

type VipLevel int

const (
	vipOne VipLevel = 1 + iota
	vipTwo
	vipThree
	vipFour
	vipFive
)

func main() {
	fmt.Printf("runtime version : %s\n", runtime.Version())
	fmt.Printf("Ldata : %v\n", Ldate)
	fmt.Printf("Ltime : %v\n", Ltime)
	fmt.Printf("Lmicroseconds : %v\n", Lmicroseconds)
	fmt.Printf("Llongfile : %v\n", Llongfile)
	fmt.Printf("Lshortfile : %v\n", Lshortfile)
	fmt.Printf("LUTC : %v\n", LUTC)
	fmt.Printf("Lmsgprefix : %v\n", Lmsgprefix)
	fmt.Printf("LstdFlags : %v\n", LstdFlags)

	fmt.Printf("log.Ldata : %v\n", log.Ldate)
	fmt.Printf("log.Ltime : %v\n", log.Ltime)
	fmt.Printf("log.Lmicroseconds : %v\n", log.Lmicroseconds)
	fmt.Printf("log.Llongfile : %v\n", log.Llongfile)
	fmt.Printf("log.Lshortfile : %v\n", log.Lshortfile)
	fmt.Printf("log.LUTC : %v\n", log.LUTC)
	fmt.Printf("log.Lmsgprefix : %v\n", log.Lmsgprefix)
	fmt.Printf("log.LstdFlags : %v\n", log.LstdFlags)

	fmt.Printf("1 << 0 : %v\n", 1<<0)

	fmt.Println(vipOne, vipTwo, vipThree, vipFour, vipFive)
}
