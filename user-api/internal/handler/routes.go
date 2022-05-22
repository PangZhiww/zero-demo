// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "zero-demo/user-api/internal/handler/user"
	"zero-demo/user-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/home/pangzhi/go/src/zero-demo/user-api/api/foo",
				Handler: user.UserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/update",
				Handler: user.UserUpdateHandler(serverCtx),
			},
		},
	)
}
