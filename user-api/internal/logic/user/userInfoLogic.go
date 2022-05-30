package user

import (
	"context"
	"fmt"
	"zero-demo/user-api/model"

	"github.com/pkg/errors"

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

	if err := l.testOne(); err != nil {
		logx.Errorf("err:%+v", err)
	}

	// fmt.Println("user info ing")
	// logx.Info("user logx info ing")
	// logx.Error("user logx error ing")

	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	if err != nil && err != model.ErrNotFound {
		fmt.Println("Err :", err)
		return nil, errors.New("查询数据失败")
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}

	return &types.UserInfoResp{
		UserId:   user.Id,
		Nickname: user.Nickname,
	}, nil
}

func (l *UserInfoLogic) testOne() error {
	return l.testTwo()
}

func (l *UserInfoLogic) testTwo() error {
	return l.testThree()
}

func (l *UserInfoLogic) testThree() error {
	return errors.Wrap(errors.New("这是故意的"), "emmmm")
}
