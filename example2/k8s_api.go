package main

import (
	"context"
	"log"

	"github.com/ericchiang/k8s"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
)

type endpoints struct {
	IPs []string `json:"ips"`
}

func GetEndpoints(namespace, service string) []string {
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
		for _, address := range endpoint.Addresses {
			ips = append(ips, *address.Ip)
		}
	}

	return ips
}
