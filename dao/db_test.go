package dao

import (
	"code-snippet/dao/mysql"
	"code-snippet/dao/redis"
	"code-snippet/pkg/setting"
	"context"
	"fmt"
	"testing"
	"time"
)

type User struct {
	ID   int64
	Name string
	Age  int // 当通过结构体查询时，gorm只会通过非零值字段进行查询，可以通过使用指针来解决
}

func TestMySQLGorm(t *testing.T) {
	m := setting.MySQLSettingS{
		Host:     "127.0.0.1",
		Port:     "3306",
		UserName: "root",
		Password: "root",
		DbName:   "testdb",
	}
	db := mysql.InitMySQL(&m)
	cli, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("mysql init err: %v \n", err))
	}
	defer cli.Close()
	// 自动迁移
	db.AutoMigrate(&User{})

	// todo: 增加增删改查操作
	user := User{
		Name: "zhangsan",
		Age:  12,
	}

	// 创建
	db.Create(&user)
	t.Log("user: ", user.ID)

	// 查询
	db.First(&user)
	t.Log("user: ", user)

	// 条件查询
	db.Where("name=?", "zhangsan").First(&user)
	t.Log("user: ", user)

	// 指定字段
	db.Select("name, age").First(&user)
	t.Log("user:", user)

}

func TestRedis(t *testing.T) {
	t.Log("start...")
	r := setting.RedisSettingS{
		Addr:     "127.0.0.1:6379",
		Password: "root",
		DB:       0,
		PoolSize: 1,
	}
	rdb := redis.InitRedis(&r)

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 执行命令获取结果
	val, err := rdb.Get(ctx, "key").Result()
	t.Log(val, err)

	// 先获取命令对象
	cmd := rdb.Get(ctx, "key")
	t.Log("val: ", cmd.Val())
	if cmd.Err() != nil {
		t.Error("err: ", cmd.Err())
	}

	// 直接执行命令获取错误
	err = rdb.Set(ctx, "key", 10, time.Hour).Err()
	if err != nil {
		t.Error("err: ", err)
	}

	// 直接执行命令获取值
	val = rdb.Get(ctx, "key").Val()
	t.Log("val: ", val)

	t.Log("end...")
}
