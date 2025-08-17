package logic

import (
	"context"
	"iot-platform/admin/internal/svc"
	"iot-platform/admin/internal/types"
	"iot-platform/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceModifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceModifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceModifyLogic {
	return &DeviceModifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeviceModifyLogic) DeviceModify(req *types.DeviceModifyRequest) (resp *types.DeviceModifyResponse, err error) {
	err = l.svcCtx.DB.Where("identity = ?", req.Identity).Updates(&models.DeviceBasic{
		ProductIdentity: req.ProductIdentity,
		Name:            req.Name,
	}).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
	}

	return
}
