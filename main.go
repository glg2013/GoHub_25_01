package main

import (
	"GoHub_25_01/bootstrap"
	config2 "GoHub_25_01/config"
	"GoHub_25_01/pkg/config"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	// 加载 config 目录下的配置信息
	config2.Initialize()
}

func main() {

	// 配置初始化，依赖命令行 --env 参数
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	// 初始化 Logger
	bootstrap.SetupLogger()

	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)

	// new 一个 Gin Engine 实例
	router := gin.New()

	// 初始化 DB
	bootstrap.SetupDB()

	// 初始化 Redis
	bootstrap.SetupRedis()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	//logger.Dump(captcha.NewCaptcha().VerifyCaptcha("NZEb3sk2d6FUSUb4Um1v", "567526"), "正确的答案")
	//logger.Dump(captcha.NewCaptcha().VerifyCaptcha("NZEb3sk2d6FUSUb4Um1v", "000000"), "错误的答案")

	// 运行服务
	err := router.Run(":3000")
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}
}
