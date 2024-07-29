package consul

import (
	"code-snippet/pkg/setting"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"math/rand"
	"net"
	"strings"
	"testing"
	"time"
)

var (
	ServerSetting *setting.ServerSettingS
)

var (
	consulUrl   = "http://127.0.0.1:8500"
	configPath  = "goProject/conf"
	consulToken = ""
)

func InitConsul(s *setting.ServerSettingS) (*api.Client, error) {
	conf := api.DefaultConfig()
	conf.Address = s.ConsulAddr
	conf.Token = s.ConsulToken
	client, err := api.NewClient(conf)
	if err != nil {
		fmt.Printf("new client error: %v\n", err)
		return nil, err
	}
	return client, nil
}

func TestConsul(t *testing.T) {
	// 获取consul设置
	cf, err := GetConfig(consulUrl, consulToken, strings.Split(configPath, ",")...)
	if err != nil {
		t.Fatal(fmt.Sprintf("get config err: %v\n", err))
	}
	// 解析consul到全局
	err = cf.ReadSection("Server", &ServerSetting)
	if err != nil {
		t.Fatal(fmt.Sprintf("parse err: %v\n", err))
	}
	conf, _ := json.Marshal(ServerSetting)
	t.Log(string(conf))

	ServerSetting.Interval *= time.Second
	ServerSetting.Deregister *= time.Minute

	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", ServerSetting.IP, ServerSetting.Port))
	if err != nil {
		t.Fatal(fmt.Sprintf("listen error: %v\n", err))
	}
	// 服务注册
	_ = RegisterService(ServerSetting)
	// 创建grpc服务
	server := grpc.NewServer()
	// 健康检查
	grpc_health_v1.RegisterHealthServer(server, &HealthImpl{})
	// reflection.Register(server) // 使用grpcurl、grpcui工具须添加该行
	t.Log("service is running...")

	go func() {
		if err = server.Serve(lis); err != nil {
			t.Error(fmt.Sprintf("start service error: %v\n", err))
			return
		}
	}()

	// 服务发现
	client, _ := InitConsul(ServerSetting)
	service := GetService(client, "serverName", "")
	if service == nil || len(service) <= 0 {
		t.Log("len(service) == 0")
	} else {
		index := rand.Intn(len(service))
		t.Log(fmt.Sprintf("Addressd: %v, Port: %v\n", service[index].Service.Address, service[index].Service.Port))
	}

	time.Sleep(10 * time.Second)
}
