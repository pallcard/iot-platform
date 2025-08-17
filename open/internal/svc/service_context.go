package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"iot-platform/device/device"
	"iot-platform/open/internal/config"
	"iot-platform/user/rpc/user_client"
)

type ServiceContext struct {
	Config    config.Config
	RpcDevice device.Device
	RpcUser   user_client.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		RpcDevice: device.NewDevice(zrpc.MustNewClient(c.DeviceRpc)),
		RpcUser:   user_client.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
