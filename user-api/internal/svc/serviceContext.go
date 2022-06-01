package svc

import (
	"zero-demo/user-api/internal/config"
	"zero-demo/user-api/internal/middleware"
	"zero-demo/user-api/model"
	"zero-demo/user-rpc/usercenter"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     model.UserModel
	UserDataModel model.UserDataModel

	TestMiddleware rest.Middleware

	UserRpcClient usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:         c,
		TestMiddleware: middleware.NewTestMiddleware().Handle,
		UserModel:      model.NewUserModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		UserDataModel:  model.NewUserDataModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		UserRpcClient:  usercenter.NewUsercenter(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
