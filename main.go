package main

import (
	"client"
	"fmt"
)

func test (name string) {
	fmt.Println(name) 
}

func main() {
	client.Run()
	//go client.Run()
	//go client.Run()
	
	var input string
    fmt.Scanln(&input)
}


