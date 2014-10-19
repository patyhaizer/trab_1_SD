package client

import (
	"fmt"
	//"bufio"
	"net"
	//"net/rpc"
	"log"
	//"time"
    //"bytes"
    "io"
    "os"
    "strings"
)
const BUFFER_SIZE = 1024
 
func Run(){
	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	//client := rpc.NewClient(conn)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	
	//read from
	var fileName string = "test.docx"

    getFileFromServer(fileName, conn)
	
	conn.Close()
}

func getFileFromServer(fileName string, connection net.Conn) {
	fmt.Println("getfile")
    //var currentByte int64 = 0

    
    var err error
    file, err := os.Create(strings.TrimSpace("src/client/"+fileName))
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(fileName)
    connection.Write([]byte("get " + fileName))
    var n int64
    n, err = io.Copy(file, connection)
    fmt.Println(n, "bytes received")
if err != nil {
    log.Fatal(err)
}
fmt.Println(n, "bytes received")

    
//    for {
//	fileBuffer := make([]byte, BUFFER_SIZE)
//	var n int
//        n,err=connection.Read(fileBuffer)
//        //cleanedFileBuffer := bytes.Trim(fileBuffer, "\x00")
//		fmt.Println(fileBuffer[:n])
//		if err == io.EOF {
//        	fmt.Println("EOF")
//            break
//        }
//        _,err = file.WriteAt(fileBuffer[:n], currentByte)
//	if err == io.EOF {
//        	fmt.Println("EOF")
//            break
//        }
//        currentByte += BUFFER_SIZE
//
//    }
fmt.Println("close")
    file.Close()
    return
}


