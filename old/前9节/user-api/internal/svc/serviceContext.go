package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	model "zero-demo/user-api/genModel"
	"zero-demo/user-api/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlConn),
	}
}
