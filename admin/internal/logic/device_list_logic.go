package logic

import (
	"context"
	"iot-platform/helper"
	"iot-platform/models"

	"iot-platform/admin/internal/svc"
	"iot-platform/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceListLogic {
	return &DeviceListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeviceListLogic) DeviceList(req *types.DeviceListRequest) (resp *types.DeviceListResponse, err error) {
	// todo: add your logic here and delete this line
	limit := helper.If(req.Size == 0, 20, req.Size).(int)
	offset := helper.If(req.Size == 0, 0, (req.Page-1)*req.Size).(int)

	resp = new(types.DeviceListResponse)
	var count int64
	data := make([]*types.DeviceLisBasic, 0)
	err = models.GetDeviceList(req.Name).Count(&count).Limit(limit).Offset(offset).Find(&data).Error
	if err != nil {
		logx.Error("[DB ERROR]:", err)
		return nil, err
	}
	resp.Count = count
	resp.List = data
	return
}
