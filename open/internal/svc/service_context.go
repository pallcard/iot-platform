package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"iot-platform/device/device"
	"iot-platform/open/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	RpcDevice device.Device
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		RpcDevice: device.NewDevice(zrpc.MustNewClient(c.Rpc)),
	}
}
