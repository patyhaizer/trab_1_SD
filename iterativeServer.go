package main 

import (
	"net"
	"net/rpc"
	"log"
	//"bytes"
    "fmt"
	"server"
   // "encoding/gob"
)

func main(){

	fmt.Println("start listening")
	fileTransfer := new (server.FileTransfer)
	rpc.Register(fileTransfer)
	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	for {
		if conn, err := listener.Accept(); err != nil {
			log.Fatal("accept error: " + err.Error())
		} else {
			log.Printf("new connection established\n")
			rpc.ServeConn(conn)
		}
	}
}