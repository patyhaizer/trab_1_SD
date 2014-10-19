package client

import (
	"fmt"
	//"bufio"
	"net"
	//"net/rpc"
	"log"
	//"time"
    "bytes"
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
	var fileName string = "test.rtf"

    getFileFromServer(fileName, conn)
	
	conn.Close()
}

func getFileFromServer(fileName string, connection net.Conn) {
	fmt.Println("getfile")
    var currentByte int64 = 0

    fileBuffer := make([]byte, BUFFER_SIZE)

    var err error
    file, err := os.Create(strings.TrimSpace("src/client/"+fileName))
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(fileName)
    connection.Write([]byte("get " + fileName))
    for {

        connection.Read(fileBuffer)
        cleanedFileBuffer := bytes.Trim(fileBuffer, "\x00")

        _, err = file.WriteAt(cleanedFileBuffer, currentByte)

        currentByte += BUFFER_SIZE

        if err == io.EOF {
            break
        }

    }

    file.Close()
    return
}


