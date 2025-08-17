package logic

import (
	"context"
	"errors"
	"fmt"
	"iot-platform/device/internal/mqtt"

	"iot-platform/device/internal/svc"
	"iot-platform/device/types/device"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendMessageLogic) SendMessage(in *device.SendMessageReq) (*device.SendMessageResp, error) {
	if in.ProductKey == "" || in.DeviceKey == "" || in.Data == "" {
		return nil, errors.New("参数异常")
	}

	topic := fmt.Sprintf("/sys/%s/%s/receive", in.ProductKey, in.DeviceKey)
	if publish := mqtt.MC.Publish(topic, 0, false, in.Data); publish.Wait() && publish.Error() != nil {
		logx.Error("[PUBLISH ERROR]:", publish.Error())
		return nil, publish.Error()
	}

	return &device.SendMessageResp{}, nil
}
