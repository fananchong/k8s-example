package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"
)

func handle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	lt, err := net.LookupIP("service-b")
	if err == nil {
		index := rand.Intn(len(lt))
		fmt.Fprintf(w, fmt.Sprintf("%v", lt[index]))
	} else {
		fmt.Fprintf(w, "")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", handle)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
}

