package main

import (
	"fmt"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:3456")
	lis, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := lis.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("on connect. addr =", conn.RemoteAddr())

		go func() {
			for {
				var b [1024]byte
				ln, err := conn.Read(b[:])
				if err != nil {
					fmt.Println(err)
					break
				}
				conn.Write([]byte(conn.LocalAddr().String() + " " + string(b[:ln])))
			}
		}()
	}
}
