// +build debug

package main

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

	// 仅举例，可以从配置文件中读取

	item := NewEndpoint()
	item.IP = "localhost"
	item.Ports[""] = 3000
	var ips []*Endpoint
	ips = append(ips, item)
	return ips
}

func GetVaildPort(namespace, service string) map[string]int {

	// 仅举例，可以从配置文件中读取

	ports := make(map[string]int)
	ports[""] = 3000
	return ports
}
