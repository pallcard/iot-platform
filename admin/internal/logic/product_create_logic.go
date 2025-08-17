package logic

import (
	"context"
	"github.com/google/uuid"
	"iot-platform/models"

	"iot-platform/admin/internal/svc"
	"iot-platform/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCreateLogic {
	return &ProductCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCreateLogic) ProductCreate(req *types.ProductCreateRequest) (resp *types.ProductCreateResponse, err error) {
	err = l.svcCtx.DB.Create(&models.ProductBasic{
		Identity: uuid.New().String(),
		Key:      uuid.New().String(),
		Name:     req.Name,
		Desc:     req.Desc,
	}).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
	}
	return
}
