package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

func main() {

	http.NewRequest("get", "localhost")

	addr, _ := net.ResolveTCPAddr("tcp4")
	conn, err := net.DialTCP("tcp4", nil, addr)
	if err != nil {
		panic(err)
	}

	fmt.Println("connect to ", conn.RemoteAddr().String())

	conn.Write(fmt.Sprintf("hello"))
	for count := 0; true; count++ {
		var b [1024]byte
		ln, err := io.ReadAtLeast(conn, b[:], 1)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(ln))
		conn.Write(fmt.Sprintf("world %d", count))
		time.Sleep(2 * time.Second)
	}
}
