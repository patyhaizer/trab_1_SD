package main

import (
	"client"
	"fmt"
)

func test (name string) {
	fmt.Println(name) 
}

func main() {
	go client.Run("1")
	//fmt.Println("saiu")
	go client.Run("2")
	//go client.Run()
	
	var input string
    fmt.Scanln(&input)
}


