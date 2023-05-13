package consumer

import (
	"admin_api/internal/model"
	"admin_api/internal/request"
	"admin_api/internal/response"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ConsumerRepo interface {
	ListByAccountIds([]string) ([]*model.Consumer, int64, error)
}

type Consumer struct {
	repo ConsumerRepo
}

func NewConsumerService(repo ConsumerRepo) *Consumer {
	return &Consumer{repo: repo}
}

func (c *Consumer) DescribeConsumer(req *request.DescribeConsumer, p *model.Pagination) ([]*response.DescribeConsumer, int, error) {
	//获取账户信息
	cs, total, err := c.repo.ListByAccountIds(req.AccountIds)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*response.DescribeConsumer{}, 0, nil
		} else {
			zap.S().Errorf("%+#v\n", err)
			return nil, -1, errors.New("获取用户列表失败")
		}
	}

	var resp []*response.DescribeConsumer
	for _, consumer := range cs {
		resp = append(resp, &response.DescribeConsumer{
			ConsumerId: consumer.Id,
			AccountId:  consumer.AccountId,
			OpenId:     consumer.OpenId,
		})
	}
	return resp, int(total), nil
}
