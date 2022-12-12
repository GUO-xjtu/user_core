package global

import (
	"github.com/xiangqin/user_core/common/logger"
	"github.com/xiangqin/user_core/common/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	EmailSetting    *setting.EmailSettingS

)

var (
	Logger *logger.Logger
)
