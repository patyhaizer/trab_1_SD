package main

import (
	"client"
	//"fmt"
	"strconv"
	"sync"
)

func main() {
	//100k file
	fileName :="ExpressoesRegulares.pdf"
	
	//creating n clients
	n := 5
	var wg sync.WaitGroup
	
	for i := 1; i < n+1; i++ {
		wg.Add(1)
        go client.Run(strconv.Itoa(i),fileName)
    }
	wg.Wait()
//	var input string
//    fmt.Scanln(&input)
}


