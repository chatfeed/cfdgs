package main

import (
	//"fmt"
	"github.com/golang/glog"
	sync "github.com/larspensjo/Go-sync-evaluation/evalsync"
	"io"
	"net"
)

type client struct {
	conn net.Conn
}

var (
	allClientsSem sync.RWMutex
	allClients    [MAX_CONN]*client
	clientNum     = 0
)

//给新连的客户端分配id
func NewClientConnect(conn net.Conn) (ok bool, index int) {
	allClientsSem.Lock()
	defer allClientsSem.Unlock()
	var i int
	for i = 0; i < MAX_CONN; i++ {
		if allClients[i] == nil {
			break
		}
	}
	if i == MAX_CONN {
		return false, 0
	}
	cli := new(client)
	cli.conn = conn
	clientNum++
	allClients[i] = cli
	return true, i
}

func HandleConnection(conn net.Conn, i int) {
	buff := make([]byte, 1024)
	for {
		n, err := conn.Read(buff)
		if err == io.EOF {
			break
		} else if err != nil {
			glog.Fatal(err)
		}
		glog.Info("ManageOneClient: command (%d bytes) %v\n", n, string(buff[:n]))
		conn.Write(buff)
	}
}
