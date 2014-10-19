package main 

import (
	"net"
	//"net/rpc"
	"log"
	"bytes"
    "fmt"
    "io"
    "os"
    "strings"
)
const BUFFER_SIZE = 1024

func main(){
	fmt.Println("start listening")
	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	for {
		if conn, err := listener.Accept(); err != nil {
			log.Fatal("accept error: " + err.Error())
		} else {
			log.Printf("new connection established\n")
			connectionHandler(conn)
		}
	}
}

func connectionHandler(connection net.Conn) {
    buffer := make([]byte, BUFFER_SIZE)

    _, error := connection.Read(buffer)
    if error != nil {
        fmt.Println("There is an error reading from connection", error.Error())
        return
    }
    fmt.Println("command recieved: " + string(buffer))

    //loop until disconntect

    cleanedBuffer := bytes.Trim(buffer, "\x00")
    cleanedInputCommandString := strings.TrimSpace(string(cleanedBuffer))
    arrayOfCommands := strings.Split(cleanedInputCommandString, " ")

    fmt.Println(arrayOfCommands[0])
    sendFileToClient(arrayOfCommands[1], connection)
    

}

func sendFileToClient(fileName string, connection net.Conn) {
    var currentByte int64 = 0
    fmt.Println("send to client")
    fileBuffer := make([]byte, BUFFER_SIZE)

    //file to read
    file, err := os.Open(strings.TrimSpace(fileName)) // For read access.
    if err != nil {

        log.Fatal(err)
    }
    var err2 error

    //read file until there is an error
    for {

        _, err2 = file.ReadAt(fileBuffer, currentByte)
        currentByte += BUFFER_SIZE
        fmt.Println(fileBuffer)
        connection.Write(fileBuffer)

        if err2 == io.EOF {
            break
        }
    }

    file.Close()
    return

}
