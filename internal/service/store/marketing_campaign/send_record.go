package marketing_campaign

import (
	"admin_api/internal/model"
	"admin_api/internal/request"
	"admin_api/internal/response"
	"admin_api/pkg/uuid"
	"errors"
	"go.uber.org/zap"
	"time"
)

type Repo interface {
	ListWithPagination(*model.Pagination) ([]*model.SendRecord, int64, error)
	Save(*model.SendRecord) (*model.SendRecord, error)
	Update(*model.SendRecord) (*model.SendRecord, error)
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

func (sr *SendRecord) CreateSendRecord(req *request.CreateSendRecord) (*response.CreateSendRecord, error) {
	record, err := sr.repo.Save(&model.SendRecord{
		BaseModel: model.BaseModel{
			Id:        uuid.New(),
			CreatedAt: time.Now().Local(),
			UpdatedAt: time.Now().Local(),
		},
		CampaignId:   req.CampaignId,
		CampaignName: req.CampaignName,
		AccountIds:   req.AccountIds,
		TotalCount:   req.TotalCount,
	})
	if err != nil {
		zap.S().Errorf("%+#v\n", err)
		return nil, errors.New("创建发送记录失败")
	}
	return &response.CreateSendRecord{Id: record.Id}, nil
}

func (sr *SendRecord) UpdateSendRecord(req *request.UpdateSendRecord) (*response.UpdateSendRecord, error) {
	if req.Id == "" {
		return nil, errors.New("发送记录Id不可为空")
	}
	if req.CampaignId == "" {
		return nil, errors.New("活动Id不可为空")
	}

	var stockSendRecordInfo []model.StockSendRecord
	for _, record := range req.StockSendRecordInfo {
		stockSendRecordInfo = append(stockSendRecordInfo, model.StockSendRecord{
			StockId:      record.StockId,
			Count:        record.Count,
			SuccessCount: record.SuccessCount,
		})
	}

	record, err := sr.repo.Update(&model.SendRecord{
		BaseModel: model.BaseModel{
			Id:        req.Id,
			UpdatedAt: time.Now().Local(),
		},
		CampaignId:        req.CampaignId,
		SurplusCount:      req.SurplusCount,
		TotalCount:        req.TotalCount,
		TotalSuccessCount: req.TotalSuccessCount,
		StockSendInfo:     stockSendRecordInfo,
	})
	if err != nil {
		return nil, err
	}
	return &response.UpdateSendRecord{
		Id:                record.Id,
		CampaignId:        record.CampaignId,
		CampaignName:      record.CampaignName,
		TotalCount:        record.TotalCount,
		TotalSuccessCount: record.TotalSuccessCount,
		CreatedAt:         record.CreatedAt.Local().Format("2006-01-02 15:04:05"),
	}, nil
}
