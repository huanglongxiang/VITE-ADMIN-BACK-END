package userservicelogic

import (
	"context"

	"sys/internal/svc"
	"sys/systemclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户个人信息
func (l *UserInfoLogic) UserInfo(in *systemclient.InfoReq) (*systemclient.InfoResp, error) {
	// todo: add your logic here and delete this line

	return &systemclient.InfoResp{}, nil
}
