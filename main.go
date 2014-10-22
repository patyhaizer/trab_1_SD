package main

import (
	"client"
	"strconv"
	"sync"
)

func main() {
	//100k file
	//fileName :="ExpressoesRegulares.pdf"
	
	//1M file
	//fileName := "levitacao.docx"
	
	//100M file
	fileName := "Sinais_e_Sistemas.pdf"
	
	//creating n clients
	n := 10
	var wg sync.WaitGroup
	
	for i := 1; i < n+1; i++ {
		wg.Add(1)
        go client.Run(strconv.Itoa(i),fileName)
    }
	wg.Wait()
}


