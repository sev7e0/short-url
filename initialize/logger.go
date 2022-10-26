package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"short-link-go/global"
	"short-link-go/utils"
)

func InitLogger() {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{
		fmt.Sprintf("%slog_%s.log", global.CONF.App.LogsAddress, utils.GetNowFormatTodayTime()), //
		"stdout",
	}
	logger, _ := config.Build()
	zap.ReplaceGlobals(logger)
	global.LOGGER = logger
}
