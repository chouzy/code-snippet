package main

import (
	"code-snippet/global"
	"code-snippet/pkg/consul"
	"code-snippet/pkg/logger"
	"code-snippet/pkg/setting"
	"flag"
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/url"
	"runtime"
	"strings"
	"time"
)

var (
	consulUrl   string
	consulToken string
	ip          string
	port        int
	config      string

	debug bool
)

func setupFlag() error {
	flag.StringVar(&consulUrl, "url", "http://127.0.0.1:8500", "consul服务地址")
	flag.StringVar(&consulToken, "token", "", "consul检验token")
	flag.StringVar(&ip, "ip", "", "服务IP")
	flag.IntVar(&port, "port", 0, "服务端口")
	flag.StringVar(&config, "config", "code-snippet/conf", "配置文件路径")
	flag.BoolVar(&debug, "debug", false, "开启debug")
	flag.Parse()

	return nil
}

func setupSetting() error {
	var s *setting.Setting
	var err error
	if debug {
		s, err = setting.NewSetting("./conf/config.yaml")
	} else {
		s, err = consul.GetConfig(consulUrl, consulToken, strings.Split(config, ",")...)
	}
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Logger", &global.LoggerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Mysql", &global.MysqlSetting)
	if err != nil {
		return err
	}

	consulUrlParse, err := url.Parse(consulUrl)
	if err != nil {
		return err
	}
	global.ServerSetting.ConsulAddr = consulUrlParse.Host

	global.ServerSetting.Interval *= time.Second
	global.ServerSetting.Deregister *= time.Minute

	return nil
}

func setupLogger() error {
	fileName := global.LoggerSetting.LogSavePath + "/" +
		global.LoggerSetting.LogFileName + "_" + time.Now().Format("20060102150405") + global.LoggerSetting.LogFileExt
	global.Log = logger.NewLogger(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    global.LoggerSetting.MaxSize,
		MaxAge:     global.LoggerSetting.MaxAge,
		MaxBackups: global.LoggerSetting.MaxBackups,
		Compress:   global.LoggerSetting.Compress,
	}, "")

	return nil
}

func initSetting() error {
	err := setupFlag()
	if err != nil {
		return fmt.Errorf("init.setupFlag err: %w", err)
	}
	err = setupSetting()
	if err != nil {
		return fmt.Errorf("init.setupSetting err: %w", err)
	}
	err = setupLogger()
	if err != nil {
		return fmt.Errorf("init.setupLogger err: %w", err)
	}

	return nil
}

func main() {
	err := initSetting()
	if err != nil {
		log.Fatal(err)
	}

	// 使用量监控日志
	go func() {
		for {
			time.Sleep(time.Minute * 5)

			stats := runtime.MemStats{}
			runtime.ReadMemStats(&stats)
			alloc, syss, idle, inuse, stack := float64(stats.HeapAlloc)/1024/1024, float64(stats.HeapSys)/1024/1024,
				float64(stats.HeapIdle)/1024/1024, float64(stats.HeapInuse)/1024/1024, float64(stats.StackInuse)/1024/1024
			global.Log.WithFields(logger.Fields{
				"GoRoutines": runtime.NumGoroutine(),
				"Memory":     fmt.Sprintf("heap_alloc=%.2fMB heap_sys=%.2fMB heap_idle=%.2fMB heap_inuse=%.2fMB stack=%.2fMB\n", alloc, syss, idle, inuse, stack),
			}).Infof("监控日志")
		}
	}()

	// 启动服务
}
