package main

import (
	"flag"
	"log"
	"time"

	"github.com/GUO-xjtu/kitex_gen/user/user"
	"github.com/xiangqin/user_core/common/logger"
	"github.com/xiangqin/user_core/common/setting"
	"github.com/xiangqin/user_core/common/tracer"
	"github.com/xiangqin/user_core/common/util"
	"github.com/xiangqin/user_core/common/version"
	"github.com/xiangqin/user_core/global"
	"github.com/xiangqin/user_core/internal/model"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	flagRun()
	// 配置初始化，读取到全局model里面
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	// 数据库orm配置初始化
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
	err = setupFlake()
	if err != nil {
		log.Fatalf("init.setupFlake err: %v", err)
	}
}

func flagRun() {
	flag.Parse()
	version.CmdParseVersion()
}
func setupSetting() error {
	newSetting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}
	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.AppSetting.DefaultContextTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	// 千万小心全局变量不要被局部变量覆盖，也就是不要用 := 符号
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		version.AppName,
		"10.21.107.4:6831",
	)
	if err != nil {
		return err
	}
	// 放到全局变量中，供后续中间件或者不同的自定义Span中打点使用
	global.Tracer = jaegerTracer
	return nil
}

func setupFlake() error {
	err := util.InitFlake(global.ServerSetting.MachineID)
	return err
}

func main() {
	svr := user.NewServer(new(UserImpl))
	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
