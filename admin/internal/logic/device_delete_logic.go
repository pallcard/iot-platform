package logic

import (
	"context"
	"gorm.io/gorm"
	"iot-platform/models"

	"iot-platform/admin/internal/svc"
	"iot-platform/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceDeleteLogic {
	return &DeviceDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeviceDeleteLogic) DeviceDelete(req *types.DeviceDeleteRequest) (resp *types.DeviceDeleteResponse, err error) {
	deviceBasic := new(models.DeviceBasic)
	err = l.svcCtx.DB.Model(new(models.DeviceBasic)).Select("key").Where("identity = ?", req.Identity).
		Find(deviceBasic).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
		return
	}

	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		// 1. 数据库中删除设备
		err = tx.Where("identity = ?", req.Identity).Delete(new(models.DeviceBasic)).Error
		if err != nil {
			logx.Error("[DB ERROR] : ", err)
			return err
		}

		return nil
	})

	return
}
