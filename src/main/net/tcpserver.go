package net

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
)

func toAddr(add string) *net.TCPAddr {
	dd, _ := net.ResolveTCPAddr("tcp", add)
	return dd
}

func startServer() {
	add := toAddr(addr)
	listener, err := net.ListenTCP("tcp", add)
	if err != nil {
		fmt.Println("listen failed ", addr)
		return
	}
	go listenAccept(listener)
	fmt.Println("listener started")
}

func listenAccept(listener *net.TCPListener) {
	for {
		tcpcon, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("listener accept error close listener")
			listener.Close()
			break
		}
		fmt.Println("tcp acceptor new connection in")
		go handleTcpConn(tcpcon)
	}
}

func tcpWrite(conn *net.TCPConn, data []byte) (n int, err error) {
	c, er := conn.Write(data)
	return c, er
}

/**
 *读取数据，放到buf中，又一次拷贝
 *
 */
func tcpRead(conn *net.TCPConn) (data []byte, err error) {
	result := bytes.NewBuffer(nil)
	buf := make([]byte, 512)
	for {
		fmt.Println("server start read")
		n, er := conn.Read(buf)
		fmt.Println("server read ")
		result.Write(buf[:n])
		//结束？？？？？
		if er != nil {
			if er == io.EOF {
				break
			} else {
				fmt.Println("tcp conn read data error")
				conn.Close()
				errors.New("tcp conn read data error")
				return
			}
		}
		break
	}
	return result.Bytes(), nil
}

func handleTcpConn(conn *net.TCPConn) {
	data, err := tcpRead(conn)
	str := string(data)
	echo := "this is a go tcp server."
	strs := []string{str, echo}
	resp := strings.Join(strs, " ")
	rdata := []byte(resp)
	n, err := tcpWrite(conn, rdata)
	if err != nil {
		fmt.Println("server send failed ")
	} else {
		fmt.Println("server send success ", n)
	}
	conn.Close()
}

func ServerStart() {
	startServer()
}
