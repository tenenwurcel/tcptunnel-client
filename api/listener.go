package api

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
)

type Listener struct {
	conn net.Conn
}

func NewListener(conn net.Conn) net.Listener {
	return &Listener{conn}
}

func (l *Listener) Accept() (net.Conn, error) {
	reader := bufio.NewReader(l.conn)
	port := 0

	for {
		s, err := reader.ReadString('\n')

		if err == io.EOF {
			continue
		}

		if err != nil {
			return nil, err
		}

		fmt.Println("============")
		fmt.Println(s)

		if strings.Contains(s, "handshake-port") {
			port, err = l.getHandshakePort(s)
			if err != nil {
				return nil, err
			}

			break
		}
	}

	host := fmt.Sprintf("%s:%d", "144.91.68.23", port)
	fmt.Println(host)
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (l *Listener) Close() error {
	err := l.conn.Close()
	if err != nil {
		return err
	}

	return nil
}

func (l *Listener) Addr() net.Addr {
	return l.conn.RemoteAddr()
}

func (l *Listener) getHandshakePort(header string) (int, error) {
	removeBreakLine := strings.Replace(header, "\n", "", -1)
	removeSpaces := strings.TrimSpace(removeBreakLine)
	spliHeader := strings.Split(removeSpaces, ":")
	port, err := strconv.Atoi(spliHeader[1])
	if err != nil {
		return 0, err
	}

	return port, nil
}
