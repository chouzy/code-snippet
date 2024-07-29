package consul

import (
	"code-snippet/pkg/setting"
	"fmt"
	"github.com/hashicorp/consul/api"
)

// RegisterService 服务注册
func RegisterService(s *setting.ServerSettingS) error {
	conf := api.DefaultConfig()
	conf.Address = s.ConsulAddr
	conf.Token = s.ConsulToken
	client, err := api.NewClient(conf)
	if err != nil {
		fmt.Printf("new client error: %v\n", err)
		return err
	}

	asr := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%v-%v-%v", s.Name, s.IP, s.Port), // 节点名
		Name:    s.Name,                                        // 服务名
		Tags:    s.Tag,                                         // 标签
		Address: s.IP,                                          // 服务ip
		Port:    s.Port,                                        // 服务端口
		Check: &api.AgentServiceCheck{
			Interval:                       s.Interval.String(),                           // 健康检查时间间隔
			GRPC:                           fmt.Sprintf("%v:%v/%v", s.IP, s.Port, s.Name), // 使用grpc执行健康检查, health.Check() 方法进行响应
			DeregisterCriticalServiceAfter: s.Deregister.String(),                         // 注销时间
		},
	}
	if err := client.Agent().ServiceRegister(asr); err != nil {
		fmt.Printf("service register error: %v\n", err)
		return err
	}
	return nil
}
