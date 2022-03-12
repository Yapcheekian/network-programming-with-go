package main

import (
	"net"
	"time"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "www.google.com:http")
	if err != nil {
		panic(err)
	}

	tcpConn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		panic(err)
	}

	err = tcpConn.SetKeepAlive(true)
	err = tcpConn.SetKeepAlivePeriod(time.Minute)
	err = tcpConn.SetLinger(-1) // anything < 0 uses the default behavior

	if err := tcpConn.SetReadBuffer(212992); err != nil {
		panic(err)
	}

	if err := tcpConn.SetWriteBuffer(212992); err != nil {
		panic(err)
	}
}
