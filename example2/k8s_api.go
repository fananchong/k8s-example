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

type Endpoint struct {
	IP    string
	Ports map[string]int
}

func NewEndpoint() *Endpoint {
	return &Endpoint{
		Ports: make(map[string]int),
	}
}

func GetEndpoints(namespace, service string) []*Endpoint {
	client, err := k8s.NewInClusterClient()
	if err != nil {
		log.Printf("unexpected error opening a connection against API server: %v\n", err)
		return nil
	}

	var ips []*Endpoint
	var endpoints corev1.Endpoints
	err = client.Get(context.Background(), namespace, service, &endpoints)
	if err != nil {
		log.Printf("unexpected error obtaining information about service endpoints: %v\n", err)
		return nil
	}

	for _, endpoint := range endpoints.Subsets {

		fmt.Println("address:", endpoint.Addresses)
		fmt.Println("ports:", endpoint.Ports)

		item := NewEndpoint()
		for _, address := range endpoint.Addresses {
			item.IP = *address.Ip
			for _, port := range endpoint.Ports {
				item.Ports[*port.Name] = int(*port.Port) + getIndex(*address.Hostname)
			}
			ips = append(ips, item)
		}
	}

	return ips
}

func GetVaildPort(namespace, service string) map[string]int {
	client, err := k8s.NewInClusterClient()
	if err != nil {
		log.Printf("unexpected error opening a connection against API server: %v\n", err)
		return nil
	}
	var endpoints corev1.Endpoints
	err = client.Get(context.Background(), namespace, service, &endpoints)
	if err != nil {
		log.Printf("unexpected error obtaining information about service endpoints: %v\n", err)
		return nil
	}

	ports := make(map[string]int)

	for _, endpoint := range endpoints.Subsets {

		fmt.Println("ports:", endpoint.Ports)

		for _, port := range endpoint.Ports {
			ports[*port.Name] = int(*port.Port) + getIndex(os.Getenv("POD_NAME"))
		}
		break
	}

	return ports
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

