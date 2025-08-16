package logic

import (
	"context"
	"errors"
	"iot-platform/helper"
	"iot-platform/models"

	"iot-platform/user/internal/svc"
	"iot-platform/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) UserLogin(req *types.UserLoginRequest) (resp *types.UserLoginResponse, err error) {
	resp = new(types.UserLoginResponse)

	ub := new(models.UserBasic)
	err = l.svcCtx.DB.Where("name = ? and password = ?", req.Username, req.Password).Find(ub).Error
	if err != nil || ub.Name == "" {
		logx.Error("[DB ERROR]:", err)
		err = errors.New("用户名密码不对")
		return
	}
	token, err := helper.GenerateToken(ub.ID, ub.Identity, ub.Name, 3600*24*30)
	if err != nil {
		logx.Error("[DB ERROR]:", err)
		err = errors.New("用户名密码不对")
		return
	}
	resp.Token = token
	return
}
