package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

// GetService 从consul中获取address，port
func GetService(client *api.Client, serviceName, tag string) []*api.ServiceEntry {
	var lastIndex uint64
	service, meta, err := client.Health().Service(serviceName, tag, true, &api.QueryOptions{
		WaitIndex: lastIndex,
	})
	if err != nil {
		fmt.Printf("Get service info err: %v\n", err)
	}
	lastIndex = meta.LastIndex
	return service
}
