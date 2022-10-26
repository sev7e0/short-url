package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"short-link-go/global"
	"short-link-go/initialize"
	"short-link-go/routers"
	"time"
)

func main() {
	initialize.ViperInit()
	initialize.InitLogger()
	err := initialize.InitTrans("zh")
	if err != nil {
		panic(err)
	}
	engine := routers.InitRouter()
	initialize.InitMysqlDb()
	initialize.InitRedisDb()
	httpServer := getCustomHttpServer(engine)
	_ = httpServer.ListenAndServe()
}

// 获取自定义HTTP SERVER
func getCustomHttpServer(engine *gin.Engine) *http.Server {
	// 创建自定义配置服务
	httpServer := &http.Server{
		//ip和端口号
		Addr: global.CONF.App.Addr,
		//调用的处理器，如为nil会调用http.DefaultServeMux
		Handler: engine,
		//计算从成功建立连接到request body(或header)完全被读取的时间
		ReadTimeout: time.Second * 10,
		//计算从request body(或header)读取结束到 response write结束的时间
		WriteTimeout: time.Second * 10,
		//请求头的最大长度，如为0则用DefaultMaxHeaderBytes
		MaxHeaderBytes: 1 << 20,
	}
	return httpServer
}
