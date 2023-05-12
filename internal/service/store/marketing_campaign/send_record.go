package marketing_campaign

import (
	"admin_api/internal/model"
	"admin_api/internal/request"
	"admin_api/internal/response"
	"errors"
	"go.uber.org/zap"
)

type Repo interface {
	ListWithPagination(*model.Pagination) ([]*model.SendRecord, int64, error)
}

type SendRecord struct {
	repo Repo
}

func NewSendRecordService(repo Repo) *SendRecord {
	return &SendRecord{repo: repo}
}

// 展示发送记录
func (sr *SendRecord) DescribeSendRecord(req *request.DescribeSendRecord) ([]*response.DescribeSendRecord, int, error) {

	//使用请求固定页数查询发送记录
	records, total, err := sr.repo.ListWithPagination(&model.Pagination{
		Offset: req.Offset,
		Limit:  req.Limit,
	})
	if err != nil {
		zap.S().Errorf("%+#v", err)
		return nil, -1, errors.New("查询发送记录列表失败")
	}
	if records == nil {
		return []*response.DescribeSendRecord{}, 0, nil
	}

	var resp []*response.DescribeSendRecord
	for _, record := range records {
		resp = append(resp, &response.DescribeSendRecord{
			Id:                record.Id,
			CampaignId:        record.CampaignId,
			CampaignName:      record.CampaignName,
			SurplusCount:      record.SurplusCount,
			TotalCount:        record.TotalCount,
			TotalSuccessCount: record.TotalSuccessCount,
			TotalFailCount:    record.TotalCount - record.TotalSuccessCount,
			CreatedAt:         record.CreatedAt.Local().Format("2006-01-02 15:04:05"),
		})
	}
	return resp, int(total), nil
}
