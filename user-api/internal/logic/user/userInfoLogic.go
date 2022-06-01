package user

import (
	"context"

	"zero-demo/user-api/internal/svc"
	"zero-demo/user-api/internal/types"
	"zero-demo/user-rpc/pb"

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

	userResp, err := l.svcCtx.UserRpcClient.GetUserInfo(l.ctx, &pb.GetUserInfoReq{
		Id: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	return &types.UserInfoResp{
		UserId:   userResp.Id,
		Nickname: userResp.Nickname,
	}, nil
}
