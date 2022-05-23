// 初始化业务依赖 存放依赖

package svc

import (
	"zero-demo/user-api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	// Kafka  kafa.Server
	// UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		// Kafka:     kafka.New(Kafaurl, topic),
		// Redis:     redis.New(url, passwd),
		// UserModel: model.NewUserModel(),
	}
}
