// 具体的业务层

package user

import (
	"context"

	"zero-demo/user-api/internal/svc"
	"zero-demo/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line

	//l.svcCtx.UserModel.Findone()
	//l.svcCtx.Redis.set()

	m := map[int64]string{
		1: "张三",
		2: "李四",
	}
	nickname := "unknown"
	if name, ok := m[req.UserId]; ok {
		nickname = name
	}

	return &types.UserInfoResp{
		UserId:   req.UserId,
		Nickname: nickname,
	}, nil
}
