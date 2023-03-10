package main

import (
	"fmt"
	"net"
	"net/rpc"
)

func main() {
	l, err := net.Listen("tcp", ":8081")

	if err != nil {
		panic(err)
	}

	go func() {
		for {
			rpc.Accept(l)
		}
	}()

	fmt.Println("Listening")
}
