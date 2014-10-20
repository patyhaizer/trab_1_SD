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
type FileTransfer struct{
}

func (t *FileTransfer) GetFile(args *Args, reply *Reply) error {
	//file to read
    file, err := os.Open(strings.TrimSpace("src/"+args.FileName+".docx")) // For read access.
    
    if err != nil && err != io.EOF {
    	log.Fatal(err)
    	defer file.Close() // make sure to close the file even if we panic.
    	return err
	}
	reply.Data = make([]byte, args.BufferSize)
	reply.N, err = file.ReadAt(reply.Data, args.CurrentByte)
	fmt.Println(reply.Data)
	if err == io.EOF {
		reply.EOF = 1
		file.Close()
		return nil
	}
	if err != nil && err != io.EOF {
    	log.Fatal(err)
	}
	reply.EOF = 0
	file.Close()
	//fmt.Println(err, "error")
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

//func connectionHandler(connection net.Conn) {
//    buffer := make([]byte, BUFFER_SIZE)
//
//    _, error := connection.Read(buffer)
//    if error != nil {
//        fmt.Println("There is an error reading from connection", error.Error())
//        return
//    }
//    fmt.Println("command recieved: " + string(buffer))
//
//    //loop until disconntect
//
//    cleanedBuffer := bytes.Trim(buffer, "\x00")
//    cleanedInputCommandString := strings.TrimSpace(string(cleanedBuffer))
//    arrayOfCommands := strings.Split(cleanedInputCommandString, " ")
//
//    fmt.Println(arrayOfCommands[0])
//    sendFileToClient("src/"+arrayOfCommands[1], connection)
//    
//
//}
//
//func sendFileToClient(fileName string, connection net.Conn) {
//   // var currentByte int64 = 0
//    fmt.Println("send to client")
//    
//
//    //file to read
//    file, err := os.Open(strings.TrimSpace(fileName)) // For read access.
//    if err != nil {
//
//        log.Fatal(err)
//    }
//    //var err2 error
//    var n int64
//	n, err = io.Copy(connection, file)
//	if err != nil {
//    	log.Fatal(err)
//	}
//	fmt.Println(n, "bytes sent")
//
//
//    //read file until there is an error
////    for {
////	fileBuffer := make([]byte, BUFFER_SIZE)
////	var n int
////        n, err2 = file.ReadAt(fileBuffer, currentByte)
////        currentByte += BUFFER_SIZE
////        fmt.Println(fileBuffer[:n])
////        connection.Write(fileBuffer[:n])
////
////        if err2 == io.EOF {
////        	fmt.Println("EOF")
////            break
////        }
////    }
//
//    file.Close()
//    return
//
//}
