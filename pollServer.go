package main 

import (
	"net"
	"net/rpc"
	"log"
    "fmt"
	"server"
)
var connectedClients int
func main(){
	
	fmt.Println("start listening")
	fileTransfer := new (server.FileTransfer)
	rpc.Register(fileTransfer)
	
	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	
	//creating a semaphore
	concurrency := 10
	sem := make(chan bool, concurrency)

	for  {
		sem <- true
		if conn, err := listener.Accept(); err != nil {
			log.Fatal("accept error: " + err.Error())
		} else {
			log.Printf("new connection established\n")
			go func (net.Conn) {
				defer func() { <-sem } ()
				rpc.ServeConn(conn)
			}(conn)
		}
		
	}
}

