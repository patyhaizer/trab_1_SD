package client

import (
	"fmt"
	//"bufio"
	"net"
	"net/rpc"
	"log"
	"time"
    //"bytes"
    //"io"
    "os"
    "strings"
)
const BUFFER_SIZE = 1024
type Reply struct {
	Data []byte
	N, EOF int
}
type Args struct {
	BufferSize int
	FileName string
	CurrentByte int64
}
func Run(id string,fileName string){
	//time of request of connection
	t0 := time.Now()
	//Creating connection with Server
	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	//Creating rpc Client
	client := rpc.NewClient(conn)
	//calling funtion to get file from server and write on client
	getFileFromServer(id,client,fileName)
	
	//time - end of file download
	t1 := time.Now()
	fmt.Printf("The download of file of client %s took %v to run.\n",id, t1.Sub(t0))
	conn.Close()
}

func getFileFromServer(id string,client *rpc.Client,fileName string) {
	
	//file to write to
	file, err := os.Create(strings.TrimSpace("src/client/"+id+fileName))
	if err != nil {
		log.Fatal(err)
	}
	var reply Reply
	args := &Args{BUFFER_SIZE,fileName,0}
	
	for {
		err = client.Call("FileTransfer.GetFile", args, &reply)
		if err != nil {
			log.Println("arith error:", err)
			break
		}

		_,err = file.WriteAt(reply.Data[:reply.N], args.CurrentByte)
		if err != nil {
			log.Println("arith error2:", err)
			break
		}
		if reply.EOF == 1 {
			break
		}
		args.CurrentByte+=BUFFER_SIZE

    }

	file.Close()
}
