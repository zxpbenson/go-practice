package main

import (
	"fmt"
	"time"
	"w.src.corp.qihoo.net/cloudsafeeng/mulan/util"
)

func main() {
	uuid := util.Uuid()
	now := time.Now().UnixNano()
	ba := util.IntToBytes(now)
	fmt.Printf("%s\n", uuid)
	fmt.Printf("%v\n", now)
	fmt.Printf("%v\n", len(ba))
	fmt.Printf("%v\n", ba)
	apd := append(ba, []byte(uuid)...)

	fmt.Printf("%v\n", apd)
	md5 := util.Md5(apd)

	fmt.Printf("%s\n", md5)
}
