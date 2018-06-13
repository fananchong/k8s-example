package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

var counter = 0

func handle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	lt, err := net.LookupIP("www.baidu.com")
	if err == nil {
		index := counter % (len(lt))
		counter = counter + 1
		fmt.Fprintf(w, fmt.Sprintf("%v", lt[index]))
	} else {
		fmt.Fprintf(w, "")
	}
}

func main() {
	http.HandleFunc("/", handle)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
}
