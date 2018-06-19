package main

import (
	"fmt"
	"net"
)

func main() {
	port := getVaildPort()
	addr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	lis, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}

	fmt.Println("start listen:", port)

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

func getVaildPort() int {
	port := 10000
	for {
		port = port + 1
		address := fmt.Sprintf(":%d", port)
		tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
		if err != nil {
			continue
		}
		listener, err := net.ListenTCP("tcp", tcpAddr)
		if err != nil {
			if listener != nil {
				listener.Close()
			}
			continue
		}
		listener.Close()
		return port
	}
	return 0
}
