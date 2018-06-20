package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ericchiang/k8s"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
)

type endpoints struct {
	IPs []string `json:"ips"`
}

func GetEndpoints(namespace, service string, base int) []string {
	client, err := k8s.NewInClusterClient()
	if err != nil {
		log.Printf("unexpected error opening a connection against API server: %v\n", err)
		return nil
	}

	ips := make([]string, 0)

	var endpoints corev1.Endpoints
	err = client.Get(context.Background(), namespace, service, &endpoints)
	if err != nil {
		log.Printf("unexpected error obtaining information about service endpoints: %v\n", err)
		return nil
	}

	for _, endpoint := range endpoints.Subsets {

		fmt.Println("address:", endpoint.Addresses)

		for _, address := range endpoint.Addresses {
			port := base + getIndex(*address.Hostname)
			ips = append(ips, fmt.Sprintf("%s:%d", *address.Ip, port))
		}
	}

	return ips
}

func GetVaildPort(base int) int {
	return base + getIndex(os.Getenv("POD_NAME"))
}

func getIndex(name string) int {
	temps := strings.Split(name, "-")
	if len(temps) == 0 {
		return 0
	}
	id, err := strconv.Atoi(temps[len(temps)-1])
	if err != nil {
		return 0
	}
	return id
}

