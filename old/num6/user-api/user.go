package main

import (
	"flag"
	"fmt"

	"zero-demo/user-api/internal/config"
	"zero-demo/user-api/internal/handler"
	"zero-demo/user-api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c) // 需要把configFile 中的内容 映射到 c

	ctx := svc.NewServiceContext(c) // 方便调用依赖 初始化业务依赖 ***
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx) // 注册路由

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start() // 启动服务
}
