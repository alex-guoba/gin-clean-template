package global

import (
	"github.com/alex-guoba/gin-clean-template/pkg/setting"
)

var (
	ServerSetting    *setting.ServerSettingS
	AppSetting       *setting.AppSettingS
	LogSetting       *setting.LogSettingS
	DatabaseSetting  *setting.DatabaseSettingS
	RatelimitSetting *setting.RatelimitSettingS
)
