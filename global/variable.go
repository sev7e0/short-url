package global

import (
	ut "github.com/go-playground/universal-translator"
	redis2 "github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"short-link-go/conf"
)

const (
	ConfigFile = "./app.yaml" // 配置文件
)

var (
	CONF   conf.ViperConfig
	LOGGER *zap.Logger
	Trans  ut.Translator
	DB     *gorm.DB
	Redis  *redis2.Client
)
