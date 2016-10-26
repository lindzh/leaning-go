package net

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
)

var addr = "127.0.0.1:6666"

func connect() (con *net.TCPConn, err error) {
	radd := toAddr(addr)
	ladd := toAddr("127.0.0.1:6660")

	conn, err := net.DialTCP("tcp", ladd, radd)
	if err != nil {
		fmt.Println("dial tcp failed")
		errors.New("dial tcp failed")
		return
	}
	return conn, nil
}

func writeData(conn *net.TCPConn, data []byte) (n int, err error) {
	return conn.Write(data)
}

func readData(conn *net.TCPConn) (data []byte, err error) {
	ret := bytes.NewBuffer(nil)
	buf := make([]byte, 512)
	for {
		n, er := conn.Read(buf)
		ret.Write(buf[:n])
		if er != nil {
			if er == io.EOF {
				break
			} else {
				fmt.Println("client read error")
				errors.New("client read error")
				return
			}
		}
	}
	return ret.Bytes(), nil
}

func satrtClient() {
	conn, err := connect()
	if err != nil {
		fmt.Println("client connect error")
		return
	} else {
		fmt.Println("client connect success")
		var echo string = "this is a go client."
		data := []byte(echo)
		n, err := writeData(conn, data)
		if err != nil {
			fmt.Println("client write data error")
			conn.Close()
			return
		} else {
			fmt.Println("client write data success ", n)
			rd, err := readData(conn)
			if err != nil {
				fmt.Println("client read data failed")
				conn.Close()
				return
			} else {
				fmt.Println("client read data success")
				sdata := string(rd)
				fmt.Println("server:", sdata)
				conn.Close()
				return
			}
		}
	}
}

func ClientStart() {
	satrtClient()
}
