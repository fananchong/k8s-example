package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3456")
	if err != nil {
		panic(err)
	}

	fmt.Println("connect to ", conn.RemoteAddr().String())

	conn.Write([]byte("hello"))
	for count := 0; true; count++ {
		var b [1024]byte
		ln, err := io.ReadAtLeast(conn, b[:], 1)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b[:ln]))
		conn.Write([]byte(fmt.Sprintf("world %d", count)))
		time.Sleep(2 * time.Second)
	}
}
