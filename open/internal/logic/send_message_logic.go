package logic

import (
	"context"
	"iot-platform/device/types/device"

	"iot-platform/open/internal/svc"
	"iot-platform/open/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendMessageLogic) SendMessage(req *types.SendMessageRequest) (resp *types.SendMessageResponse, err error) {
	_, err = l.svcCtx.RpcDevice.SendMessage(l.ctx, &device.SendMessageReq{
		ProductKey: req.ProductKey,
		DeviceKey:  req.DeviceKey,
		Data:       req.Data,
	})
	return
}
