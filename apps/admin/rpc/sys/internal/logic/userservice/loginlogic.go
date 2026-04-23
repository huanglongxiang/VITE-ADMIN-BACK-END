package userservicelogic

import (
	"context"

	"sys/internal/svc"
	"sys/systemclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户登录
func (l *LoginLogic) Login(in *systemclient.LoginReq) (*systemclient.LoginResp, error) {
	// todo: add your logic here and delete this line

	return &systemclient.LoginResp{}, nil
}
