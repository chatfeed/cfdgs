package main

import (
	"fmt"
	"flag"
)

var uFlag *string = flag.String("u","suser0","user name")
var aFlag *string = flag.String("a","127.0.0.1:57482","network address")
func main() {
	flag.Parse()
	user := *uFlag
	addr := *aFlag
	fmt.Println(user,addr)
}