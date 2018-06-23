package main

import (
	"fmt"
	"net"
)

func main() {
	ports := GetVaildPort("k8s-example3", "service-b")
	if ports == nil {
		panic("")
	}
	fmt.Println("ports:", ports)
	port := ports[""]
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
