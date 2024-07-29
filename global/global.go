package global

import (
	"code-snippet/pkg/logger"
	"code-snippet/pkg/setting"
)

var (
	ServerSetting *setting.ServerSettingS
	MysqlSetting  *setting.MySQLSettingS
	RedisSetting  *setting.RedisSettingS
	LoggerSetting *setting.LogSettingS

	Log *logger.Logger
)
