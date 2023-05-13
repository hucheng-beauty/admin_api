package consumer

import (
	"admin_api/internal/model"
	"admin_api/internal/request"
	"admin_api/internal/response"
	"admin_api/pkg/uuid"
	"errors"
	"go.uber.org/zap"
	"time"
)

type CouponRepo interface {
	Save(record *model.Coupon) (*model.Coupon, error)
}

type Coupon struct {
	repo CouponRepo
}

func NewCouponService(repo CouponRepo) *Coupon {
	return &Coupon{repo: repo}
}

func (c *Coupon) CreateCoupon(req *request.CreateCoupon) (*response.CreateCoupon, error) {
	record, err := c.repo.Save(&model.Coupon{
		BaseModel: model.BaseModel{
			Id:        uuid.New(),
			CreatedAt: time.Now().Local(),
			UpdatedAt: time.Now().Local(),
		},
		CouponId:     req.CouponId,
		CouponName:   req.CouponName,
		CouponType:   req.CouponType,
		CouponAmount: req.CouponAmount,
		CreateTime:   req.CreateTime.Local(),
		UsedTime:     time.Time{}.Local(),
		Description:  req.Description,
		Status:       req.Status,
	})
	if err != nil {
		zap.S().Errorf("%+#v\n", err)
		return nil, errors.New("新建领取记录失败")
	}
	return &response.CreateCoupon{
		Id:       record.Id,
		CouponId: record.CouponId,
	}, nil
}
