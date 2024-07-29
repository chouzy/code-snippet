package mysql

import (
	"code-snippet/pkg/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitMySQL 初始化MySQL连接
func InitMySQL(m *setting.MySQLSettingS) *gorm.DB {
	if m.DbName == "" {
		return nil
	}
	sqlConf := mysql.Config{
		DSN:                       fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.UserName, m.Password, m.Host, m.Port, m.DbName),
		DefaultStringSize:         255,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	db, err := gorm.Open(mysql.New(sqlConf))
	if err != nil {
		panic(fmt.Sprintf("MySQL init failed, err: %v \n", err))
	}
	fmt.Println("MySQL init success.")
	return db
}
