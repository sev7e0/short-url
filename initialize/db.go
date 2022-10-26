package initialize

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"short-link-go/global"
)

// InitMysqlDb 初始化数据库
func InitMysqlDb() {
	mysqlConf := global.CONF.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConf.User, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.DbName)
	open, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败！")
	}
	db, err := open.DB()
	db.SetMaxIdleConns(mysqlConf.MaxConn)
	db.SetMaxOpenConns(mysqlConf.MaxOpen)
	global.DB = open
}

// InitRedisDb 初始化数据库
func InitRedisDb() {
	redisConfig := global.CONF.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	global.Redis = client
}
