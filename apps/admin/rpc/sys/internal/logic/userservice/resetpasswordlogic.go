package userservicelogic

import (
	"context"

	"sys/internal/svc"
	"sys/systemclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReSetPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReSetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReSetPasswordLogic {
	return &ReSetPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 重置用户密码
func (l *ReSetPasswordLogic) ReSetPassword(in *systemclient.ReSetPasswordReq) (*systemclient.ReSetPasswordResp, error) {
	// todo: add your logic here and delete this line

	return &systemclient.ReSetPasswordResp{}, nil
}
