package main

import (
	"fmt"
	"io"
	"net"
	"tcptunnel-client/api"
)

func main() {
	server, err := api.NewServer("tcp", "95.111.251.152:8081")
	if err != nil {
		panic(err)
	}

	serverListener := server.Listen()

	for {
		conn, err := serverListener.Accept()
		fmt.Println("ACCEPTED SERVER CONN")
		if err != nil {
			panic(err)
		}

		handle(conn)
	}
}

func handle(conn net.Conn) {
	localConn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic("Port not binded on localhost")
	}

	go func() {
		_, err = io.Copy(localConn, conn)
		if err != nil {
			panic(err)
		}

		conn.Close()
	}()

	go func() {
		_, err := io.Copy(conn, localConn)
		if err != nil {
			panic(err)
		}

		localConn.Close()
	}()
}
