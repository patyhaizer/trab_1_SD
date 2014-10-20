package main 

import (
	"net"
	"net/rpc"
	"log"
	//"bytes"
    "fmt"
    "io"
    "os"
    "strings"
   // "encoding/gob"
)
const BUFFER_SIZE = 1024
var conn net.Conn
type Reply struct {
	Data []byte
	N ,EOF int
}

type Args struct {
	BufferSize int
	FileName string
	CurrentByte int64
}

type FileTransfer struct{}

//function to be called by client
func (t *FileTransfer) GetFile(args *Args, reply *Reply) error {
	//file to read
    file, err := os.Open(strings.TrimSpace("src/"+args.FileName)) // For read access.
    
    if err != nil && err != io.EOF {
    	log.Fatal(err)
    	defer file.Close() // make sure to close the file even if we panic.
    	return err
	}
	
	reply.Data = make([]byte, args.BufferSize)
	//filling buffer with file data
	reply.N, err = file.ReadAt(reply.Data, args.CurrentByte)

	if err == io.EOF {
		reply.EOF = 1
		file.Close()
		return nil
	}
	if err != nil && err != io.EOF {
		file.Close()
    	log.Fatal(err)
	}
	reply.EOF = 0
	file.Close()

	return err
}
func main(){

	fmt.Println("start listening")
	fileTransfer := new (FileTransfer)
	rpc.Register(fileTransfer)
	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	for {
		var err error
		if conn, err = listener.Accept(); err != nil {
			log.Fatal("accept error: " + err.Error())
		} else {
			log.Printf("new connection established\n")
			rpc.ServeConn(conn)
		}
	}
}