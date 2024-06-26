package main

import "net/textproto"
import "fmt"

func main() {
	key := "abc-def"
	fmt.Printf("%s -> %s\n", key, textproto.CanonicalMIMEHeaderKey(key))
}
