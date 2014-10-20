package client

import (
	"fmt"
	//"bufio"
	"net"
	"net/rpc"
	"log"
	//"time"
    //"bytes"
    //"io"
    "os"
    "strings"
    //"encoding/gob"
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
func Run(id string){
	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	client := rpc.NewClient(conn)
	//file to write to
	fileName := "test"
	file, err := os.Create(strings.TrimSpace("src/client/"+fileName+id+".docx"))
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
		fmt.Println(reply.Data)
		_,err = file.WriteAt(reply.Data[:reply.N], args.CurrentByte)
		if err != nil {
			log.Println("arith error2:", err)
			break
		}
		args.CurrentByte+=BUFFER_SIZE
		if reply.EOF == 1 {
			break
		}

    }
	fmt.Println("fim")
	file.Close()
	
	conn.Close()
}

//func getFileFromServer(fileName string, connection net.Conn) {
//	fmt.Println("getfile")
//    //var currentByte int64 = 0
//
//    
//    var err error
//    file, err := os.Create(strings.TrimSpace("src/client/"+fileName))
//    if err != nil {
//        log.Fatal(err)
//    }
//    fmt.Println(fileName)
//    connection.Write([]byte("get " + fileName))
//    var n int64
//    n, err = io.Copy(file, connection)
//    fmt.Println(n, "bytes received")
//if err != nil {
//    log.Fatal(err)
//}
//fmt.Println(n, "bytes received")
//
//    
////    for {
////	fileBuffer := make([]byte, BUFFER_SIZE)
////	var n int
////        n,err=connection.Read(fileBuffer)
////        //cleanedFileBuffer := bytes.Trim(fileBuffer, "\x00")
////		fmt.Println(fileBuffer[:n])
////		if err == io.EOF {
////        	fmt.Println("EOF")
////            break
////        }
////        _,err = file.WriteAt(fileBuffer[:n], currentByte)
////	if err == io.EOF {
////        	fmt.Println("EOF")
////            break
////        }
////        currentByte += BUFFER_SIZE
////
////    }
//fmt.Println("close")
//    file.Close()
//    return
//}
//
//
