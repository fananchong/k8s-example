package main


import (
	"fmt"
	"context"
	"encoding/json"
	"log"
	"strings"
	"unicode"
	"unsafe"

	"github.com/ericchiang/k8s"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
)

type endpoints struct {
	IPs []string `json:"ips"`
}

func GetEndpoints(namespace, service) []string {
	ns := C.GoString(namespace)
	svc := C.GoString(service)

	client, err := k8s.NewInClusterClient()
	if err != nil {
		log.Printf("unexpected error opening a connection against API server: %v\n", err)
		return nil
	}

	ips := make([]string, 0)

	var endpoints corev1.Endpoints
	err = client.Get(context.Background(), ns, svc, &endpoints)
	if err != nil {
		log.Printf("unexpected error obtaining information about service endpoints: %v\n", err)
		return nil
	}

	for _, endpoint := range endpoints.Subsets {
		for _, address := range endpoint.Addresses {
			ips = append(ips, *address.Ip)
		}
	}

	return ips
}
