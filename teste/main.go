package main

import (
	"fmt"
	"net"
)

func main() {
	server, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}

	for {
		fmt.Println("waiting for conn")
		conn, err := server.Accept()
		if err != nil {
			panic(err)
		}

		fmt.Println(conn.RemoteAddr())
	}
}
