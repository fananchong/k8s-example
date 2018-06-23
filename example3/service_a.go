package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func handle(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	for i := 0; i < sessionCount; i++ {
		if session := sessions[i]; session != nil {
			session.Write([]byte(strconv.Itoa(i)))
		}
	}
}

func main() {
	go func() {
		rand.Seed(time.Now().UnixNano())
		http.HandleFunc("/", handle)
		err := http.ListenAndServe(":8000", nil)
		if err != nil {
			log.Fatal("ERROR: ", err)
			panic("")
		}
	}()
	for {
		checkEndpoints()
		time.Sleep(5 * time.Second)
	}
}

var sessions [1024]*net.TCPConn
var sessionCount int
var mutex sync.Mutex

func checkEndpoints() {
	if eps := GetEndpoints("k8s-example3", "service-b"); eps != nil {
		sessionCount = len(eps)
		for _, ep := range eps {
			mutex.Lock()
			if sessions[ep.Index] == nil {
				sessions[ep.Index] = newSession(ep.Index, ep.IP, ep.Ports[""])
			}
			mutex.Unlock()
		}
	}
}

func newSession(index int, ip string, port int) *net.TCPConn {
	addr, _ := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d", ip, port))
	conn, err := net.DialTCP("tcp4", nil, addr)
	if err != nil {
		panic(err)
	}

	fmt.Println("connect to ", conn.RemoteAddr().String())

	go func() {
		for {
			var tempbuf [1024]byte
			readnum, err := io.ReadAtLeast(conn, tempbuf[:], 1)
			if err != nil {
				fmt.Println(err)
				mutex.Lock()
				sessions[index] = nil
				mutex.Unlock()
				return
			}
			fmt.Println("recv data:", string(tempbuf[:readnum]))
		}
	}()
	return conn
}
