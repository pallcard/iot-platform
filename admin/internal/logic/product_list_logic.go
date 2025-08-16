package logic

import (
	"context"
	"iot-platform/helper"
	"iot-platform/models"

	"iot-platform/admin/internal/svc"
	"iot-platform/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductListLogic {
	return &ProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductListLogic) ProductList(req *types.ProductListRequest) (resp *types.ProductListResponse, err error) {
	limit := helper.If(req.Size == 0, 20, req.Size).(int)
	offset := helper.If(req.Size == 0, 0, (req.Page-1)*req.Size).(int)

	resp = new(types.ProductListResponse)
	var count int64
	list := make([]*types.ProductLisBasic, 0)
	err = models.GetProductList(req.Name).Count(&count).Limit(limit).Offset(offset).Find(&list).Error
	if err != nil {
		logx.Error("[DB ERROR]:", err)
		return
	}
	for _, v := range list {
		v.CreatedAt = helper.RFC3339ToNormalTime(v.CreatedAt)
	}

	resp.Count = count
	resp.List = list
	return
}
