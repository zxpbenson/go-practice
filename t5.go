package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("runtime version : %s\n", runtime.Version())
	//ci := make(map[string]interface{})
	ci := make(map[string]any)
	ci["k1"] = "v1"
	fmt.Printf("k1 : %s\n", ci["k1"])
}
