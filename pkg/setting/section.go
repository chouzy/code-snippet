package setting

import "time"

type ServerSettingS struct {
	Name        string
	Tag         []string
	IP          string
	Port        int
	ConsulAddr  string
	ConsulToken string
	Interval    time.Duration
	Deregister  time.Duration
	RunMode     string
}

type MySQLSettingS struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	UserName string `mapstructure:"userName" json:"userName" yaml:"userName"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DbName   string `mapstructure:"dbName" json:"dbName" yaml:"dbName"`
}

type RedisSettingS struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DB       int    `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	PoolSize int    `mapstructure:"poolSize" json:"poolSize" yaml:"poolSize"`
}

type LogSettingS struct {
	LogFileName string
	LogFileExt  string
	LogSavePath string
	MaxSize     int
	MaxAge      int
	MaxBackups  int
	Compress    bool
}
