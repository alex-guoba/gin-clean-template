package global

import (
	"github.com/alex-guoba/gin-clean-template/pkg/logger"
	"github.com/alex-guoba/gin-clean-template/pkg/setting"
)

var (
	ServerSetting    *setting.ServerSettingS
	AppSetting       *setting.AppSettingS
	DatabaseSetting  *setting.DatabaseSettingS
	RatelimitSetting *setting.RatelimitSettingS
	Logger           *logger.Logger
)
