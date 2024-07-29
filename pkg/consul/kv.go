package consul

import (
	"bytes"
	"code-snippet/pkg/setting"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
)

// GetConfig 获取consul中的配置
func GetConfig(url, token string, paths ...string) (*setting.Setting, error) {
	// 连接consul
	config := api.DefaultConfig()
	config.Address = url
	config.Token = token
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	// 读取配置文件
	var kv []byte
	for _, path := range paths {
		pair, _, err := client.KV().Get(path, nil)
		if err != nil {
			panic(err)
		}
		if len(kv) > 0 {
			kv = append(kv, []byte("\n\n")...)
		}
		kv = append(kv, pair.Value...)
	}
	// 解析数据
	vp := viper.New()
	vp.SetConfigType("yaml")
	err = vp.ReadConfig(bytes.NewBuffer(kv))
	if err != nil {
		return nil, err
	}

	s := &setting.Setting{
		Vp: vp,
	}
	return s, nil
}
