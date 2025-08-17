package logic

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"iot-platform/api"
	"iot-platform/helper"
	"iot-platform/models"

	"iot-platform/admin/internal/svc"
	"iot-platform/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceCreateLogic {
	return &DeviceCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeviceCreateLogic) DeviceCreate(req *types.DeviceCreateRequest) (resp *types.DeviceCreateResponse, err error) {
	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		// 1. 数据库中新增设备
		deviceBasic := &models.DeviceBasic{
			Identity:        uuid.New().String(),
			ProductIdentity: req.ProductIdentity,
			Name:            req.Name,
			Key:             uuid.New().String(),
			Secret:          uuid.New().String(),
		}
		err = tx.Create(deviceBasic).Error
		if err != nil {
			logx.Error("[DB ERROR] : ", err)
			return err
		}

		// 2. EMQX 中新增认证设备
		err = api.CreateAuthUser(&api.CreateAuthUserRequest{
			UserId:   deviceBasic.Key,
			Password: helper.Md5(deviceBasic.Key + deviceBasic.Secret),
		})
		if err != nil {
			logx.Error("[CreateAuthUse ERROR] : ", err)
			return err
		}

		return nil
	})

	return
}
