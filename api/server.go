package api

import "net"

type Server struct {
	protocol string
	host     string
	conn     net.Conn
}

func NewServer(protocol string, host string) (*Server, error) {
	conn, err := net.Dial(protocol, host)
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
