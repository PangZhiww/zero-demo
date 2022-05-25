package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	// config.go中的名字 要与 user-api.yaml中名字一样
	DB struct {
		DataSource string
	}
}
