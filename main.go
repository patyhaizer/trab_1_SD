package main

import (
	"client"
	"fmt"
	"strconv"
)

func main() {
	fileName :="test.docx"
	
	//creating n clients
	n :=5
	
	for i := 0; i < n; i++ {
        go client.Run(strconv.Itoa(i),fileName)
    }
	
	var input string
    fmt.Scanln(&input)
}


