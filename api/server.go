package api

import (
	"fmt"
	"log"
	"net"
	"os"
)

type Server struct {
	protocol string
	host     string
	conn     net.Conn
}

func NewServer(protocol string, port string) (*Server, error) {
	host := os.Getenv("TCPTUNNELHOST")
	if host == "" {
		log.Fatal("no host was provided")
	}

	fmt.Println(host)
	connString := fmt.Sprintf("%s%s", host, port)
	conn, err := net.Dial(protocol, connString)
	if err != nil {
		return &Server{}, err
	}

	return &Server{
		protocol: protocol,
		host:     host,
		conn:     conn,
	}, nil
}

func (s *Server) Listen() net.Listener {
	return NewListener(s.conn)
}
