package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"iot-platform/helper"
	"iot-platform/models"
	"sort"

	"iot-platform/user/rpc/internal/svc"
	"iot-platform/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type OpenAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOpenAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OpenAuthLogic {
	return &OpenAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OpenAuthLogic) OpenAuth(in *user.OpenAuthReq) (*user.OpenAuthRsp, error) {
	data := make(map[string]interface{})
	if err := json.Unmarshal(in.Body, &data); err != nil {
		return nil, err
	}
	ub := new(models.UserBasic)
	err := l.svcCtx.DB.Model(new(models.UserBasic)).Select("app_secret").Where("app_key = ?", data["app_key"]).First(ub).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err.Error())
		return nil, err
	}

	arr := make([]string, 0)
	for k, _ := range data {
		arr = append(arr, k)
	}
	sort.Strings(arr)

	var s string
	for _, v := range arr {
		if v != "sign" {
			s += data[v].(string)
		}
	}
	md5 := helper.Md5(s)
	fmt.Println(md5)
	if md5 != data["sign"].(string) {
		return nil, errors.New("签名不正确")
	}
	return &user.OpenAuthRsp{}, nil
}
