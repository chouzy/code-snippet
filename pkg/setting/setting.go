package setting

import (
	"fmt"
	"github.com/spf13/viper"
)

type Setting struct {
	Vp *viper.Viper
}

func NewSetting(path string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigFile(path)
	err := vp.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("config file err: %w", err)
	}
	return &Setting{Vp: vp}, nil
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	return s.Vp.UnmarshalKey(k, v)
}
