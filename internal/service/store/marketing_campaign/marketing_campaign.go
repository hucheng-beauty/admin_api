package marketing_campaign

import (
	"admin_api/internal/model"
	"admin_api/internal/request"
	"admin_api/pkg/uuid"
	"admin_api/utils"

	"fmt"
	"time"
)

func NewBM() model.BM {
	n := time.Now()
	return model.BM{
		Id:        uuid.New(),
		CreatedAt: n,
		UpdatedAt: n,
	}
}

// 创建活动和券批次
func (s *MarCampaignService) CreateMarCampaignAndCouponBatch(in *request.CreateMarketingCampaignRequest, userId string) (*model.MarketingCampaign, []*model.CouponBatch, error) {
	if userId == `` {
		return nil, nil, fmt.Errorf("not user_id")
	}
	//获取营销活动和卷批次
	mr, cbs := in.ToModel()
	//if cbs ==0 则没有卷批次
	if len(cbs) == 0 {
		return nil, nil, fmt.Errorf("the CouponBatchs length should not eq 0")
	}

	mr.BM = NewBM()
	//生成活动编号
	mr.CampaignNu = utils.DefaultGenerator("HD")

	//返回所有可用的卷模板id和对应的卷模板
	templateMap, err := s.CouponBatchTemplates2Map(cbs)
	if err != nil {
		return nil, nil, err
	}
	//
	mr.State = model.Submit                // 创建状态
	mr.CouponNumber = 0                    //活动卷数量
	mr.MaxAmount = 0                       //总预算
	mr.CouponBatchNumber = int64(len(cbs)) //活动卷批次数量
	mr.UserId = userId                     //账户id
	//设置卷批次信息
	for _, item := range cbs {
		template := templateMap[item.TemplateID]

		if template == nil {
			return nil, nil, fmt.Errorf("not found template by id :%s", item.TemplateID)
		}
		item.CouponTemplateBaseInfo = template.CouponTemplateBaseInfo //关联卷模板
		item.Id = uuid.New()
		item.MarketingCampaignID = mr.Id
		item.MaxAmount = item.MaxCoupons * item.CouponAmount
		item.CreatedAt = mr.CreatedAt
		item.UpdatedAt = mr.UpdatedAt
		// 批次计算到活动身上
		mr.CouponNumber += item.MaxCoupons
		mr.MaxAmount += item.MaxAmount

		if item.MaxCouponsPerUser > 60 {
			return nil, nil, fmt.Errorf("单个用户可领个数不能超过60")
		}
		if item.MaxAmountByDay*item.MaxCoupons > item.MaxAmount {
			return nil, nil, fmt.Errorf("每日发放上限不能大于总预算")
		}
	}
	//创建营销活动
	err = s.mcr.Create(mr)
	if err != nil {
		return nil, nil, err
	}
	//创建卷批次
	err = s.cbr.InsertMany(cbs)
	return mr, cbs, err
}

// 修改券状态并且写入日志  根据活动id和状态
func (s *MarCampaignService) SetCampaignState(campaignId string, state model.State) (*model.MarketingCampaign, error) {
	return s.actCampaignState(campaignId, state, true)
}

// todo 统一使用写日志
func (s *MarCampaignService) actCampaignState(id string, state model.State, disAct bool) (*model.MarketingCampaign, error) {
	campaign, err := s.MarCampaignById(id)
	if err != nil {
		return nil, err
	}
	//没有该条记录
	if campaign == nil {
		return nil, nil
	}
	// 不做操作校验
	if !disAct {
		if !campaign.CanAction() {
			return nil, nil
		}
	}
	campaign.State = state
	// 修改状态
	err = s.UpdateMarCampaignStateById(campaign.Id, campaign.State)
	if err != nil {
		return nil, err
	}
	// 创建活动日志
	err = s.CreateMarketingCampaignLogWithModel(campaign)
	if err != nil {
		return nil, err
	}
	// 写入券日志
	return campaign, s.CreateCouponLogWithMarCampaign(campaign)
}

// 找到活动id
func (s *MarCampaignService) MarCampaignById(id string) (*model.MarketingCampaign, error) {
	return s.mcr.FindById(id)
}

// 更改活动状态
func (s *MarCampaignService) UpdateMarCampaignStateById(id string, state model.State) error {
	return s.mcr.UpdateStateById(id, state)
}

func (s *MarCampaignService) UpdateMarCampaignCouponSurplusNumber(id string, successCount int) error {
	if id == `` || successCount == 0 {
		return nil
	}
	return s.mcr.UpdateSurplusNumberById(id, successCount)
}
