package main

import (
	"flag"
	"net"
	"os"
	//	sync "github.com/larspensjo/Go-sync-evaluation/evalsync"
	"fmt"
	"github.com/asjustas/goini"
	"github.com/golang/glog"
)

// type uconn struct{
// 	conn	net.Conn
// 	sync.RWMutex
// }
var (
	conf_path = flag.String("c", "../s1.conf", "config path")
)

func main() {
	flag.Parse()
	c, err := goini.Load(*conf_path)
	if err != nil {
		panic(err)
	}
	addr := c.Str("server", "addr")

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		glog.Fatal("Failed listening: ", err, "\n")
		return
	}
	fmt.Println("Listen on :", addr)
	for failures := 0; failures < 100; {
		conn, err := listener.Accept()
		if err != nil {
			glog.Fatal("Failed listening: ", err, "\n")
			failures++
		}

		if ok, index := NewClientConnect(conn); ok {
			go HandleConnection(conn, index)
		}
	}
	glog.Fatal("Too many listener.Accept() errors, giving up")
	os.Exit(1)
}
